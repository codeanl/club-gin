package model

import (
	"gorm.io/gorm"
)

type Active struct {
	gorm.Model
	ClubID    uint   `gorm:"not null" json:"clubId" form:"clubId"`             //社团id
	Title     string `gorm:"size:1024" json:"title" form:"title"`              // 消息标题
	Content   string `gorm:"type:text;not null" json:"content" form:"content"` // 消息内容
	Cover     string `json:"cover" form:"cover"`                               //封面图
	StartTime string `json:"starttime" form:"starttime"`                       //活动时间
	EndTime   string `json:"endtime" form:"endtime"`                           //结束时间
	Place     string `json:"place" form:"place"`                               //活动地点
	Creator   string `json:"creator" form:"creator"`                           //活动创建人
	MaxPeople string `json:"maxpeaple" form:"maxpeaple"`                       //限制人数
}
