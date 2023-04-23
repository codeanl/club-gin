package dao

import (
	"school-bbs/model"
)

// 获取某个社团消息列表和总数
func GetMessageClub(cid uint) ([]*model.Notice, int64, error) {
	var count int64
	list := make([]*model.Notice, 0)
	// 查询总数
	if err := DB.Model(new(model.Notice)).Count(&count).Where("club_id=?", cid).Error; err != nil {
		return nil, 0, err
	}
	// 查询列表
	if err := DB.Model(new(model.Notice)).Where("club_id=?", cid).
		Find(&list).Error; err != nil {
		return nil, 0, err
	}
	return list, count, nil
}

// 创建消息
func CreateMessage(message *model.Notice) (err error) {
	return DB.Create(message).Error
}

// 删除消息
func MessageDelete(uid uint) (err error) {
	message := &model.Notice{}          // 声明一个用户变量
	err = DB.First(&message, uid).Error // 通过 ID 查询用户
	if err != nil {
		return err
	}
	err = DB.Delete(&message).Error
	if err != nil {
		return err
	}
	return nil
}

// GetMessageById 根据 id 获取消息
func GetMessageById(uId uint) (message *model.Notice, err error) {
	err = DB.Model(&model.Notice{}).Where("id=?", uId).
		First(&message).Error
	return
}

// UpdateMessageById 根据 id 更新信息
func UpdateMessageById(uId uint, message *model.Notice) error {
	return DB.Model(&model.Notice{}).Where("id=?", uId).
		Updates(&message).Error
}

// 查询我的信息
func GetMyMessage(id uint) ([]*model.Notice, error) {
	list := make([]*model.Notice, 0)
	// 查询列表
	if err := DB.Model(new(model.Notice)).
		Find(&list).Where("userId=? OR userId=?", id, 0).Error; err != nil {
		return nil, err
	}
	return list, nil
}
