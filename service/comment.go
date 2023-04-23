package service

import (
	"github.com/gin-gonic/gin"
	"school-bbs/dao"
	"school-bbs/model"
	"school-bbs/serializer"
	"school-bbs/util/e"
)

// 评论
type CommentService struct {
	UserId       int64  ` json:"userId" form:"userId"`                                      // 用户编号
	EntityId     int64  ` json:"entityId" form:"entityId"`                                  // 被评论文章编号
	Content      string `gorm:"type:text;not null" json:"content" form:"content"`           // 内容
	QuoteId      int64  `gorm:"default:0;not null"  json:"quoteId" form:"quoteId"`          // 引用的评论编号
	LikeCount    int64  `gorm:"not null;default:0" json:"likeCount" form:"likeCount"`       // 点赞数量
	CommentCount int64  `gorm:"not null;default:0" json:"commentCount" form:"commentCount"` // 评论数量
}

// 获取评论列表
func (service CommentService) GetCommentList(c *gin.Context) serializer.Response {
	code := e.SUCCESS
	list, count, err := dao.GetCommentList(service.EntityId)
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

// 发表评论
func (service CommentService) CommentAdd(c *gin.Context, id uint) serializer.Response {
	code := e.SUCCESS
	if service.Content == "" {
		code = e.InvalidParams
		return serializer.Response{
			Code: code,
			Msg:  e.GetMsg(code),
		}
	}
	// 数据的插入
	comment := &model.Comment{
		UserId:   int64(id),
		EntityId: service.EntityId,
		Content:  service.Content,
		QuoteId:  service.QuoteId,
	}
	err := dao.CreateComment(comment)
	if err != nil {
		code = e.ERROR
		return serializer.Response{
			Code: code,
			Msg:  e.GetMsg(code),
		}
	}
	return serializer.Response{
		Code: code,
		Data: serializer.BuildComment(comment),
		Msg:  e.GetMsg(code),
	}
}

// 删除评论
func (service CommentService) CommentDelete(c *gin.Context, myid int64, cid uint) serializer.Response {
	code := e.SUCCESS
	//通过评论uid查询具体信息
	comment, err := dao.GetCommentById(cid)
	if err != nil {
		code = e.ERROR
		return serializer.Response{
			Code:  code,
			Msg:   e.GetMsg(code),
			Error: err.Error(),
		}
	}
	if comment.UserId != myid {
		code = e.ERROR
		return serializer.Response{
			Code: code,
			Msg:  e.GetMsg(code),
			Data: "这不是你的评论",
		}
	} else {
		err = dao.CommentDelete(cid)
		if err != nil {
			code = e.ERROR
			return serializer.Response{
				Code:  code,
				Msg:   e.GetMsg(code),
				Error: err.Error(),
			}
		}
	}

	return serializer.Response{
		Code: code,
		Msg:  e.GetMsg(code),
	}
}
