package initializers

import "myproject/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.Admins{})
	DB.AutoMigrate(&models.Product{})
	DB.AutoMigrate(&models.ProductsPackage{})
	DB.AutoMigrate(&models.User{})
}
