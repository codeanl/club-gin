package model

import "gorm.io/gorm"

// 收藏社团
type FavoriteClub struct {
	gorm.Model
	UserId int64 `gorm:"index:idx_favorite_user_id;not null" json:"userId" form:"userId"`    // 用户编号
	ClubId int64 `gorm:"index:idx_favorite_article_id;not null" json:"clubId" form:"clubId"` // 收藏社团编号
}
