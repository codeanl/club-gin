package serializer

import "school-bbs/model"

type Comment struct {
	UserId       int64  `gorm:"index:idx_comment_user_id;not null" json:"userId" form:"userId"`       // 用户编号
	EntityId     int64  `gorm:"index:idx_comment_entity_id;not null" json:"entityId" form:"entityId"` // 被评论文章编号
	Content      string `gorm:"type:text;not null" json:"content" form:"content"`                     // 内容
	QuoteId      int64  `gorm:"not null"  json:"quoteId" form:"quoteId"`                              // 引用的评论编号
	LikeCount    int64  `gorm:"not null;default:0" json:"likeCount" form:"likeCount"`                 // 点赞数量
	CommentCount int64  `gorm:"not null;default:0" json:"commentCount" form:"commentCount"`           // 评论数量
}

// BuildClub 序列化社团
func BuildComment(comment *model.Comment) Comment {
	return Comment{
		UserId:       comment.UserId,
		EntityId:     comment.EntityId,
		Content:      comment.Content,
		QuoteId:      comment.QuoteId,
		LikeCount:    comment.LikeCount,
		CommentCount: comment.CommentCount,
	}
}
