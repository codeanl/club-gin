package dao

import (
	"school-bbs/model"
)

// 查询是否已点赞
func CheckArticleLikeExist(articlelike uint, userid uint) (exist bool, err error) {
	var count int64
	if err := DB.Model(&model.ArticleLike{}).Where("user_id = ? and article_id = ?", userid, articlelike).Count(&count).Error; err != nil {
		// 查询出错，处理错误的逻辑
	}
	if count > 0 {
		return true, nil
	} else {
		return false, err
	}
}

// 点赞数据插入
func CreateArticleLike(articlelike *model.ArticleLike) (err error) {
	return DB.Create(articlelike).Error
}

// 取消点赞
func ArticleUnLike(articleid uint, userid uint) (err error) {
	if err := DB.Where(&model.ArticleLike{
		UserId:    userid,
		ArticleId: articleid,
	}).Unscoped().Delete(&model.ArticleLike{}).Error; err != nil {
		return err
	}
	return nil
}
