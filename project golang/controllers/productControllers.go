package controllers

import (
	"myproject/initializers"
	"myproject/models"
	"net/http"
	"regexp"
	"strconv"

	"github.com/gin-gonic/gin"
)

func SeeProduct(c *gin.Context) {
	// Tạo mảng products để lưu tất cả Product
	var products []models.Product

	// Tìm tất cả product trong database
	result := initializers.DB.Find(&products)
	if result.Error != nil {
		// nếu phát hiện lỗi thông qua biến result thì trả lại error
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve products",
		})
		return
	}

	// trả JSON của tất cả Product tìm được
	c.JSON(http.StatusOK, gin.H{
		"products": products,
	})
}

func AddProduct(c *gin.Context) {
	var product models.Product

	if err := c.Bind(&product); err != nil {
		// Trả về lỗi nếu không thể phân tích yêu cầu
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request format",
		})
		return
	}

	// Kiểm tra xem các trường cần thiết có bị bỏ trống không
	if product.NameProduct == "" || product.Descriptions == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "NameProduct and Descriptions cannot be empty",
		})
		return
	}

	// Kiểm tra xem NameProduct có bắt đầu bằng chữ cái và không chứa ký tự đặc biệt không
	isValidName, _ := regexp.MatchString(`^[a-zA-Z][a-zA-Z0-9 ]*$`, product.NameProduct)
	if !isValidName {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid NameProduct. It must start with a letter and contain no special characters.",
		})
		return
	}

	// Kiểm tra xem sản phẩm đã tồn tại trong cơ sở dữ liệu chưa
	var existingProduct models.Product
	if result := initializers.DB.Where("name_product = ?", product.NameProduct).First(&existingProduct); result.Error == nil {
		c.JSON(http.StatusConflict, gin.H{
			"error": "Product already exists",
		})
		return
	}

	// Thêm sản phẩm vào cơ sở dữ liệu
	if result := initializers.DB.Create(&product); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error adding the product",
		})
		return
	}

	// Trả về thông tin sản phẩm đã được thêm
	c.JSON(http.StatusOK, gin.H{
		"message": "Product added successfully",
		"product": product,
	})
}

func UpdateProduct(c *gin.Context) {
	// Lấy ID từ tham số URL
	idParam := c.Param("id")

	// Chuyển đổi id từ string sang integer
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID format",
		})
		return
	}

	// Tạo biến product để chứa dữ liệu từ yêu cầu
	var product models.Product

	// Ràng buộc dữ liệu từ yêu cầu JSON
	if err := c.BindJSON(&product); err != nil {
		// Trả về lỗi nếu không thể phân tích yêu cầu
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request format",
		})
		return
	}

	// Kiểm tra xem các trường cần thiết có bị bỏ trống không
	if product.NameProduct == "" || product.Descriptions == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "NameProduct and Descriptions cannot be empty",
		})
		return
	}

	// Kiểm tra xem NameProduct có bắt đầu bằng chữ cái và không chứa ký tự đặc biệt không
	isValidName, _ := regexp.MatchString(`^[a-zA-Z][a-zA-Z0-9 ]*$`, product.NameProduct)
	if !isValidName {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid NameProduct. It must start with a letter and contain no special characters.",
		})
		return
	}

	// Kiểm tra xem sản phẩm có tồn tại trong cơ sở dữ liệu không
	var existingProduct models.Product
	if result := initializers.DB.First(&existingProduct, id); result.Error != nil {
		// Trả về lỗi nếu không tìm thấy sản phẩm
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Product not found",
		})
		return
	}

	// Cập nhật thông tin sản phẩm
	existingProduct.NameProduct = product.NameProduct
	existingProduct.Descriptions = product.Descriptions

	// Lưu các thay đổi vào cơ sở dữ liệu
	if result := initializers.DB.Save(&existingProduct); result.Error != nil {
		// Trả về lỗi nếu có vấn đề trong việc lưu dữ liệu
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error updating the product",
		})
		return
	}

	// Trả về thông tin sản phẩm đã được cập nhật
	c.JSON(http.StatusOK, gin.H{
		"message": "Product updated successfully",
		"product": existingProduct,
	})
}

func DeleteProduct(c *gin.Context) {
	// Lấy ID từ tham số URL
	idParam := c.Param("id")

	// Chuyển đổi ID từ chuỗi sang số nguyên
	id, err := strconv.Atoi(idParam)
	if err != nil {
		// Trả về lỗi nếu không thể chuyển đổi ID
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID format",
		})
		return
	}

	// Tạo biến để chứa sản phẩm cần xóa
	var product models.Product

	// Tìm sản phẩm trong cơ sở dữ liệu theo ID
	if result := initializers.DB.First(&product, id); result.Error != nil {
		// Trả về lỗi nếu không tìm thấy sản phẩm
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Product not found",
		})
		return
	}

	// Xóa sản phẩm khỏi cơ sở dữ liệu
	if result := initializers.DB.Delete(&product, id); result.Error != nil {
		// Trả về lỗi nếu có vấn đề trong việc xóa sản phẩm
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error deleting the product",
		})
		return
	}

	// Trả về thông báo thành công
	c.JSON(http.StatusOK, gin.H{
		"message": "Product deleted successfully",
	})
}
