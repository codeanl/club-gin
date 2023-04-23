package service

import (
	"github.com/gin-gonic/gin"
	"school-bbs/dao"
	"school-bbs/model"
	"school-bbs/serializer"
	"school-bbs/util"
	"school-bbs/util/e"
	"strconv"
)

type TagService struct {
	Name string `gorm:"not null" json:"name" form:"name" `
}

// 获取标签列表
func (service TagService) GetTagList(c *gin.Context) serializer.Response {
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

	list, count, err := dao.GetTagList(page, size)
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

// 创建标签
func (service TagService) TagAdd(c *gin.Context) serializer.Response {
	code := e.SUCCESS
	if service.Name == "" {
		code = e.InvalidParams
		return serializer.Response{
			Code: code,
			Msg:  e.GetMsg(code),
		}
	}
	// 判断标签是否已存在
	_, exist, err := dao.CheckTagExist(service.Name)
	if err != nil {
		code = e.ERROR
		return serializer.Response{
			Code: code,
			Msg:  e.GetMsg(code),
		}
	}
	if exist {
		code = e.ERROR
		return serializer.Response{
			Code: code,
			Msg:  e.GetMsg(code),
			Data: "该标签已存在",
		}
	}

	// 数据的插入
	tag := &model.Tag{
		Name: service.Name,
	}
	err = dao.CreateTag(tag)
	if err != nil {
		code = e.ERROR
		return serializer.Response{
			Code: code,
			Msg:  e.GetMsg(code),
		}
	}
	return serializer.Response{
		Code: code,
		Data: tag,
		Msg:  e.GetMsg(code),
	}
}
