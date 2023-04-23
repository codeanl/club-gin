package service

import (
	"github.com/gin-gonic/gin"
	"school-bbs/dao"
	"school-bbs/model"
	"school-bbs/serializer"
	"school-bbs/util/e"
)

// 用户点赞文章
type ArticleLikeService struct {
	UserId    uint `json:"userId" form:"userId"`       // 用户
	ArticleId uint `json:"articleId" form:"articleId"` // 文章编号
}

// 点赞
func (service ArticleLikeService) ArticleLike(c *gin.Context, myid uint) serializer.Response {
	code := e.SUCCESS
	// 判断是否已点赞
	exist, err := dao.CheckArticleLikeExist(service.ArticleId, myid)
	if err != nil {
		code = e.ERROR
		return serializer.Response{
			Code: code,
			Msg:  e.GetMsg(code),
		}
	}
	if exist {
		code = e.ERROR
		return serializer.Response{
			Code: code,
			Msg:  e.GetMsg(code),
			Data: "您已点赞，无需重复点赞",
		}
	} else {
		// 数据的插入
		articlelike := &model.ArticleLike{
			UserId:    myid,
			ArticleId: service.ArticleId,
		}
		err = dao.CreateArticleLike(articlelike)
		if err != nil {
			code = e.ERROR
			return serializer.Response{
				Code: code,
				Msg:  e.GetMsg(code),
			}
		}
		article, _ := dao.GetArticleById(service.ArticleId)
		article.LikeCount = article.LikeCount + 1
		err := dao.UpdateArticleById(service.ArticleId, article)
		if err != nil {
			code = e.ERROR
			return serializer.Response{
				Code: code,
				Msg:  e.GetMsg(code),
			}
		}
		return serializer.Response{
			Code: code,
			Msg:  e.GetMsg(code),
			Data: "点赞成功",
		}
	}
}

// 取消点赞
func (service ArticleLikeService) ArticleUnLike(c *gin.Context, myid uint) serializer.Response {
	code := e.SUCCESS
	// 判断是否已点赞
	exist, err := dao.CheckArticleLikeExist(service.ArticleId, myid)
	if err != nil {
		code = e.ERROR
		return serializer.Response{
			Code: code,
			Msg:  e.GetMsg(code),
		}
	}
	if !exist {
		code = e.ERROR
		return serializer.Response{
			Code: code,
			Msg:  e.GetMsg(code),
			Data: "您已取消点赞，无需重复取消点赞",
		}
	} else {
		err = dao.ArticleUnLike(service.ArticleId, myid)
		if err != nil {
			code = e.ERROR
			return serializer.Response{
				Code: code,
				Msg:  e.GetMsg(code),
			}
		}
		// comment表的like-1
		article, _ := dao.GetArticleById(service.ArticleId)
		article.LikeCount = article.LikeCount - 1
		err := dao.UpdateArticleById(service.ArticleId, article)
		if err != nil {
			code = e.ERROR
			return serializer.Response{
				Code: code,
				Msg:  e.GetMsg(code),
			}
		}
		return serializer.Response{
			Code: code,
			Msg:  e.GetMsg(code),
			Data: "取消点赞成功",
		}
	}
}
