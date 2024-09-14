package main

import (
	"myproject/controllers"
	"myproject/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()
	r.POST("/Admin/CreateAdmin", controllers.AddAdmin)
	r.POST("/Admin/Login", controllers.Login)
	r.GET("/Admin/Product", controllers.SeeProduct)
	r.GET("/Admin/ProductPackage/:id", controllers.GetProductPackageByProductID)
	r.POST("/Admin/AddUser", controllers.AddUser)
	r.GET("/Admin/GetAllUser", controllers.GetAllUsers)
	r.GET("/Admin/UserInfo/:id", controllers.GetUserInfoByID)
	r.GET("/Admin/ProductPackageInfo/:id", controllers.GetProductPackageInfo)
	r.GET("/Admin/GetAllAdmin", controllers.GetAllAdmins)
	r.POST("/Admin/AddProduct", controllers.AddProduct)
	r.POST("/Admin/UpdateProduct/:id", controllers.UpdateProduct)
	r.POST("/Admin/DeleteProduct/:id", controllers.DeleteProduct)
	r.Run()
}
