package service

import (
	"github.com/gin-gonic/gin"
	"school-bbs/dao"
	"school-bbs/model"
	"school-bbs/serializer"
	"school-bbs/util/e"
)

type ArticleTagService struct {
	ArticleId int64 `gorm:"not null;index:idx_article_id;" json:"articleId" form:"articleId"` // 文章编号
	TagId     int64 `gorm:"not null;index:idx_article_tag_tag_id;" json:"tagId" form:"tagId"` // 标签编号
}

// 文章添加标签
func (service ArticleTagService) ArticleTagAdd(c *gin.Context) serializer.Response {
	code := e.SUCCESS
	// 数据的插入
	articleTag := &model.ArticleTag{
		ArticleId: service.ArticleId,
		TagId:     service.TagId,
	}

	err := dao.ArticleTagCreate(articleTag)
	if err != nil {
		code = e.ERROR
		return serializer.Response{
			Code: code,
			Msg:  e.GetMsg(code),
		}
	}
	tag, _ := dao.GetTagById(uint(service.TagId))
	tag.Usecn = tag.Usecn + 1
	err = dao.UpdateTagById(uint(service.TagId), tag)
	if err != nil {
		code = e.ERROR
		return serializer.Response{
			Code: code,
			Msg:  e.GetMsg(code),
		}
	}
	return serializer.Response{
		Code: code,
		Data: articleTag,
		Msg:  e.GetMsg(code),
	}
}
