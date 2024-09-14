package controllers

import (
	"fmt"
	"myproject/handlers"
	"myproject/initializers"
	"myproject/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func AddAdmin(c *gin.Context) {
	// tạo biến để chứa input data
	var body struct {
		Username    string `json:"username"`
		Password    string `json:"password"`
		Email       string `json:"email"`
		Address     string `json:"address"`
		VipAdmin    string `json:"vipAdmin"`
		PhoneNumber string `json:"phoneNumber"`
	}

	// lưu các input data vào biến vừa tạo
	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	// Kiểm tra xem có thông tin nào bị bỏ trông không
	if body.Username == "" || body.Password == "" || body.Email == "" || body.Address == "" || body.PhoneNumber == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "All fields are required",
		})
		return
	}

	// Kiểm tra nếu username của admin đã tồn tại hay chưa
	var existingAdmin models.Admins
	if err := initializers.DB.Where("username = ?", body.Username).First(&existingAdmin).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{
			// Nếu đã tồn tại username đó trong database của admin thì gửi error và quay về
			"warning": "Admin with this username already exists",
		})
		return
	}

	// Kiểm tra xem password có mạnh không
	if !handlers.IsStrongPassword(body.Password) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Password must be at least 12 characters long and include upper/lowercase letters, numbers, and special characters",
		})
		return
	}

	// Băm mật khẩu
	hashedPassword, err := handlers.HashPassword(body.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})
		return
	}

	// tạo một biến admin với các input data để chuẩn bị thêm
	admin := models.Admins{
		Username:    body.Username,
		Password:    hashedPassword,
		Email:       body.Email,
		Address:     body.Address,
		VipAdmin:    body.VipAdmin,
		PhoneNumber: body.PhoneNumber,
	}

	// Thêm vào database
	if result := initializers.DB.Create(&admin); result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create admin",
		})
		return
	}

	// Return success message
	c.JSON(http.StatusOK, gin.H{
		"message": "Admin created successfully",
	})
}

func Login(c *gin.Context) {
	// Tạo biến body để lưu username, password từ input
	var body struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	// Đọc dữ liệu vào biến body
	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}
	fmt.Println(body)

	// Tìm kiếm admin trong database dựa vào username
	var admin models.Admins
	result := initializers.DB.First(&admin, "username = ?", body.Username)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Admin not found or query error",
		})
		return
	}
	fmt.Println(admin)

	//So sánh mật khẩu trong database và mật khẩu vừa input
	err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(body.Password))
	if err != nil {
		// Trả lại error nếu mật khẩu sai
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid username or password",
		})
		return
	}

	//Trả lại message khi so sánh password trả về đúng
	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
	})
}

func GetAllAdmins(c *gin.Context) {
	// Tạo mảng để lưu tất cả admin
	var admins []models.Admins

	// Lấy tất cả admin có trong cơ sở dữ liệu
	if result := initializers.DB.Find(&admins); result.Error != nil {
		// Trả về lỗi nếu không thể lấy danh sách admin
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Không thể lấy danh sách admin",
		})
		return
	}

	// Kiểm tra xem mảng admin có trống không
	if len(admins) == 0 {
		// Trả về thông báo nếu không có admin nào
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Không có admin nào",
		})
		return
	}

	// Trả lại thông tin của tất cả admin
	c.JSON(http.StatusOK, gin.H{
		"admins": admins,
	})
}
