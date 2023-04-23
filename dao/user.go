package dao

import (
	"school-bbs/model"
)

// 查询邮箱是否存在
func GetUserEmail(email string) (user *model.User, exist bool, err error) {
	var count int64
	err = DB.Model(&model.User{}).Where("email = ?", email).Count(&count).Error
	if count == 0 {
		return nil, false, err
	}
	if err != nil {
		return nil, false, err
	}
	return user, true, nil
}

// 创建用户
func CreateUser(user *model.User) (err error) {
	return DB.Create(user).Error
}

// ExistOrNotByUserName 根据username判断是否存在该名字
func ExistOrNotByUserName(userName string) (user *model.User, exist bool, err error) {
	var count int64
	err = DB.Model(&model.User{}).Where("username=?", userName).
		Count(&count).Error
	if count == 0 {
		return nil, false, err
	}
	err = DB.Model(&model.User{}).Where("username=?", userName).
		First(&user).Error
	if err != nil {
		return nil, false, err
	}
	return user, true, nil
}

// GetUserById 根据 id 获取用户
func GetUserById(uId uint) (user *model.User, err error) {
	err = DB.Model(&model.User{}).Where("id=?", uId).
		First(&user).Error
	return
}

// 找回账号密码
func RetrieveUser(email string) (user *model.User, err error) {
	err = DB.Model(&model.User{}).Where("email=?", email).
		First(&user).Error
	return user, err
}

// 查询用户
func GetUserByNickname(nickname string) ([]*model.User, error) {
	var users []*model.User
	if err := DB.Model(&model.User{}).Where("nick_name LIKE ?", "%"+nickname+"%").Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// UpdateUserById 根据 id 更新用户信息
func UpdateUserById(uId uint, user *model.User) error {
	return DB.Model(&model.User{}).Where("id=?", uId).
		Updates(&user).Error
}

// GetUserList 返回列表和总数
func GetUserList(page int, size int) ([]*model.User, int64, error) {
	var count int64
	list := make([]*model.User, 0)

	// 查询用户总数
	if err := DB.Model(new(model.User)).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	// 查询列表
	if err := DB.Model(new(model.User)).
		Offset((page - 1) * size).Limit(size).
		Find(&list).Error; err != nil {
		return nil, 0, err
	}

	return list, count, nil
}

// 删除用户
func DeleteUser(uid uint) (err error) {
	user := &model.User{}            // 声明一个用户变量
	err = DB.First(&user, uid).Error // 通过 ID 查询用户
	if err != nil {
		return err
	}

	err = DB.Delete(&user).Error
	if err != nil {
		return err
	}
	return nil
}
