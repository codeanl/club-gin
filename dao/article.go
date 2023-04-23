package dao

import "school-bbs/model"

// GetArticleList 返回列表和总数
func GetArticleList(page int, size int) ([]*model.Article, int64, error) {
	var count int64
	list := make([]*model.Article, 0)

	// 查询文章总数
	if err := DB.Model(new(model.Article)).Count(&count).Error; err != nil {
		return nil, 0, err
	}
	// 查询列表
	if err := DB.Model(new(model.Article)).
		Offset((page - 1) * size).Limit(size).
		Find(&list).Error; err != nil {
		return nil, 0, err
	}
	return list, count, nil
}

// 创建文章
func CreateArticle(article *model.Article) (err error) {
	return DB.Create(article).Error
}

// 删除文章
func ArticleDelete(uid uint) (err error) {
	article := &model.Article{}         // 声明一个用户变量
	err = DB.First(&article, uid).Error // 通过 ID 查询用户
	if err != nil {
		return err
	}
	err = DB.Delete(&article).Error
	if err != nil {
		return err
	}
	return nil
}

// GetClubById 根据 id 获取文章
func GetArticleById(uId uint) (article *model.Article, err error) {
	err = DB.Model(&model.Article{}).Where("id=?", uId).
		First(&article).Error
	return
}

// UpdateArticleById 根据 id 更新文章信息
func UpdateArticleById(uId uint, article *model.Article) error {
	return DB.Model(&model.Article{}).Where("id=?", uId).
		Updates(&article).Error
}

// 查询我的文章
func GetMyArticle(id uint) ([]*model.Article, error) {
	list := make([]*model.Article, 0)
	// 查询列表
	if err := DB.Model(new(model.Article)).
		Find(&list).Where("userId=?", id).Error; err != nil {
		return nil, err
	}
	return list, nil
}

// CreateArticleImg 创建文章图片
func CreateArticleImg(articleImg *model.ArticleImg) error {
	return DB.Model(&model.ArticleImg{}).Create(&articleImg).Error
}

// 查询文章
func GetArticleDetail(name string) ([]*model.Article, error) {
	var article []*model.Article
	if err := DB.Model(&model.Article{}).Where("title LIKE ? OR content LIKE ?  ", "%"+name+"%", "%"+name+"%").Find(&article).Error; err != nil {
		return nil, err
	}
	return article, nil
}
