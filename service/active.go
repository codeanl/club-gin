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

type ActiveService struct {
	ClubID    uint   `gorm:"not null" json:"clubId" form:"clubId"`             //社团id
	Title     string `gorm:"size:1024" json:"title" form:"title"`              // 消息标题
	Content   string `gorm:"type:text;not null" json:"content" form:"content"` // 消息内容
	Cover     string `json:"cover" form:"cover"`                               //封面图
	StartTime string `json:"starttime" form:"starttime"`                       //活动时间
	EndTime   string `json:"endtime" form:"endtime"`                           //结束时间
	Place     string `json:"place" form:"place"`                               //活动地点
	Creator   string `json:"creator" form:"creator"`                           //活动创建人
	MaxPeople string `json:"maxpeaple" form:"maxpeaple"`                       //限制人数
}

// 创建活动
func (service ActiveService) ActiveAdd(c *gin.Context, id uint, files []*multipart.FileHeader) serializer.Response {
	code := e.SUCCESS
	if service.Title == "" || service.Content == "" {
		code = e.InvalidParams
		return serializer.Response{
			Code: code,
			Msg:  e.GetMsg(code),
		}
	}
	club, err := dao.GetClubById(service.ClubID)
	if err != nil {
		code = e.ERROR
		return serializer.Response{
			Code: code,
			Data: e.GetMsg(code),
		}
	}
	if club.PresidentID == id {
		tmp, _ := files[0].Open()
		path1, err := util.UploadToQiNiu(tmp, files[0].Size)
		if err != nil {
			code = e.ERROR
			return serializer.Response{
				Code: code,
				Data: e.GetMsg(code),
			}
		}
		active := &model.Active{
			ClubID:    service.ClubID,
			Title:     service.Title,
			Content:   service.Content,
			Cover:     path1,
			StartTime: service.StartTime,
			EndTime:   service.EndTime,
			Place:     service.Place,
			MaxPeople: service.MaxPeople,
		}
		err = dao.CreateActive(active)
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
			tmp, _ = file.Open()
			path, err := util.UploadToQiNiu(tmp, file.Size)
			if err != nil {
				code = e.ERROR
				return serializer.Response{
					Code:  code,
					Data:  e.GetMsg(code),
					Error: path,
				}
			}
			activeImg := &model.ActiveImg{
				ActiveID: active.ID,
				ImgPath:  path,
			}
			err = dao.CreateActiveImg(activeImg)
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
			Data: serializer.BuildActive(active),
			Msg:  e.GetMsg(code),
		}
	} else {
		code = e.ERROR
		return serializer.Response{
			Code: code,
			Data: "你不是该社团的社长，不能创建活动",
		}
	}
}

// 获取活动列表
func (service ActiveService) GetActiveList(c *gin.Context) serializer.Response {
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
	list, count, err := dao.GetActiveList(page, size)
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

// 获取某个社团的活动列表
func (service ActiveService) GetActiveListByClub(c *gin.Context) serializer.Response {
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
	list, count, err := dao.GetActiveListByClub(page, size, service.ClubID)
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

// 删除活动
func (service ActiveService) ActiveDelete(c *gin.Context, aid uint, uid uint) serializer.Response {
	code := e.SUCCESS
	activeform, _ := dao.GetActiveById(aid)
	if activeform.Title == "" {
		code = e.ERROR
		return serializer.Response{
			Code: code,
			Data: "该活动不存在",
		}
	} else {
		club, err := dao.GetClubById(activeform.ClubID)
		if err != nil {
			code = e.ERROR
			return serializer.Response{
				Code: code,
				Data: e.GetMsg(code),
			}
		}
		if club.PresidentID == uid {
			err := dao.ActiveDelete(aid)
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
				Data: "你不是该社团的社长，不能删除该活动",
			}
		}
	}
}

// 修改活动信息
func (service ActiveService) ActiveUpdate(c *gin.Context, aid uint, uid uint) serializer.Response {
	var err error
	code := e.SUCCESS
	active, err := dao.GetActiveById(aid)
	if service.ClubID != 0 {
		active.ClubID = service.ClubID
	}
	if service.Title != "" {
		active.Title = service.Title
	}
	if service.Content != "" {
		active.Content = service.Content
	}
	if service.StartTime != "" {
		active.StartTime = service.StartTime
	}
	if service.EndTime != "" {
		active.EndTime = service.EndTime
	}
	if service.Place != "" {
		active.Place = service.Place
	}
	if service.MaxPeople != "" {
		active.MaxPeople = service.MaxPeople
	}
	activeform, _ := dao.GetActiveById(aid)
	club, err := dao.GetClubById(activeform.ClubID)
	if club.PresidentID == uid {
		err = dao.UpdateActiveById(aid, active)
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
			Data: serializer.BuildActive(active),
			Msg:  e.GetMsg(code),
		}
	} else {
		code = e.ERROR
		return serializer.Response{
			Code: code,
			Data: "你不是该社团的社长，不更新该活动",
		}
	}
}

// 搜索活动
func (service ActiveService) GetActiveDetail(c *gin.Context, name string) serializer.Response {
	code := e.SUCCESS
	data, err := dao.GetActiveDetail(name)
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
