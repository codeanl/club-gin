package dao

import "school-bbs/model"

// CreateProductImg 创建活动图片
func CreateClubImg(clubImg *model.ClubImg) error {
	return DB.Model(&model.ClubImg{}).Create(&clubImg).Error
}
