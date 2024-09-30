package models

import "gorm.io/gorm"

type Product_image struct {
	gorm.Model
	Name         string
	Descriptions string
	Thumb        string
	Type         int
	Content      string
	Slug         string
	ID_products  int
	// ProductType  Product_type `gorm:"foreignKey:Type;references:ID"`
	Product Product `gorm:"foreignKey:ID_products;references:ID"`
}
