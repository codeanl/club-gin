package model

import "gorm.io/gorm"

type ArticleImg struct {
	gorm.Model
	ArticleID uint `gorm:"not null"`
	ImgPath   string
}
