package domain

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Code  string
	Price uint
}
