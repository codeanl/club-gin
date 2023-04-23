package service

import (
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"school-bbs/dao"
	"school-bbs/model"
	"school-bbs/serializer"
	"school-bbs/util"
	"school-bbs/util/e"
	"strconv"
	"sync"
)

// 文章
type ArticleService struct {
	UserId        uint   `gorm:"" json:"userId" form:"userId"`                        // 所属用户编号
	Title         string `gorm:"" json:"title" form:"title"`                          // 标题
	Summary       string `gorm:"" json:"summary" form:"summary"`                      // 摘要
	Content       string `gorm:"" json:"content" form:"content"`                      // 内容
	InvolveClubId uint   `gorm:"default:0" json:"involveClubId" form:"involveClubId"` //涉及社团id 0为不涉及
	CommentCount  int64  `gorm:"default:0" json:"commentCount" form:"commentCount"`   // 评论数量
	LikeCount     int64  `gorm:"default:0" json:"likeCount" form:"likeCount"`         // 点赞数量
}

// 获取文章列表
func (service ArticleService) GetArticleList(c *gin.Context) serializer.Response {
	code := e.SUCCESS
	size, _ := strconv.Atoi(c.DefaultQuery("size", util.DefaultSize))
	page, err := strconv.Atoi(c.DefaultQuery("page", util.DefaultPage))
	if err != nil {
		code = e.ERROR
		return serializer.Response{
			Code: code,
			Msg:  e.GetMsg(code),
		}
	}

	list, count, err := dao.GetArticleList(page, size)
	if err != nil {
		code = e.ERROR
		return serializer.Response{
			Code: code,
			Msg:  e.GetMsg(code),
		}
	}

	return serializer.Response{
		Code: code,
		Msg:  e.GetMsg(code),
		Data: map[string]interface{}{
			"list":  list,
			"count": count,
		},
	}
}

// 创建文章
func (service ArticleService) ArticleAdd(c *gin.Context, id uint, files []*multipart.FileHeader) serializer.Response {
	code := e.SUCCESS
	if service.Title == "" || service.Summary == "" || service.Content == "" {
		code = e.InvalidParams
		return serializer.Response{
			Code: code,
			Msg:  e.GetMsg(code),
		}
	}
	articcleform, _ := dao.GetArticleById(service.InvolveClubId)
	var involveClubId uint
	if articcleform.Title != "" {
		involveClubId = service.InvolveClubId
	} else {
		involveClubId = 0
	}

	// 数据的插入
	article := &model.Article{
		UserId:        id,
		Title:         service.Title,
		Summary:       service.Summary,
		Content:       service.Content,
		InvolveClubId: involveClubId,
	}
	err := dao.CreateArticle(article)
	if err != nil {
		code = e.ERROR
		return serializer.Response{
			Code: code,
			Msg:  e.GetMsg(code),
			Data: err,
		}
	}

	wg := new(sync.WaitGroup)
	wg.Add(len(files))
	for _, file := range files {
		tmp, _ := file.Open()
		path, err := util.UploadToQiNiu(tmp, file.Size)
		if err != nil {
			code = e.ERROR
			return serializer.Response{
				Code:  code,
				Data:  e.GetMsg(code),
				Error: path,
			}
		}
		articleImg := &model.ArticleImg{
			ArticleID: article.ID,
			ImgPath:   path,
		}
		err = dao.CreateArticleImg(articleImg)
		if err != nil {
			code = e.ERROR
			return serializer.Response{
				Code: code,
				Msg:  e.GetMsg(code),
			}
		}
		wg.Done()
	}
	wg.Wait()
	return serializer.Response{
		Code: code,
		Data: article,
		Msg:  e.GetMsg(code),
	}
}

// 删除文章
func (service ArticleService) ArticleDelete(c *gin.Context, aid uint, myid uint) serializer.Response {
	code := e.SUCCESS
	articleform, err := dao.GetArticleById(aid)
	if articleform.UserId == myid {
		err = dao.ArticleDelete(aid)
		if err != nil {
			code = e.ERROR
			return serializer.Response{
				Code:  code,
				Msg:   e.GetMsg(code),
				Error: err.Error(),
			}
		}
		return serializer.Response{
			Code: code,
			Msg:  e.GetMsg(code),
		}
	} else {
		code = e.ERROR
		return serializer.Response{
			Code: code,
			Msg:  e.GetMsg(code),
			Data: "该文章不是你的",
		}
	}

}

// 修改文章信息
func (service ArticleService) ArticleUpdate(c *gin.Context, aid uint, myid uint) serializer.Response {
	var err error
	code := e.SUCCESS
	article, err := dao.GetArticleById(aid)
	if service.Content != "" {
		article.Content = service.Content
	}
	if service.Title != "" {
		article.Title = service.Title
	}
	if service.Summary != "" {
		article.Summary = service.Summary
	}
	articleform, err := dao.GetArticleById(aid)
	if articleform.UserId == myid {
		err = dao.UpdateArticleById(aid, article)
		if err != nil {
			code = e.ERROR
			return serializer.Response{
				Code:  code,
				Msg:   e.GetMsg(code),
				Error: err.Error(),
			}
		}

		return serializer.Response{
			Code: code,
			Data: serializer.BuildArticle(article),
			Msg:  e.GetMsg(code),
		}
	} else {
		code = e.ERROR
		return serializer.Response{
			Code: code,
			Msg:  e.GetMsg(code),
			Data: "该文章不是你的",
		}
	}
}

// 搜索文章
func (service ArticleService) GetArticleDetail(c *gin.Context, name string) serializer.Response {
	code := e.SUCCESS
	data, err := dao.GetArticleDetail(name)
	if err != nil {
		code = e.ERROR
		return serializer.Response{
			Code: code,
			Msg:  e.GetMsg(code),
		}
	}
	return serializer.Response{
		Code: code,
		Msg:  e.GetMsg(code),
		Data: data,
	}
}

// 查询我的文章
func (service ArticleService) GetMyArticle(c *gin.Context, id uint) serializer.Response {
	code := e.SUCCESS
	data, err := dao.GetMyArticle(id)
	if err != nil {
		code = e.ERROR
		return serializer.Response{
			Code: code,
			Msg:  e.GetMsg(code),
		}
	}
	return serializer.Response{
		Code: code,
		Msg:  e.GetMsg(code),
		Data: data,
	}
}
