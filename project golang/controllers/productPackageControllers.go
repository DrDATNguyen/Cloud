package controllers

import (
	"myproject/initializers"
	"myproject/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SeeProductsPackage(c *gin.Context) {
	var packages []models.ProductsPackage

	// Retrieve all packages from the database
	if err := initializers.DB.Find(&packages).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve product packages",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"packages": packages,
	})
}

func SeeProductsPackageByProductSlug(c *gin.Context) {
	productSlug := c.Param("slug") // Get the slug from URL parameters

	var product models.Product
	// Find the product by its slug
	if err := initializers.DB.Where("slug = ?", productSlug).First(&product).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Product not found",
		})
		return
	}

	var packages []models.ProductsPackage
	// Find all packages associated with the product's ID
	if err := initializers.DB.Where("product_id = ?", product.ID).Find(&packages).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve product packages",
		})
		return
	}

	c.HTML(http.StatusOK, "productPackage.html", gin.H{
		"product":  product.Name,
		"packages": packages,
	})
}

func AddProductsPackage(c *gin.Context) {
	var productsPackage models.ProductsPackage

	// Bind JSON body to the struct
	if err := c.BindJSON(&productsPackage); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request format",
		})
		return
	}

	// Validate required fields
	if productsPackage.NameProductPackage == "" || productsPackage.Price <= 0 || productsPackage.ProductID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Name, Price, and ProductID fields are required",
		})
		return
	}

	// Add to the database
	if err := initializers.DB.Create(&productsPackage).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to add product package",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Product package added successfully",
		"package": productsPackage,
	})
}

func UpdateProductsPackage(c *gin.Context) {
	// Get the ID from the URL parameter
	id := c.Param("id")

	var productsPackage models.ProductsPackage

	// Find the existing package by ID
	if err := initializers.DB.First(&productsPackage, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Product package not found",
		})
		return
	}

	// Bind the updated data to the struct
	if err := c.BindJSON(&productsPackage); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request format",
		})
		return
	}

	// Validate required fields
	if productsPackage.NameProductPackage == "" || productsPackage.Price <= 0 || productsPackage.ProductID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Name, Price, and ProductID fields are required",
		})
		return
	}

	// Save the updated package
	if err := initializers.DB.Save(&productsPackage).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update product package",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Product package updated successfully",
		"package": productsPackage,
	})
}

func DeleteProductsPackage(c *gin.Context) {
	// Get the ID from the URL parameter
	id := c.Param("id")

	var productsPackage models.ProductsPackage

	// Find the package by ID
	if err := initializers.DB.First(&productsPackage, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Product package not found",
		})
		return
	}

	// Delete the package
	if err := initializers.DB.Delete(&productsPackage).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete product package",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Product package deleted successfully",
	})
}
