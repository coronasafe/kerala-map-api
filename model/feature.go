package model

import "github.com/jinzhu/gorm"

// Feature struct
type Feature struct {
	gorm.Model
	Data        string `gorm:"not null" json:"data"`
	Description string `gorm:"not null" json:"description"`
}
