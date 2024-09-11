package controllers

import (
	"myproject/initializers"
	"myproject/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SeeProduct(c *gin.Context) {
	var products []models.Product

	// Query the database to get all products
	result := initializers.DB.Find(&products)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve products",
		})
		return
	}

	// Return the list of products as JSON
	c.JSON(http.StatusOK, gin.H{
		"products": products,
	})
}
