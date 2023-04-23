package service

import (
	"github.com/gin-gonic/gin"
	"school-bbs/dao"
	"school-bbs/model"
	"school-bbs/serializer"
	"school-bbs/util/e"
)

type UserFAnsService struct {
	NoticerID  uint `json:"noticerId "form:"noticerId"`   //关注者
	FollowerID uint `json:"followerId "form:"followerId"` //被关注者
}

// 关注
func (service UserFAnsService) UserLike(c *gin.Context, myid uint) serializer.Response {
	code := e.SUCCESS
	// 判断是否已关注
	exist, err := dao.CheckUserLikeExist(myid, service.FollowerID)
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
			Data: "您已关注，无需重复关注",
		}
	} else {
		// 数据的插入
		userfans := &model.UserFans{
			NoticerID:  myid,
			FollowerID: service.FollowerID,
		}
		err = dao.CreateUserLike(userfans)
		if err != nil {
			code = e.ERROR
			return serializer.Response{
				Code: code,
				Msg:  e.GetMsg(code),
			}
		}
		noticer, _ := dao.GetUserById(myid) //关注者
		noticer.AttentionCnt = noticer.AttentionCnt + 1
		err := dao.UpdateUserById(myid, noticer)
		if err != nil {
			code = e.ERROR
			return serializer.Response{
				Code: code,
				Msg:  e.GetMsg(code),
			}
		}
		follower, _ := dao.GetUserById(service.FollowerID) //被关注者
		follower.Fans = follower.Fans + 1
		err = dao.UpdateUserById(service.FollowerID, follower)
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
			Data: "关注成功",
		}
	}
}

// 取消关注
func (service UserFAnsService) UserUnLike(c *gin.Context, myid uint) serializer.Response {
	code := e.SUCCESS
	// 判断是否已关注
	exist, err := dao.CheckUserLikeExist(myid, service.FollowerID)
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
			Data: "您已取消关注，无需重复取消关注",
		}
	} else {
		err = dao.UnUserLike(myid, service.FollowerID)
		if err != nil {
			code = e.ERROR
			return serializer.Response{
				Code: code,
				Msg:  e.GetMsg(code),
			}
		}

		noticer, _ := dao.GetUserById(myid) //关注者
		noticer.AttentionCnt = noticer.AttentionCnt - 1
		err := dao.UpdateUserById(myid, noticer)
		if err != nil {
			code = e.ERROR
			return serializer.Response{
				Code: code,
				Msg:  e.GetMsg(code),
			}
		}
		follower, _ := dao.GetUserById(service.FollowerID) //被关注者
		follower.Fans = follower.Fans - 1
		err = dao.UpdateUserById(service.FollowerID, follower)
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
