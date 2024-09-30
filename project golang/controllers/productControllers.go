package controllers

import (
	"myproject/handlers"
	"myproject/initializers"
	"myproject/models"
	"net/http"
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

	// Bind the JSON request body to the product struct
	if err := c.Bind(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	// Validate fields
	if product.Name == "" || product.Descriptions == "" || product.Content == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Required fields are missing"})
		return
	}

	// Check if product already exists
	var existingProduct models.Product
	if result := initializers.DB.Where("slug = ?", product.Slug).First(&existingProduct); result.Error == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Product already exists"})
		return
	}

	// Create slug and thumbnail based on product name
	product.Slug = handlers.GenerateSlug(product.Name)
	product.Thumb = handlers.GenerateThumb(product.Name)

	// Insert new product into the database
	if result := initializers.DB.Create(&product); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product added successfully", "product": product})
}

func UpdateProduct(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var product models.Product
	if err := c.BindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	// Find existing product
	var existingProduct models.Product
	if result := initializers.DB.First(&existingProduct, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	// Update fields
	existingProduct.Name = product.Name
	existingProduct.Descriptions = product.Descriptions
	existingProduct.Content = product.Content
	existingProduct.Slug = handlers.GenerateSlug(product.Name)
	existingProduct.Thumb = handlers.GenerateThumb(product.Name)

	// Save changes
	if result := initializers.DB.Save(&existingProduct); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product updated successfully", "product": existingProduct})
}

func DeleteProduct(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var product models.Product
	if result := initializers.DB.First(&product, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	// Delete the product
	if result := initializers.DB.Delete(&product); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}

func SeeProductsByTypeSlug(c *gin.Context) {
	slug := c.Param("slug")

	var productType models.Product_type
	if result := initializers.DB.Where("slug = ?", slug).First(&productType); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product type not found"})
		return
	}

	var products []models.Product
	if result := initializers.DB.Where("id_product_type = ?", productType.ID).Find(&products); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve products"})
		return
	}

	c.HTML(http.StatusOK, "products_by_type.html", gin.H{
		"ProductType": productType,
		"Products":    products,
	})
}

func SeeProductBySlug(c *gin.Context) {
	slug := c.Param("slug")

	var product models.Product
	if result := initializers.DB.Where("slug = ?", slug).First(&product); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"product": product})
}
