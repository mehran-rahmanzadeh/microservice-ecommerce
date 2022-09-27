package model

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Title       string
	Description string
	Image       string
	Children    []Category `gorm:"foreignkey:ParentID"`
	ParentID    *uint
	Depth       int `gorm:"default:0"`
}

// BeforeCreate hook to incrementation of depth field in hierarchy
func (c *Category) BeforeCreate(tx *gorm.DB) (err error) {
	if c.ParentID == nil {
		return
	}
	var parent Category
	tx.First(&parent, "id = ?", c.ParentID)
	c.Depth = parent.Depth + 1
	return
}
