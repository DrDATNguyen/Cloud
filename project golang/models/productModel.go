package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name            string
	Descriptions    string
	Content         string
	Thumb           string
	Slug            string
	ID_product_type int
	ProductType     Product_type `gorm:"foreignKey:ID_product_type;references:ID"`
}
