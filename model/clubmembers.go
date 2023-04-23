package model

import "gorm.io/gorm"

type ClubMembers struct {
	gorm.Model
	ClubID uint `gorm:"default:0" json:"clubID,omitempty" form:"clubId"`
	UserID uint `gorm:"default:0" json:"userID,omitempty" form:"userId"`
}
