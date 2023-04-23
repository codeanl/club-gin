package api

import (
	"github.com/gin-gonic/gin"
	"school-bbs/service"
	"school-bbs/util"
)

// 点赞
func ArticleLike(c *gin.Context) {
	var service service.ArticleLikeService
	chaim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&service); err == nil {
		res := service.ArticleLike(c, chaim.ID)
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
	}
}

// 取消点赞
func ArticleUnLike(c *gin.Context) {
	var service service.ArticleLikeService
	chaim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&service); err == nil {
		res := service.ArticleUnLike(c, chaim.ID)
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
	}
}
