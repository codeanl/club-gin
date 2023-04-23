package dao

import "school-bbs/model"

// 创建活动
func CreateActive(active *model.Active) (err error) {
	return DB.Create(active).Error
}

// CreateActiveImg 创建活动图片
func CreateActiveImg(activeImg *model.ActiveImg) error {
	return DB.Model(&model.ActiveImg{}).Create(&activeImg).Error
}

// GetArticleList 返回列表和总数
func GetActiveList(page int, size int) ([]*model.Active, int64, error) {
	var count int64
	list := make([]*model.Active, 0)

	// 查询活动总数
	if err := DB.Model(new(model.Active)).Count(&count).Error; err != nil {
		return nil, 0, err
	}
	// 查询列表
	if err := DB.Model(new(model.Active)).
		Offset((page - 1) * size).Limit(size).
		Find(&list).Error; err != nil {
		return nil, 0, err
	}
	return list, count, nil
}

// 获取某个社团的活动列表和总数
func GetActiveListByClub(page int, size int, cid uint) ([]*model.Active, int64, error) {
	var count int64
	list := make([]*model.Active, 0)

	// 查询活动总数
	if err := DB.Model(new(model.Active)).Where("club_id=?", cid).Count(&count).Error; err != nil {
		return nil, 0, err
	}
	// 查询列表
	if err := DB.Model(new(model.Active)).
		Offset((page - 1) * size).Limit(size).
		Find(&list).Error; err != nil {
		return nil, 0, err
	}
	return list, count, nil
}

// 删除活动
func ActiveDelete(uid uint) (err error) {
	active := &model.Active{}          // 声明一个用户变量
	err = DB.First(&active, uid).Error // 通过 ID 查询用户
	if err != nil {
		return err
	}
	err = DB.Delete(&active).Error
	if err != nil {
		return err
	}
	return nil
}

// GetClubById 根据 id 获取活动
func GetActiveById(uId uint) (active *model.Active, err error) {
	err = DB.Model(&model.Active{}).Where("id=?", uId).
		First(&active).Error
	return
}

// UpdateActiveById 根据 id 更新活动信息
func UpdateActiveById(uId uint, active *model.Active) error {
	return DB.Model(&model.Active{}).Where("id=?", uId).
		Updates(&active).Error
}

// 查询活动
func GetActiveDetail(name string) ([]*model.Active, error) {
	var active []*model.Active
	if err := DB.Model(&model.Active{}).Where("title LIKE ? OR content LIKE ?  ", "%"+name+"%", "%"+name+"%").Find(&active).Error; err != nil {
		return nil, err
	}
	return active, nil
}
