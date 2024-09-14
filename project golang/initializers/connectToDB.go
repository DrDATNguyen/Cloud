package initializers

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var JwtKey = []byte(os.Getenv("Cloud_Secret_key")) // Lấy jwtKey trong file .env

func ConnectToDB() {
	var err error
	dsn := os.Getenv("db") // Lấy dsn của database trong file .env
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failled to connect to DB")
	}
}
