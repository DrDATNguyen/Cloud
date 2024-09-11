package controllers

import (
	"fmt"
	"myproject/initializers"
	"myproject/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {
	var body struct {
		Username string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}
	fmt.Println(body)

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})
	}

	admin := models.Admins{Username: body.Username, Password: string(hash), Email: "123@gmail.com", Address: "Street 1", VipAdmin: "Vip 1", PhoneNumber: "123"}
	result := initializers.DB.Create(&admin)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create admin",
		})
	}

	c.JSON(http.StatusOK, gin.H{})
}

func Login(c *gin.Context) {
	var body struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}
	fmt.Println(body)

	var admin models.Admins
	result := initializers.DB.First(&admin, "username = ?", body.Username)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Admin not found or query error",
		})
		return
	}
	fmt.Println(admin)

	err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(body.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid username or password",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
	})
}
