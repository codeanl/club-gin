package dao

import (
	"school-bbs/model"
)

// 查询是否已点赞
func CheckCommentLikeExist(commentid uint, userid uint) (exist bool, err error) {
	var count int64
	if err := DB.Model(&model.CommentLike{}).Where("user_id = ? and comment_id = ?", userid, commentid).Count(&count).Error; err != nil {
		// 查询出错，处理错误的逻辑
	}
	if count > 0 {
		return true, nil
	} else {
		return false, err
	}
}

// 点赞数据插入
func CreateCommengLike(commentlike *model.CommentLike) (err error) {
	return DB.Create(commentlike).Error
}

// 退出社团
func UnLike(commentid uint, userid uint) (err error) {
	if err := DB.Where(&model.CommentLike{
		CommentID: commentid,
		UserID:    userid,
	}).Delete(&model.CommentLike{}).Error; err != nil {
		return err
	}
	return nil
}
