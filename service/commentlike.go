package service

import (
	"github.com/gin-gonic/gin"
	"school-bbs/dao"
	"school-bbs/model"
	"school-bbs/serializer"
	"school-bbs/util/e"
)

type CommentLikeService struct {
	CommentID uint `gorm:"not null" json:"commentid"form:"commentid"`
	UserID    uint
}

// 点赞
func (service CommentLikeService) CommentLike(c *gin.Context, myid uint) serializer.Response {
	code := e.SUCCESS

	// 判断是否已点赞
	exist, err := dao.CheckCommentLikeExist(service.CommentID, myid)
	if err != nil {
		code = e.ERROR
		return serializer.Response{
			Code: code,
			Msg:  e.GetMsg(code),
		}
	}
	if exist {
		code = e.ERROR
		return serializer.Response{
			Code: code,
			Msg:  e.GetMsg(code),
			Data: "您已点赞，无需重复点赞",
		}
	} else {
		// 数据的插入
		clubmembers := &model.CommentLike{
			CommentID: service.CommentID,
			UserID:    myid,
		}
		err = dao.CreateCommengLike(clubmembers)
		if err != nil {
			code = e.ERROR
			return serializer.Response{
				Code: code,
				Msg:  e.GetMsg(code),
			}
		}
		//todo comment表的like+1
		comment, _ := dao.GetCommentById(service.CommentID)
		comment.LikeCount = comment.LikeCount + 1
		err := dao.UpdateCommentById(service.CommentID, comment)
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
			Data: "点赞成功",
		}
	}
}

// 取消点赞
func (service CommentLikeService) CommentUnLike(c *gin.Context, myid uint) serializer.Response {
	code := e.SUCCESS
	// 判断是否已点赞
	exist, err := dao.CheckCommentLikeExist(service.CommentID, myid)
	if err != nil {
		code = e.ERROR
		return serializer.Response{
			Code: code,
			Msg:  e.GetMsg(code),
		}
	}
	if !exist {
		code = e.ERROR
		return serializer.Response{
			Code: code,
			Msg:  e.GetMsg(code),
			Data: "您已取消点赞，无需重复取消点赞",
		}
	} else {
		err = dao.UnLike(service.CommentID, myid)
		if err != nil {
			code = e.ERROR
			return serializer.Response{
				Code: code,
				Msg:  e.GetMsg(code),
			}
		}
		// comment表的like-1
		comment, _ := dao.GetCommentById(service.CommentID)
		comment.LikeCount = comment.LikeCount - 1
		err := dao.UpdateCommentById(service.CommentID, comment)
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
			Data: "取消点赞成功",
		}
	}
}
