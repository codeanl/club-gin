package dao

import "school-bbs/model"

// 查询是否已加入社团
func CheckJoinClubExist(userid uint, clubid uint) (exist bool, err error) {

	var count int64
	if err := DB.Model(&model.ClubMembers{}).Where("user_id = ? and club_id = ?", userid, clubid).Count(&count).Error; err != nil {
		// 查询出错，处理错误的逻辑
	}
	if count > 0 {
		return true, nil
	} else {
		return false, err
	}
}

// 学生社团数据插入
func CreateJoinClub(clubmembers *model.ClubMembers) (err error) {
	return DB.Create(clubmembers).Error
}

// 退出社团
func QuidClub(userid uint, clubid uint) (err error) {
	if err := DB.Where(&model.ClubMembers{UserID: userid, ClubID: clubid}).Delete(&model.ClubMembers{}).Error; err != nil {
		return err
	}
	return nil
}
