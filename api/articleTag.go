package api

import (
	"github.com/gin-gonic/gin"
	"school-bbs/service"
)

// 文章添加标签
func ArticleTagAdd(c *gin.Context) {
	var service service.ArticleTagService
	if err := c.ShouldBind(&service); err == nil {
		res := service.ArticleTagAdd(c)
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
	}
}
