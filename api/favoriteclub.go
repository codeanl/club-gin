package api

import (
	"github.com/gin-gonic/gin"
	"school-bbs/service"
	"school-bbs/util"
)

// 查询我收藏的社团
func GetMyFavoriteClub(c *gin.Context) {
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	var service service.FavoriteClubService
	if err := c.ShouldBind(&service); err == nil {
		res := service.GetMyFavoriteClub(c, claims.ID)
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
	}
}

// 加入收藏
func JoinFavorite(c *gin.Context) {
	var service service.FavoriteClubService
	chaim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&service); err == nil {
		res := service.JoinFavorite(c, int64(chaim.ID))
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
	}
}

// 移除收藏
func QuitFavorite(c *gin.Context) {
	var service service.FavoriteClubService
	chaim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&service); err == nil {
		res := service.QuitFavorite(c, chaim.ID)
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
	}
}
