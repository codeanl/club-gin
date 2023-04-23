package api

import (
	"github.com/gin-gonic/gin"
	"school-bbs/service"
)

// 获取标签列表 按使用次数排序
func GetTagList(c *gin.Context) {
	var service service.TagService
	if err := c.ShouldBind(&service); err == nil {
		res := service.GetTagList(c)
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
	}
}

// 创建标签
func TagAdd(c *gin.Context) {
	var service service.TagService
	if err := c.ShouldBind(&service); err == nil {
		res := service.TagAdd(c)
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
	}
}
