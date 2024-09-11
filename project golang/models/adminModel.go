package models

import "time"

type Admins struct {
	ID          uint `gorm:"primaryKey"`
	Username    string
	Password    string
	Email       string
	PhoneNumber string
	Address     string
	VipAdmin    string
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}
