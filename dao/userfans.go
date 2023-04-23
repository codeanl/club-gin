package dao

import (
	"school-bbs/model"
)

// 查询是否已关注
func CheckUserLikeExist(myid uint, followerId uint) (exist bool, err error) {
	var count int64
	if err := DB.Model(&model.UserFans{}).Where("noticer_id = ? and follower_id = ?", myid, followerId).Count(&count).Error; err != nil {
		// 查询出错，处理错误的逻辑
	}
	if count > 0 {
		return true, nil
	} else {
		return false, err
	}
}

// 关注数据插入
func CreateUserLike(userfans *model.UserFans) (err error) {
	return DB.Create(userfans).Error
}

// 退出关注
func UnUserLike(myid uint, followerId uint) (err error) {
	if err := DB.Where(&model.UserFans{
		NoticerID:  myid,
		FollowerID: followerId,
	}).Delete(&model.UserFans{}).Error; err != nil {
		return err
	}
	return nil
}
