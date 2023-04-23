package api

import (
	"github.com/gin-gonic/gin"
	"school-bbs/service"
	"school-bbs/util"
	"strconv"
)

// 获取某个文章的评论和次数 按like_count点赞次数排序
func GetCommentList(c *gin.Context) {
	var service service.CommentService
	if err := c.ShouldBind(&service); err == nil {
		res := service.GetCommentList(c)
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
	}
}

// 发表评论
func CommentAdd(c *gin.Context) {
	var service service.CommentService
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&service); err == nil {
		res := service.CommentAdd(c, claims.ID)
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
	}
}

// 删除评论
func CommentDelete(c *gin.Context) {
	cid, _ := strconv.Atoi(c.PostForm("cid"))
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	var service service.CommentService
	if err := c.ShouldBind(&service); err == nil {
		res := service.CommentDelete(c, int64(claims.ID), uint(cid))
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
	}
}
