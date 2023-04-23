package model

import "gorm.io/gorm"

// 用户点赞文章
type ArticleLike struct {
	gorm.Model
	UserId    uint `gorm:"not null;uniqueIndex:idx_user_like_unique;" json:"userId" form:"userId"`                                   // 用户
	ArticleId uint `gorm:"not null;uniqueIndex:idx_user_like_unique;index:idx_user_like_articleId;" json:"articleId" form:"article"` // 文章编号
}
