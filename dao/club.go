package dao

import "school-bbs/model"

// GetClubList 返回列表和总数
func GetClubList(page int, size int) ([]*model.Club, int64, error) {
	var count int64
	list := make([]*model.Club, 0)

	// 查询社团总数
	if err := DB.Model(new(model.Club)).Count(&count).Error; err != nil {
		return nil, 0, err
	}
	// 查询列表
	if err := DB.Model(new(model.Club)).
		Offset((page - 1) * size).Limit(size).
		Find(&list).Error; err != nil {
		return nil, 0, err
	}
	return list, count, nil
}

// 查询社长所创建的社团
func UserClubInfo(uid uint) ([]*model.Club, int64, error) {
	var count int64
	club := make([]*model.Club, 0)

	if err := DB.Model(new(model.Club)).Where("president_id=?", uid).Count(&count).Error; err != nil {
		return nil, 0, err
	}
	err := DB.Model(&model.Club{}).Where("president_id=?", uid).
		First(&club).Error
	return club, count, err
}

// 查询社团是否存在
func CheckClubExist(name string) (club *model.Club, exist bool, err error) {
	var count int64
	err = DB.Model(&model.Club{}).Where("name = ?", name).Count(&count).Error
	if count == 0 {
		return nil, false, err
	}
	if err != nil {
		return nil, false, err
	}
	return club, true, nil
}

// 创建社团
func CreateClub(club *model.Club) (err error) {
	return DB.Create(club).Error
}

// 删除社团
func ClubDelete(uid uint) (err error) {
	club := &model.Club{}            // 声明一个用户变量
	err = DB.First(&club, uid).Error // 通过 ID 查询用户
	if err != nil {
		return err
	}
	err = DB.Delete(&club).Error
	if err != nil {
		return err
	}
	return nil
}

// GetClubById 根据 id 获取社团
func GetClubById(uId uint) (club *model.Club, err error) {
	err = DB.Model(&model.Club{}).Where("id=?", uId).
		First(&club).Error
	return
}

// UpdateClubById 根据 id 更新社团信息
func UpdateClubById(uId uint, club *model.Club) error {
	return DB.Model(&model.Club{}).Where("id=?", uId).
		Updates(&club).Error
}

// 查询社团
func GetClubByNickname(name string) ([]*model.Club, error) {
	var club []*model.Club
	if err := DB.Model(&model.Club{}).Where("name LIKE ?", "%"+name+"%").Find(&club).Error; err != nil {
		return nil, err
	}
	return club, nil
}

// 我加入的社团
func GetMyClubInfo(myid uint) ([]*model.ClubMembers, int64, error) {
	var count int64
	list := make([]*model.ClubMembers, 0)
	// 查询社团总数
	if err := DB.Model(new(model.ClubMembers)).Where("user_id=?", myid).Count(&count).Error; err != nil {
		return nil, 0, err
	}
	// 查询列表
	if err := DB.Model(new(model.ClubMembers)).Where("user_id=?", myid).Find(&list).Error; err != nil {
		return nil, 0, err
	}
	return list, count, nil
}
