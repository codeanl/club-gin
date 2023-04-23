package dao

import "school-bbs/model"

// 查询是否已加入收藏
func CheckJoinFavoriteExist(userid int64, clubid int64) (exist bool, err error) {
	var count int64
	if err := DB.Model(&model.FavoriteClub{}).Where("user_id = ? and club_id = ?", userid, clubid).Count(&count).Error; err != nil {
		// 查询出错，处理错误的逻辑
	}
	if count > 0 {
		return true, nil
	} else {
		return false, err
	}
}

// 收藏数据插入
func CreateFavorite(favorite *model.FavoriteClub) (err error) {
	return DB.Create(favorite).Error
}

// 取消收藏
func QuidFavorite(userid int64, clubid int64) (err error) {
	if err := DB.Where(&model.FavoriteClub{UserId: userid, ClubId: clubid}).Delete(&model.FavoriteClub{}).Error; err != nil {
		return err
	}
	return nil
}

// 我收藏的社团
func GetMyFavoriteClub(myid uint) ([]*model.FavoriteClub, int64, error) {
	var count int64
	list := make([]*model.FavoriteClub, 0)
	// 查询社团总数
	if err := DB.Model(new(model.FavoriteClub)).Where("user_id=?", myid).Count(&count).Error; err != nil {
		return nil, 0, err
	}
	// 查询列表
	if err := DB.Model(new(model.FavoriteClub)).Where("user_id=?", myid).Find(&list).Error; err != nil {
		return nil, 0, err
	}
	return list, count, nil
}
