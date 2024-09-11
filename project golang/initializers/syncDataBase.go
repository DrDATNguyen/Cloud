package initializers

import "myproject/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.Admins{})
	DB.AutoMigrate(&models.Product{})
}
