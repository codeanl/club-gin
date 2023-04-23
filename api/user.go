package api

import (
	"github.com/gin-gonic/gin"
	"school-bbs/service"
	"school-bbs/util"
)

// 用户注册
func Register(c *gin.Context) {
	var service service.UserService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Register(c)
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
	}
}

// 用户登录
func Login(c *gin.Context) {
	var service service.UserService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Login(c)
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
	}
}

// 发送验证码
func SendCode(c *gin.Context) {
	var service service.UserService
	if err := c.ShouldBind(&service); err == nil {
		res := service.SendCode(c)
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
	}
}

// 更新头像
func UploadAvatar(c *gin.Context) {
	file, fileHeader, _ := c.Request.FormFile("file")
	fileSize := fileHeader.Size
	var service service.UserService
	chaim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&service); err == nil {
		res := service.UploadAvatar(c, chaim.ID, file, fileSize)
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
	}
}

// 用户修改密码
func UpdatePassword(c *gin.Context) {
	oldpassword := c.PostForm("oldpassword")
	newpassword := c.PostForm("newpassword")
	var service service.UserService
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&service); err == nil {
		res := service.UpdatePassword(c, claims.ID, oldpassword, newpassword)
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
	}
}

// 用户信息修改
func UserUpdate(c *gin.Context) {
	var service service.UserService
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&service); err == nil {
		res := service.Update(c, claims.ID)
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
	}
}

// 获取用户列表
func GetUserList(c *gin.Context) {
	var service service.UserService
	if err := c.ShouldBind(&service); err == nil {
		res := service.GetUserList(c)
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
	}
}

// 用户注销
func UserMyDelete(c *gin.Context) {
	var service service.UserService
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&service); err == nil {
		res := service.UserMyDelete(c, claims.ID)
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
	}
}

// 获取自己的详细信息
func GetMyInfo(c *gin.Context) {
	var service service.UserService
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&service); err == nil {
		res := service.GetMyInfo(c, claims.ID)
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
	}
}

// 查询用户
func GetUserDetail(c *gin.Context) {
	var service service.UserService
	if err := c.ShouldBind(&service); err == nil {
		res := service.GetUserDetail(c)
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
	}
}

// 找回账号密码
func Retrieve(c *gin.Context) {
	var service service.UserService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Retrieve(c)
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
	}
}
