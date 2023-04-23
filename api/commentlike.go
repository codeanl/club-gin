package api

import (
	"github.com/gin-gonic/gin"
	"school-bbs/service"
	"school-bbs/util"
)

// 点赞
func CommentLike(c *gin.Context) {
	var service service.CommentLikeService
	chaim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&service); err == nil {
		res := service.CommentLike(c, chaim.ID)
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
	}
}

// 取消点赞
func CommentUnLike(c *gin.Context) {
	var service service.CommentLikeService
	chaim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&service); err == nil {
		res := service.CommentUnLike(c, chaim.ID)
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
	}
}
