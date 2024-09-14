package controllers

import (
	"myproject/initializers"
	"myproject/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetProductPackageByProductID(c *gin.Context) {
	// Lấy ID từ params
	productIDParam := c.Param("id")
	productID, err := strconv.Atoi(productIDParam) // Chuyển ID từ dạng string sang integer rồi lưu vào biến productID
	if err != nil {
		// Nếu có lỗi thì trả về error
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid product ID",
		})
		return
	}

	// Tạo biến productPackages để lưu lại dữ liệu tìm được
	var productPackages []models.ProductsPackage
	// Tìm productPackage dựa vào ProductID
	if err := initializers.DB.Where("product_id = ?", productID).Find(&productPackages).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve product packages",
		})
		return
	}

	// Nếu Product đó không có bất cứ dữ liệu productPackage nào thì trả về message
	if len(productPackages) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "No product packages found for this product ID",
		})
		return
	}

	// Trả về thông tin Product Package khi tìm thành công
	c.JSON(http.StatusOK, gin.H{
		"product_packages": productPackages,
	})
}

func GetProductPackageInfo(c *gin.Context) {
	// Lấy Id từ param
	idParam := c.Param("id")

	// Chuyển đổi id từ string sang integer và lưu vào biến id
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			// Báo lỗi nếu có sai sót khi chuyển đổi
			"error": "Invalid ID format",
		})
		return
	}

	var productPackage models.ProductsPackage

	// Lấy product package từ database thông qua ID
	if result := initializers.DB.Preload("Product").First(&productPackage, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			// Báo lỗi nếu tồn tại result.Error đồng nghĩa với có lỗi khi tìm dữ liệu trong database
			"error": "Product package not found",
		})
		return
	}

	// Trả về thông tin product Package
	c.JSON(http.StatusOK, gin.H{
		"productPackage": productPackage,
	})
}
