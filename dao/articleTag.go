package dao

import "school-bbs/model"

// 创建标签
func ArticleTagCreate(articleTag *model.ArticleTag) (err error) {
	return DB.Create(articleTag).Error
}
