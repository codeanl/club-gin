package model

import "gorm.io/gorm"

// 标签
type Tag struct {
	gorm.Model
	Name  string `gorm:"size:32;unique;not null" json:"name" form:"name"`
	Usecn uint   `gorm:"default:0" json:"usecn" form:"usecn"`
}
