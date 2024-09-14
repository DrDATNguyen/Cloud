package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email       string `gorm:"unique;not null"`
	UserName    string `gorm:"not null"`
	PhoneNumber string `gorm:"unique;not null"`
	Pass        string `gorm:"not null"`
	Token       string `gorm:"not null"`
}
