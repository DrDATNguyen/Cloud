package controllers

import (
	"myproject/initializers"
	"myproject/models"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func generateSlug(name string) string {
	// Convert to lowercase and replace spaces with dashes
	return strings.ToLower(strings.ReplaceAll(name, " ", "-"))
}

func SeeProductType(c *gin.Context) {
	// Create a slice to hold the product types
	var productTypes []models.Product_type

	// Query the database to find all product types
	if result := initializers.DB.Find(&productTypes); result.Error != nil {
		// Return an error if the query fails
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve product types",
		})
		return
	}

	// Return the list of product types as JSON
	c.HTML(http.StatusOK, "productType.html", gin.H{
		"ProductTypes": productTypes,
	})
}

func AddProductType(c *gin.Context) {
	var productType models.Product_type

	// Bind incoming JSON to productType struct
	if err := c.Bind(&productType); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request format",
		})
		return
	}

	// Validate required fields
	if productType.NameProductType == "" || productType.Descriptions == "" || productType.Content == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "NameProductType, Descriptions, and Content cannot be empty",
		})
		return
	}

	// Generate thumb and slug based on NameProductType
	productType.Slug = generateSlug(productType.NameProductType)
	productType.Thumb = "images/thumbs/" + generateSlug(productType.NameProductType) + ".jpg"

	// Insert the new product type into the database
	if result := initializers.DB.Create(&productType); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error adding the product type",
		})
		return
	}

	// Return success response
	c.JSON(http.StatusOK, gin.H{
		"message":     "Product type added successfully",
		"productType": productType,
	})
}

func UpdateProductType(c *gin.Context) {
	// Get the ID from the URL parameter
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID format",
		})
		return
	}

	// Bind JSON to productType struct
	var productType models.Product_type
	if err := c.BindJSON(&productType); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request format",
		})
		return
	}

	// Validate required fields
	if productType.NameProductType == "" || productType.Descriptions == "" || productType.Content == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "NameProductType, Descriptions, and Content cannot be empty",
		})
		return
	}

	// Find the existing product type by ID
	var existingProductType models.Product_type
	if result := initializers.DB.First(&existingProductType, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Product type not found",
		})
		return
	}

	// Update the fields
	existingProductType.NameProductType = productType.NameProductType
	existingProductType.Descriptions = productType.Descriptions
	existingProductType.Content = productType.Content
	existingProductType.Slug = generateSlug(productType.NameProductType)
	existingProductType.Thumb = "images/thumbs/" + generateSlug(productType.NameProductType) + ".jpg"

	// Save the changes
	if result := initializers.DB.Save(&existingProductType); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error updating the product type",
		})
		return
	}

	// Return success response
	c.JSON(http.StatusOK, gin.H{
		"message":     "Product type updated successfully",
		"productType": existingProductType,
	})
}

// Delete an existing Product_type
func DeleteProductType(c *gin.Context) {
	// Get the ID from the URL parameter
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID format",
		})
		return
	}

	// Find the product type by ID
	var productType models.Product_type
	if result := initializers.DB.First(&productType, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Product type not found",
		})
		return
	}

	// Delete the product type
	if result := initializers.DB.Delete(&productType, id); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error deleting the product type",
		})
		return
	}

	// Return success response
	c.JSON(http.StatusOK, gin.H{
		"message": "Product type deleted successfully",
	})
}
