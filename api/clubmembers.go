package api

import (
	"github.com/gin-gonic/gin"
	"school-bbs/service"
	"school-bbs/util"
)

// 学生加入社团
func JoinClub(c *gin.Context) {
	var service service.ClubMemberService
	chaim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&service); err == nil {
		res := service.JoinClub(c, chaim.ID)
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
	}
}

// 学生退出社团
func QuitClub(c *gin.Context) {
	var service service.ClubMemberService
	chaim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&service); err == nil {
		res := service.QuitClub(c, chaim.ID)
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
	}
}
