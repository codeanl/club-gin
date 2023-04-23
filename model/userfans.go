package model

import "gorm.io/gorm"

type UserFans struct {
	gorm.Model
	NoticerID  uint `json:"noticerId "form:"noticerId"`   //关注者
	FollowerID uint `json:"followerId "form:"followerId"` //被关注者
}
