package api

import (
	"github.com/gin-gonic/gin"
	"school-bbs/service"
	"school-bbs/util"
	"strconv"
)

// 获取社团列表和总数
func GetClubList(c *gin.Context) {
	var service service.ClubService
	if err := c.ShouldBind(&service); err == nil {
		res := service.GetClubList(c)
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
	}
}

// 创建社团
func ClubAdd(c *gin.Context) {
	logo, fileHeader, _ := c.Request.FormFile("logo")
	fileSize := fileHeader.Size
	form, _ := c.MultipartForm()
	files := form.File["file"]
	var service service.ClubService
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&service); err == nil {
		res := service.ClubAdd(c, claims.ID, logo, fileSize, files)
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
	}
}

// 查询社长所创建的社团 (我创建的社团)
func UserClubInfo(c *gin.Context) {
	var service service.ClubService
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&service); err == nil {
		res := service.UserClubInfo(c, claims.ID)
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
	}
}

// 社团注销
func ClubDelete(c *gin.Context) {
	cid, _ := strconv.Atoi(c.PostForm("cid"))
	var service service.ClubService
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&service); err == nil {
		res := service.ClubDelete(c, claims.ID, uint(cid))
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
	}
}

// 更新社团信息
func ClubUpdate(c *gin.Context) {
	id, _ := strconv.Atoi(c.PostForm("id"))
	var service service.ClubService
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&service); err == nil {
		res := service.Update(c, uint(id), claims.ID)
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
	}
}

// 查询社团
func GetClubInfo(c *gin.Context) {
	var service service.ClubService
	if err := c.ShouldBind(&service); err == nil {
		res := service.GetClubInfo(c)
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
	}
}

// 更新头像
func UploadClubAvatar(c *gin.Context) {
	file, fileHeader, _ := c.Request.FormFile("file")
	fileSize := fileHeader.Size
	id, _ := strconv.Atoi(c.PostForm("id"))
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	var service service.ClubService
	if err := c.ShouldBind(&service); err == nil {
		res := service.UploadClubAvatar(c, uint(id), file, fileSize, claims.ID)
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
	}
}

// 转让社长
func TransferClub(c *gin.Context) {
	cid, _ := strconv.Atoi(c.PostForm("cid")) //社团id
	uid, _ := strconv.Atoi(c.PostForm("uid")) //被转让人的id
	var service service.ClubService
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&service); err == nil {
		res := service.TransferClub(c, uint(cid), uint(uid), claims.ID)
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
	}
}

// 我加入的社团
func GetMyClubInfo(c *gin.Context) {
	var service service.ClubService
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&service); err == nil {
		res := service.GetMyClubInfo(c, claims.ID)
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
	}
}
