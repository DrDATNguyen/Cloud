package main

import (
	"myproject/controllers"
	"myproject/initializers"
	"myproject/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("view/admin/*.html")
	r.Static("/static", "./view/admin")

	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/login")
	})

	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "sign-in.html", gin.H{})
	})

	r.GET("/dashboard", func(c *gin.Context) {
		c.HTML(http.StatusOK, "dashboard.html", gin.H{
			"title": "Main website",
		})
	})

	r.GET("/product-type", func(c *gin.Context) {
		var productTypes []models.Product_type

		// Fetch all product types from the database
		if err := initializers.DB.Find(&productTypes).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to retrieve product types",
			})
			return
		}

		// Pass data to HTML template
		c.HTML(http.StatusOK, "productType.html", gin.H{
			"ProductTypes": productTypes,
		})
	})

	r.POST("/Admin/CreateAdmin", controllers.AddAdmin)
	r.POST("/Admin/Login", controllers.Login)
	r.GET("/Admin/ProductType", controllers.SeeProductType)
	r.GET("/Admin/Product", controllers.SeeProduct)
	r.GET("/Admin/Product/:slug", controllers.SeeProductsByTypeSlug)
	r.POST("/Admin/AddProduct", controllers.AddProduct)
	r.POST("/Admin/UpdateProduct/:id", controllers.UpdateProduct)
	r.POST("/Admin/AddUser", controllers.AddUser)
	r.GET("/Admin/GetAllUser", controllers.GetAllUsers)
	r.GET("/Admin/UserInfo/:id", controllers.GetUserInfoByID)
	r.GET("/Admin/GetAllAdmin", controllers.GetAllAdmins)
	r.POST("/Admin/AddProductType", controllers.AddProductType)
	r.POST("/Admin/UpdateProductType/:id", controllers.UpdateProductType)
	r.POST("/Admin/DeleteProductType/:id", controllers.DeleteProductType)

	r.GET("/Admin/ProductPackage/:slug", controllers.SeeProductsPackageByProductSlug)
	r.POST("/Admin/AddProductPackage", controllers.AddProductsPackage)
	r.POST("/Admin/UpdateProductPackage/:id", controllers.UpdateProduct)
	r.POST("/Admin/DeleteProductPackage/:id", controllers.DeleteProduct)
	r.Run()
}
