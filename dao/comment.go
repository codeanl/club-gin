package dao

import "school-bbs/model"

// GetCommentList 返回列表和总数
func GetCommentList(id int64) ([]*model.Comment, int64, error) {
	var count int64
	list := make([]*model.Comment, 0)
	// 查询评论总数
	if err := DB.Model(new(model.Comment)).Count(&count).Where("entity_id = ?", id).Error; err != nil {
		return nil, 0, err
	}
	// 查询列表
	if err := DB.Model(new(model.Comment)).Where("entity_id = ?", id).Order("like_count DESC").
		Find(&list).Error; err != nil {
		return nil, 0, err
	}
	return list, count, nil
}

// 删除评论
func CommentDelete(uid uint) (err error) {
	comment := &model.Comment{}         // 声明一个用户变量
	err = DB.First(&comment, uid).Error // 通过 ID 查询用户
	if err != nil {
		return err
	}
	err = DB.Delete(&comment).Error
	if err != nil {
		return err
	}
	return nil
}

// 创建评论
func CreateComment(comment *model.Comment) (err error) {
	return DB.Create(comment).Error
}

// GetCommentById 根据 id 获取评论
func GetCommentById(uId uint) (comment *model.Comment, err error) {
	err = DB.Model(&model.Comment{}).Where("id=?", uId).
		First(&comment).Error
	return comment, err
}
func UpdateCommentById(uId uint, comment *model.Comment) error {
	return DB.Model(&model.Comment{}).Where("id=?", uId).
		Updates(&comment).Error
}
