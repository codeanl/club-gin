package service

import (
	"github.com/gin-gonic/gin"
	"school-bbs/dao"
	"school-bbs/model"
	"school-bbs/serializer"
	"school-bbs/util/e"
)

type ClubMemberService struct {
	ClubID uint `json:"clubID,omitempty" form:"clubId"`
	UserID uint `json:"userID,omitempty" form:"userId"`
}

// 学生加入社团
func (service ClubMemberService) JoinClub(c *gin.Context, id uint) serializer.Response {
	code := e.SUCCESS
	if service.ClubID == 0 {
		code = e.InvalidParams
		return serializer.Response{
			Code: code,
			Msg:  e.GetMsg(code),
		}
	}
	// 判断是否已加入社团
	exist, err := dao.CheckJoinClubExist(id, service.ClubID)
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
			Data: "您已加入该社团",
		}
	} else {
		// 数据的插入
		clubmembers := &model.ClubMembers{
			ClubID: service.ClubID,
			UserID: id,
		}
		clubform, _ := dao.GetClubById(service.ClubID)
		_, exist, _ = dao.CheckClubExist(clubform.Name)
		if exist {
			err = dao.CreateJoinClub(clubmembers)
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
				Data: "加入社团成功",
			}
		} else {
			return serializer.Response{
				Code: code,
				Msg:  e.GetMsg(code),
				Data: "不存在该社团",
			}
		}
	}
}

// 学生退出社团
func (service ClubMemberService) QuitClub(c *gin.Context, id uint) serializer.Response {
	code := e.SUCCESS
	if service.ClubID == 0 {
		code = e.InvalidParams
		return serializer.Response{
			Code: code,
			Msg:  e.GetMsg(code),
		}
	}
	// 判断是否已加入社团
	exist, err := dao.CheckJoinClubExist(id, service.ClubID)
	if err != nil {
		code = e.ERROR
		return serializer.Response{
			Code: code,
			Msg:  e.GetMsg(code),
		}
	}
	if exist {
		err := dao.QuidClub(id, service.ClubID)
		if err != nil {
			code := e.ERROR
			return serializer.Response{
				Code: code,
				Msg:  e.GetMsg(code),
			}
		}
		return serializer.Response{
			Code: code,
			Msg:  e.GetMsg(code),
			Data: "退出成功",
		}
	} else {
		return serializer.Response{
			Code: code,
			Msg:  e.GetMsg(code),
			Data: "您未加入社团，不能退出",
		}
	}

}
