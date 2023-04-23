package api

import (
	"github.com/gin-gonic/gin"
	"school-bbs/service"
	"school-bbs/util"
	"strconv"
)

// 获取文章列表
func GetArticleList(c *gin.Context) {
	var service service.ArticleService
	if err := c.ShouldBind(&service); err == nil {
		res := service.GetArticleList(c)
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
	}
}

// 创建文章
func ArticleAdd(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["file"]
	var service service.ArticleService
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&service); err == nil {
		res := service.ArticleAdd(c, claims.ID, files)
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
	}
}

// 删除文章
func ArticleDelete(c *gin.Context) {
	aid, _ := strconv.Atoi(c.PostForm("aid"))
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	var service service.ArticleService
	if err := c.ShouldBind(&service); err == nil {
		res := service.ArticleDelete(c, uint(aid), claims.ID)
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
	}
}

// 更新文章信息
func ArticleUpdate(c *gin.Context) {
	aid, _ := strconv.Atoi(c.PostForm("aid"))
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))

	var service service.ArticleService
	if err := c.ShouldBind(&service); err == nil {
		res := service.ArticleUpdate(c, uint(aid), claims.ID)
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
	}
}

// 搜索文章
func GetArticleDetail(c *gin.Context) {
	name := c.PostForm("name")
	var service service.ArticleService
	if err := c.ShouldBind(&service); err == nil {
		res := service.GetArticleDetail(c, name)
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
	}
}

// 查询我的文章
func GetMyArticle(c *gin.Context) {
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	var service service.ArticleService
	if err := c.ShouldBind(&service); err == nil {
		res := service.GetMyArticle(c, claims.ID)
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
	}
}
