package controllers

import (
	"myproject/handlers"
	"myproject/initializers"
	"myproject/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AddUser(c *gin.Context) {
	// Tạo biến chứa dữ liệu input
	var requestData struct {
		Email    string `json:"email"`
		UserName string `json:"username"`
		Phone    string `json:"phone"`
		Pass     string `json:"pass"`
	}

	// Lưu dữ liệu input vào biến vừa khởi tạo
	if err := c.BindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}

	email := requestData.Email
	userName := requestData.UserName
	phone := requestData.Phone
	password := requestData.Pass

	// Xử Lí khi một trong các dữ liệu input trống
	if email == "" || userName == "" || phone == "" || password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email, Username, Phone, and Password are required"})
		return
	}

	// Kiểm tra liệu password có mạnh không
	if !handlers.IsStrongPassword(password) {
		c.JSON(http.StatusBadRequest, gin.H{
			// Trả lại các yêu cầu mà một password cần có
			"error": "Password must be at least 12 characters long and include upper/lower case, numbers, and special characters",
		})
		return
	}

	// Băm Password
	hashedPassword, err := handlers.HashPassword(password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	// Check xem Email hay số điện thoại có tồn tại hay chưa
	var existingUser models.User
	if err := initializers.DB.Where("email = ? OR phone_number = ?", email, phone).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Email or phone number already exists"})
		return
	}

	// Tạo token user
	token, err := handlers.GenerateJWT(0, email, phone)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// Tạo biến mới với model user để chuẩn bị thêm vào bảng user
	user := models.User{
		Email:       email,
		UserName:    userName,
		PhoneNumber: phone,
		Pass:        hashedPassword,
		Token:       token,
	}

	// Thêm vào bảng User
	if result := initializers.DB.Create(&user); result.Error != nil {
		// Trả lại thông báo lỗi khi thêm vào không thành công
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

	// Trả lại thông tin vừa thêm vào và thông báo đã thành công
	c.JSON(http.StatusOK, gin.H{
		"ID":       user.ID,
		"UserName": user.UserName,
		"Email":    user.Email,
		"Phone":    user.PhoneNumber,
		"Token":    user.Token,
		"Message":  "Registration successful",
	})
}

func GetAllUsers(c *gin.Context) {
	// Tạo mảng để lưu tất cả người dùng
	var users []models.User

	// Lấy tất cả người dùng có trong cơ sở dữ liệu
	if result := initializers.DB.Find(&users); result.Error != nil {
		// Trả về lỗi nếu không thể lấy danh sách người dùng
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Không thể lấy danh sách người dùng",
		})
		return
	}

	// Kiểm tra xem mảng người dùng có trống không
	if len(users) == 0 {
		// Trả về thông báo nếu không có người dùng nào
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Không có người dùng nào",
		})
		return
	}

	// Trả lại thông tin của tất cả người dùng
	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

func GetUserInfoByID(c *gin.Context) {
	// Lấy ID từ tham số URL
	idParam := c.Param("id")

	// Chuyển đổi ID từ chuỗi sang số nguyên
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			// Trả về lỗi nếu không thể chuyển đổi ID
			"error": "Định dạng ID không hợp lệ",
		})
		return
	}

	var user models.User

	// Lấy dữ liệu người dùng theo ID
	if result := initializers.DB.First(&user, id); result.Error != nil {
		// Nếu không tìm thấy người dùng, trả về lỗi
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Người dùng không tồn tại",
			})
		} else {
			// Nếu có lỗi khác, trả về lỗi chung
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Lỗi hệ thống, không thể truy cập dữ liệu",
			})
		}
		return
	}

	// Trả về thông tin người dùng
	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}
