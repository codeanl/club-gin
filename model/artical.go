package model

import "gorm.io/gorm"

// 文章
type Article struct {
	gorm.Model
	UserId        uint   `gorm:"index:idx_article_user_id" json:"userId" form:"userId"` // 所属用户编号
	Title         string `gorm:"" json:"title" form:"title"`                            // 标题
	Summary       string `gorm:"" json:"summary" form:"summary"`                        // 摘要
	Content       string `gorm:"" json:"content" form:"content"`                        // 内容
	CommentCount  int64  `gorm:"default:0" json:"commentCount" form:"commentCount"`     // 评论数量
	LikeCount     int64  `gorm:"default:0" json:"likeCount" form:"likeCount"`           // 点赞数量
	InvolveClubId uint   `gorm:"default:0" json:"involveClubId" form:"involveClubId"`   //涉及社团id 0为不涉及
}
