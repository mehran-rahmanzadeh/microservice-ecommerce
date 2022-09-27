package model

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Title       string
	Description string
	Image       string
	Parent      *Category `gorm:"foreignkey:ParentID"`
	ParentID    *uint
}
