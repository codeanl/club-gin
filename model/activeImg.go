package model

import (
	"gorm.io/gorm"
)

type ActiveImg struct {
	gorm.Model
	ActiveID uint `gorm:"not null"`
	ImgPath  string
}
