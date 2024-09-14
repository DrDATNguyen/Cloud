package models

import "gorm.io/gorm"

type ProductsPackage struct {
	gorm.Model
	NameProduct string  `gorm:"type:varchar(255);not null"`
	RAM         *string `gorm:"type:varchar(255);"`
	CPU         *string `gorm:"type:varchar(255);"`
	Storage     *string `gorm:"type:varchar(255);"`
	Price       float64 `gorm:"type:float;not null"`
	ProductID   int     `gorm:"not null"`
	Product     Product `gorm:"foreignKey:ProductID;references:ID"`
}
