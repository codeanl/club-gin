package api

import (
	"github.com/gin-gonic/gin"
	"school-bbs/service"
	"school-bbs/util"
	"strconv"
)

// 获取某个社团消息列表(社团通知)
func GetClubMessageList(c *gin.Context) {
	var service service.MessageService
	if err := c.ShouldBind(&service); err == nil {
		res := service.GetClubMessageList(c)
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
	}
}

// 发布通知
func MessageAdd(c *gin.Context) {
	var service service.MessageService
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&service); err == nil {
		res := service.MessageAdd(c, claims.ID)
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
	}
}

// 删除消息
func MessageDelete(c *gin.Context) {
	nid, _ := strconv.Atoi(c.PostForm("nid"))
	var service service.MessageService
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))

	if err := c.ShouldBind(&service); err == nil {
		res := service.MessageDelete(c, uint(nid), claims.ID)
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
	}
}

// 更新通知
func MessageUpdate(c *gin.Context) {
	nid, _ := strconv.Atoi(c.PostForm("nid"))
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))

	var service service.MessageService
	if err := c.ShouldBind(&service); err == nil {
		res := service.MessageUpdate(c, claims.ID, uint(nid))
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
	}
}

// 查询我的社团消息
func GetMyMessage(c *gin.Context) {
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	var service service.MessageService
	if err := c.ShouldBind(&service); err == nil {
		res := service.GetMyMessage(c, claims.ID)
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
	}
}
