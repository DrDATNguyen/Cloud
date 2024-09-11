package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	NameProduct  string
	Descriptions string
}
