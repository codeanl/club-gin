package service

import (
	"github.com/gin-gonic/gin"
	"school-bbs/dao"
	"school-bbs/model"
	"school-bbs/serializer"
	"school-bbs/util/e"
)

// 收藏社团
type FavoriteClubService struct {
	UserId int64 `gorm:"index:idx_favorite_user_id;not null" json:"userId" form:"userId"`    // 用户编号
	ClubId int64 `gorm:"index:idx_favorite_article_id;not null" json:"clubId" form:"clubId"` // 收藏社团编号
}

// 加入收藏
func (service FavoriteClubService) JoinFavorite(c *gin.Context, myid int64) serializer.Response {
	code := e.SUCCESS
	// 判断是否已加入收藏
	exist, err := dao.CheckJoinFavoriteExist(myid, service.ClubId)
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
			Data: "您已加入收藏，无需再次添加",
		}
	} else {
		// 数据的插入
		favorite := &model.FavoriteClub{
			UserId: myid,
			ClubId: service.ClubId,
		}
		err = dao.CreateFavorite(favorite)
		if err != nil {
			code = e.ERROR
			return serializer.Response{
				Code: code,
				Msg:  e.GetMsg(code),
			}
		}
		//user收藏个数+1
		myinfo, _ := dao.GetUserById(uint(myid))
		myinfo.FavouriteCnt = myinfo.FavouriteCnt + 1
		err := dao.UpdateUserById(uint(myid), myinfo)
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
			Data: "加入收藏成功",
		}
	}
}

// 移除收藏
func (service FavoriteClubService) QuitFavorite(c *gin.Context, myid uint) serializer.Response {
	code := e.SUCCESS
	// 判断是否已加入收藏
	exist, err := dao.CheckJoinClubExist(myid, uint(service.ClubId))
	if err != nil {
		code = e.ERROR
		return serializer.Response{
			Code: code,
			Msg:  e.GetMsg(code),
		}
	}
	if !exist {
		err := dao.QuidFavorite(int64(myid), service.ClubId)
		if err != nil {
			code := e.ERROR
			return serializer.Response{
				Code: code,
				Msg:  e.GetMsg(code),
			}
		}
		myinfo, _ := dao.GetUserById(myid)
		myinfo.FavouriteCnt = myinfo.FavouriteCnt - 1
		err = dao.UpdateUserById(myid, myinfo)
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
			Data: "取消收藏成功",
		}
	} else {
		return serializer.Response{
			Code: code,
			Msg:  e.GetMsg(code),
			Data: "您未加入收藏，不能取消收藏",
		}
	}

}

// 查询我收藏的社团
func (service FavoriteClubService) GetMyFavoriteClub(c *gin.Context, myid uint) serializer.Response {
	code := e.SUCCESS
	clubs, count, _ := dao.GetMyFavoriteClub(myid)

	return serializer.Response{
		Code: code,
		Msg:  e.GetMsg(code),
		Data: map[string]interface{}{
			"list":  clubs,
			"count": count,
		},
	}
}
