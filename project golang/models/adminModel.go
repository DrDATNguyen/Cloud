package models

import "time"

type Admins struct {
	ID          uint   `gorm:"primaryKey"`
	Username    string `gorm:"type:varchar(255);not null"`
	Password    string `gorm:"type:varchar(255);not null"`
	Email       string
	PhoneNumber string
	Address     string
	VipAdmin    string
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}
