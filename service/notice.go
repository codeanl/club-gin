package service

import (
	"github.com/gin-gonic/gin"
	"school-bbs/dao"
	"school-bbs/model"
	"school-bbs/serializer"
	"school-bbs/util/e"
)

type MessageService struct {
	FromId       int64  `gorm:"not null" json:"fromId" form:"fromId"`                                      // 消息发送人
	FromType     int64  `gorm:"not null" json:"formType" form:"formType"`                                  // 消息类型 1-管理员通知 2-社团管理员通知
	UserId       int64  `gorm:"default:0;not null;index:idx_message_user_id;" json:"userId" form:"userId"` // 用户编号(消息接收人) 0为全部成员的通知
	Title        string `gorm:"size:1024" json:"title" form:"title"`                                       // 消息标题
	Content      string `gorm:"type:text;not null" json:"content" form:"content"`                          // 消息内容
	QuoteContent string `gorm:"type:text" json:"quoteContent" form:"quoteContent"`                         // 引用内容
	ClubId       uint   `gorm:"not null" json:"clubId" form:"clubId"`                                      //所属社团
}

// 获取某个社团消息列表
func (service MessageService) GetClubMessageList(c *gin.Context) serializer.Response {
	code := e.SUCCESS
	list, count, err := dao.GetMessageClub(service.ClubId)
	if err != nil {
		code = e.ERROR
		return serializer.Response{
			Code: code,
			Msg:  e.GetMsg(code),
		}
	}
	return serializer.Response{
		Code: code,
		Msg:  e.GetMsg(code),
		Data: map[string]interface{}{
			"list":  list,
			"count": count,
		},
	}
}

// 创建消息
func (service MessageService) MessageAdd(c *gin.Context, myid uint) serializer.Response {
	code := e.SUCCESS
	if service.Title == "" || service.Content == "" || &service.ClubId == nil {
		code = e.InvalidParams
		return serializer.Response{
			Code: code,
			Msg:  e.GetMsg(code),
		}
	}
	clubform, _ := dao.GetClubById(service.ClubId)
	if clubform.PresidentID == myid {
		// 数据的插入
		message := &model.Notice{
			FromId:       int64(myid),
			FromType:     2,
			UserId:       service.UserId,
			Title:        service.Title,
			Content:      service.Content,
			QuoteContent: service.QuoteContent,
			ClubId:       service.ClubId,
		}
		err := dao.CreateMessage(message)
		if err != nil {
			code = e.ERROR
			return serializer.Response{
				Code: code,
				Msg:  e.GetMsg(code),
			}
		}
		return serializer.Response{
			Code: code,
			Data: message,
			Msg:  e.GetMsg(code),
		}
	} else {
		return serializer.Response{
			Code: code,
			Msg:  e.GetMsg(code),
			Data: "你不是该社团管理员",
		}
	}

}

// 删除消息
func (service MessageService) MessageDelete(c *gin.Context, nid uint, myid uint) serializer.Response {
	code := e.SUCCESS
	notice, _ := dao.GetMessageById(nid)
	if notice.FromId == int64(myid) {
		err := dao.MessageDelete(nid)
		if err != nil {
			code = e.ERROR
			return serializer.Response{
				Code:  code,
				Msg:   e.GetMsg(code),
				Error: err.Error(),
			}
		}
		return serializer.Response{
			Code: code,
			Msg:  e.GetMsg(code),
		}
	} else {
		return serializer.Response{
			Code: code,
			Msg:  e.GetMsg(code),
			Data: "你不是该社团管理员",
		}
	}
}

// 修改消息信息
func (service MessageService) MessageUpdate(c *gin.Context, myid uint, nid uint) serializer.Response {
	var err error
	code := e.SUCCESS
	message, _ := dao.GetMessageById(nid)
	if service.Title != "" {
		message.Title = service.Title
	}
	if service.Content != "" {
		message.Content = service.Content
	}
	if service.QuoteContent != "" {
		message.QuoteContent = service.QuoteContent
	}
	if message.FromId == int64(myid) {
		err = dao.UpdateMessageById(nid, message)
		if err != nil {
			code = e.ERROR
			return serializer.Response{
				Code:  code,
				Msg:   e.GetMsg(code),
				Error: err.Error(),
			}
		}
		return serializer.Response{
			Code: code,
			Data: serializer.BuildMessage(message),
			Msg:  e.GetMsg(code),
		}
	} else {
		return serializer.Response{
			Code: code,
			Msg:  e.GetMsg(code),
			Data: "你不是该社团管理员",
		}
	}
}

// 查询我的所有消息
func (service MessageService) GetMyMessage(c *gin.Context, myid uint) serializer.Response {
	code := e.SUCCESS
	var allMessages []*model.Notice
	clubs, _, _ := dao.GetMyClubInfo(myid)
	for _, club := range clubs {
		messages, _, _ := dao.GetMessageClub(club.ClubID)
		allMessages = append(allMessages, messages...)
	}

	return serializer.Response{
		Code: code,
		Msg:  e.GetMsg(code),
		Data: allMessages,
	}
}
