package dao

import "school-bbs/model"

// GetClubList 返回列表和总数 按使用次数排序
func GetTagList(page int, size int) ([]*model.Tag, int64, error) {
	var count int64
	list := make([]*model.Tag, 0)

	// 查询标签总数
	if err := DB.Model(new(model.Tag)).Count(&count).Error; err != nil {
		return nil, 0, err
	}
	// 查询列表
	if err := DB.Model(new(model.Tag)).Order("usecn DESC").
		Offset((page - 1) * size).Limit(size).
		Find(&list).Error; err != nil {
		return nil, 0, err
	}
	return list, count, nil
}

// 查询标签是否存在
func CheckTagExist(name string) (tag *model.Tag, exist bool, err error) {
	var count int64
	err = DB.Model(&model.Tag{}).Where("name = ?", name).Count(&count).Error
	if count == 0 {
		return nil, false, err
	}
	if err != nil {
		return nil, false, err
	}
	return tag, true, nil
}

// 创建标签
func CreateTag(tag *model.Tag) (err error) {
	return DB.Create(tag).Error
}

// 删除标签
func TagDelete(uid uint) (err error) {
	tag := &model.Tag{}             // 声明一个用户变量
	err = DB.First(&tag, uid).Error // 通过 ID 查询用户
	if err != nil {
		return err
	}
	err = DB.Delete(&tag).Error
	if err != nil {
		return err
	}
	return nil
}
func UpdateTagById(uId uint, club *model.Tag) error {
	return DB.Model(&model.Tag{}).Where("id=?", uId).
		Updates(&club).Error
}
func GetTagById(uId uint) (tag *model.Tag, err error) {
	err = DB.Model(&model.Tag{}).Where("id=?", uId).
		First(&tag).Error
	return
}
