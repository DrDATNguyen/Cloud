package models

import "gorm.io/gorm"

type ProductsPackage struct {
	gorm.Model
	NameProductPackage string  `gorm:"type:varchar(255);not null"`
	RAM                string  `gorm:"type:varchar(255);"`
	CPU                string  `gorm:"type:varchar(255);"`
	Storage            string  `gorm:"type:varchar(255);"`
	Price              float64 `gorm:"type:float;not null"`
	ProductID          int     `gorm:"not null"`
	Hourly             bool    `gorm:"default:false"`
	Monthly            bool    `gorm:"default:false"`
	Quarterly          bool    `gorm:"default:false"`
	Biannually         bool    `gorm:"default:false"`
	Annually           bool    `gorm:"default:false"`
	Biennially         bool    `gorm:"default:false"`
	Triennially        bool    `gorm:"default:false"`
	Quinquennially     bool    `gorm:"default:false"`
	Decennially        bool    `gorm:"default:false"`
	Data_stransfer     string  `gorm:"type:varchar(255);"`
	Bandwidth          string  `gorm:"type:varchar(255);"`
	Tax                float64
	Product            Product `gorm:"foreignKey:ProductID;references:ID"`
}
