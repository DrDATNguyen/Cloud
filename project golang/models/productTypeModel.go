package models

import "gorm.io/gorm"

type Product_type struct {
	gorm.Model
	NameProductType string
	Descriptions    string
	Content         string
	Thumb           string
	Slug            string
}
