package model

import "gorm.io/gorm"

type ClubImg struct {
	gorm.Model
	ClubID  uint `gorm:"not null"`
	ImgPath string
}
