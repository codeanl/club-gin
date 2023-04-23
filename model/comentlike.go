package model

import "gorm.io/gorm"

type CommentLike struct {
	gorm.Model
	CommentID uint `gorm:"not null" json:"commentid"form:"commentid"`
	UserID    uint
}
