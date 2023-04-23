package model

import (
	"gorm.io/gorm"
)

type Club struct {
	gorm.Model
	Name        string `gorm:"" json:"name" form:"name"`       //名字
	Phone       string `gorm:"" json:"phone" form:"phone"`     //联系方式
	Desc        string `gorm:"" json:"desc" form:"desc"`       //简介
	Purpose     string `gorm:"" json:"purpose" form:"purpose"` //宗旨
	Avatar      string `gorm:""  json:"avatar" form:"avatar"`  //社团头像
	PresidentID uint   //会长id
	PeoPleCln   int    `gorm:"default:0"  json:"peoPleCln" form:"peoPleCln"`
}
