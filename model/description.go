package model

import "github.com/jinzhu/gorm"

// Description struct
type Description struct {
	gorm.Model
	District string `gorm:"not null" json:"district"`
	LSGD     string `gorm:"not null" json:"lsgd"`
	Data     string `gorm:"not null" json:"data"`
}
