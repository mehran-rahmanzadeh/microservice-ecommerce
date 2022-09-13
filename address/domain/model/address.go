package model

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	UserID uint
	PostalAddress string
	PostalCode string
}
