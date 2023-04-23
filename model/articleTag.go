package model

import "gorm.io/gorm"

// 文章标签
type ArticleTag struct {
	gorm.Model
	ArticleId int64 `gorm:"not null;index:idx_article_id;" json:"articleId" form:"articleId"` // 文章编号
	TagId     int64 `gorm:"not null;index:idx_article_tag_tag_id;" json:"tagId" form:"tagId"` // 标签编号
}
