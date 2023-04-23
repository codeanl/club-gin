package serializer

import "school-bbs/model"

// 消息
type Message struct {
	ID           uint
	FromId       int64  `gorm:"not null" json:"fromId" form:"fromId"`                                      // 消息发送人
	FromType     int64  `gorm:"not null" json:"formType" form:"formType"`                                  // 消息类型 1-管理员通知 2-社团管理员通知
	UserId       int64  `gorm:"default:0;not null;index:idx_message_user_id;" json:"userId" form:"userId"` // 用户编号(消息接收人) 0为全部成员的通知
	Title        string `gorm:"size:1024" json:"title" form:"title"`                                       // 消息标题
	Content      string `gorm:"type:text;not null" json:"content" form:"content"`                          // 消息内容
	QuoteContent string `gorm:"type:text" json:"quoteContent" form:"quoteContent"`                         // 引用内容
	ClubId       uint   `gorm:"not null" json:"clubId" form:"clubId"`                                      //所属社团
}

func BuildMessage(message *model.Notice) Message {
	return Message{
		ID:           message.ID,
		FromId:       message.FromId,
		FromType:     message.FromType,
		UserId:       message.UserId,
		Title:        message.Title,
		Content:      message.Content,
		QuoteContent: message.QuoteContent,
	}
}
