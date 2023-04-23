package serializer

import (
	"school-bbs/model"
)

type Club struct {
	ID          uint
	Name        string `gorm:"not null" form:"name;" json:"name"` //社团名字
	PresidentID uint   //会长id
	Phone       string `gorm:"" form:"phone;" json:"phone"`    //联系方式
	Desc        string `gorm:"" json:"desc" form:"desc"`       //简介
	Purpose     string `gorm:"" json:"purpose" form:"purpose"` //宗旨
	Avatar      string `gorm:""  json:"avatar" form:"avatar"`  //社团头像
	PeoPleCln   int    `gorm:"default:0"  json:"peoPleCln" form:"peoPleCln"`
}

// BuildClub 序列化社团
func BuildClub(club *model.Club) Club {
	return Club{
		ID:          club.ID,
		Name:        club.Name,
		PresidentID: club.PresidentID,
		Phone:       club.Phone,
		Desc:        club.Desc,
		Purpose:     club.Purpose,
		Avatar:      club.Avatar,
		PeoPleCln:   club.PeoPleCln,
	}
}
