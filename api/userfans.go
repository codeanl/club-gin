package api

import (
	"github.com/gin-gonic/gin"
	"school-bbs/service"
	"school-bbs/util"
)

// 关注
func UserLike(c *gin.Context) {
	var service service.UserFAnsService
	chaim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&service); err == nil {
		res := service.UserLike(c, chaim.ID)
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
	}
}

// 取消关注
func UserUnLike(c *gin.Context) {
	var service service.UserFAnsService
	chaim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&service); err == nil {
		res := service.UserUnLike(c, chaim.ID)
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
	}
}
