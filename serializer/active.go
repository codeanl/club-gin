package serializer

import (
	"school-bbs/model"
)

type Active struct {
	ClubID    uint   `gorm:"not null" json:"clubId" form:"clubId"`
	Title     string `gorm:"size:1024" json:"title" form:"title"`              // 消息标题
	Content   string `gorm:"type:text;not null" json:"content" form:"content"` // 消息内容
	Cover     string `json:"cover" form:"cover"`
	StartTime string `json:"starttime" form:"starttime"` //活动时间
	EndTime   string `json:"endtime" form:"endtime"`
	Place     string `json:"place" form:"place"`     //活动地点
	Creator   string `json:"creator" form:"creator"` //活动创建人
	MaxPeople string `json:"maxpeaple" form:"maxpeaple"`
}

// BuildActive 序列化活动
func BuildActive(active *model.Active) Active {
	return Active{
		ClubID:    active.ClubID,
		Title:     active.Title,
		Content:   active.Content,
		Cover:     active.Cover,
		StartTime: active.StartTime,
		EndTime:   active.EndTime,
		Place:     active.Place,
		Creator:   active.Creator,
		MaxPeople: active.MaxPeople,
	}
}
