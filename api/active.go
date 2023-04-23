package api

import (
	"github.com/gin-gonic/gin"
	"school-bbs/service"
	"school-bbs/util"
	"strconv"
)

// 创建活动
func ActiveAdd(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["file"]
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	var service service.ActiveService
	if err := c.ShouldBind(&service); err == nil {
		res := service.ActiveAdd(c, claims.ID, files)
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
	}
}

// 获取活动列表
func GetActiveList(c *gin.Context) {
	var service service.ActiveService
	if err := c.ShouldBind(&service); err == nil {
		res := service.GetActiveList(c)
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
	}
}

// 获取某个社团的活动列表
func GetActiveListByClub(c *gin.Context) {
	var service service.ActiveService
	if err := c.ShouldBind(&service); err == nil {
		res := service.GetActiveListByClub(c)
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
	}
}

// 删除活动
func ActiveDelete(c *gin.Context) {
	aid, _ := strconv.Atoi(c.PostForm("aid"))
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	var service service.ActiveService
	if err := c.ShouldBind(&service); err == nil {
		res := service.ActiveDelete(c, uint(aid), claims.ID)
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
	}
}

// 更新活动信息
func ActiveUpdate(c *gin.Context) {
	aid, _ := strconv.Atoi(c.PostForm("aid"))
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	var service service.ActiveService
	if err := c.ShouldBind(&service); err == nil {
		res := service.ActiveUpdate(c, uint(aid), claims.ID)
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
	}
}

// 搜索活动
func GetActiveDetail(c *gin.Context) {
	name := c.PostForm("name")
	var service service.ActiveService
	if err := c.ShouldBind(&service); err == nil {
		res := service.GetActiveDetail(c, name)
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
	}
}
