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

type ClubService struct {
	ID          uint
	Name        string `gorm:"" json:"name" form:"name"`       //名字
	Phone       string `gorm:"" json:"phone" form:"phone"`     //联系方式
	Desc        string `gorm:"" json:"desc" form:"desc"`       //简介
	Purpose     string `gorm:"" json:"purpose" form:"purpose"` //宗旨
	Avatar      string `gorm:""  json:"avatar" form:"avatar"`  //社团头像
	PresidentID uint   //会长id
	PeoPleCln   int    `gorm:"default:0"  json:"peoPleCln" form:"peoPleCln"`
}

// 获取社团列表
func (service ClubService) GetClubList(c *gin.Context) serializer.Response {
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

	list, count, err := dao.GetClubList(page, size)
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

// 创建社团
func (service ClubService) ClubAdd(c *gin.Context, id uint, logo multipart.File, fileSize int64, files []*multipart.FileHeader) serializer.Response {
	code := e.SUCCESS

	// 判断社团是否已存在
	_, exist, err := dao.CheckClubExist(service.Name)
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
			Data: "该社团已入驻",
		}
	}
	path, err := util.UploadToQiNiu(logo, fileSize)
	if err != nil {
		code = e.ERROR
		return serializer.Response{
			Code:  code,
			Data:  e.GetMsg(code),
			Error: path,
		}
	}
	// 数据的插入
	club := &model.Club{
		Name:        service.Name,
		PresidentID: id,
		Phone:       service.Phone,
		Desc:        service.Desc,
		Purpose:     service.Purpose,
		Avatar:      path,
	}
	err = dao.CreateClub(club)
	if err != nil {
		code = e.ERROR
		return serializer.Response{
			Code: code,
			Msg:  e.GetMsg(code),
		}
	}

	wg := new(sync.WaitGroup)
	wg.Add(len(files))
	for _, file := range files {
		tmp, _ := file.Open()
		path1, err := util.UploadToQiNiu(tmp, file.Size)
		if err != nil {
			code = e.ERROR
			return serializer.Response{
				Code:  code,
				Data:  e.GetMsg(code),
				Error: path,
			}
		}
		clubimg := &model.ClubImg{
			ClubID:  club.ID,
			ImgPath: path1,
		}
		err = dao.CreateClubImg(clubimg)
		if err != nil {
			code = e.ERROR
			return serializer.Response{
				Code:  code,
				Msg:   e.GetMsg(code),
				Error: err.Error(),
			}
		}
		wg.Done()
	}
	wg.Wait()
	clubmembers := &model.ClubMembers{
		ClubID: club.ID,
		UserID: id,
	}
	err = dao.CreateJoinClub(clubmembers)
	if err != nil {
		code = e.ERROR
		return serializer.Response{
			Code: code,
			Msg:  e.GetMsg(code),
		}
	}

	return serializer.Response{
		Code: code,
		Data: club,
		Msg:  e.GetMsg(code),
	}
}

// 查询社长所创建的社团
func (service ClubService) UserClubInfo(c *gin.Context, id uint) serializer.Response {
	code := e.SUCCESS

	club, count, err := dao.UserClubInfo(id)
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
			"list":  club,
			"count": count,
		},
	}
}

// 社团注销
func (service ClubService) ClubDelete(c *gin.Context, uid uint, cid uint) serializer.Response {
	code := e.SUCCESS
	clubinfo, _ := dao.GetClubById(cid)
	if clubinfo.PresidentID == uid {
		err := dao.ClubDelete(cid)
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
			Data: "你不是该社团的管理员",
		}
	}

}

// 社团修改信息
func (service ClubService) Update(c *gin.Context, id uint, uid uint) serializer.Response {
	var err error
	code := e.SUCCESS
	club, err := dao.GetClubById(id)

	if service.Phone != "" {
		club.Phone = service.Phone
	}
	if service.Desc != "" {
		club.Desc = service.Desc
	}
	if service.Purpose != "" {
		club.Purpose = service.Purpose
	}
	clubform, err := dao.GetClubById(id)
	if clubform.PresidentID == uid {
		err = dao.UpdateClubById(id, club)
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
			Data: serializer.BuildClub(club),
			Msg:  e.GetMsg(code),
		}
	} else {
		code = e.ERROR
		return serializer.Response{
			Code: code,
			Msg:  e.GetMsg(code),
			Data: "你不是该社团的会长，不能修改",
		}
	}

}

// 查询社团
func (service ClubService) GetClubInfo(c *gin.Context) serializer.Response {
	code := e.SUCCESS
	club, err := dao.GetClubByNickname(service.Name)
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
		Data: club,
	}
}

// 更新社团头像
func (service *ClubService) UploadClubAvatar(c *gin.Context, id uint, file multipart.File, fileSize int64, uid uint) serializer.Response {
	code := e.SUCCESS
	var err error
	path, err := util.UploadToQiNiu(file, fileSize)
	if err != nil {
		code = e.ERROR
		return serializer.Response{
			Code:  code,
			Data:  e.GetMsg(code),
			Error: path,
		}
	}
	club, err := dao.GetClubById(id)
	if err != nil {
		code = e.ERROR
		return serializer.Response{
			Code:  code,
			Msg:   e.GetMsg(code),
			Error: err.Error(),
		}
	}
	club.Avatar = path
	clubform, err := dao.GetClubById(id)
	if clubform.PresidentID == uid {
		err = dao.UpdateClubById(id, club)
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
			Data: serializer.BuildClub(club),
			Msg:  e.GetMsg(code),
		}
	} else {
		code = e.ERROR
		return serializer.Response{
			Code: code,
			Msg:  e.GetMsg(code),
			Data: "你不是该社团的会长，不能修改",
		}
	}
}

// 转让社长
func (service *ClubService) TransferClub(c *gin.Context, cid uint, uid uint, myid uint) serializer.Response {
	code := e.SUCCESS
	var err error
	club, err := dao.GetClubById(cid)
	if err != nil {
		code = e.ERROR
		return serializer.Response{
			Code:  code,
			Msg:   e.GetMsg(code),
			Error: err.Error(),
		}
	}
	if club.PresidentID == myid {
		//查询被转让的人是否存在
		userform, err := dao.GetUserById(uid)
		if userform.Username != "" {
			exist, _ := dao.CheckJoinClubExist(uid, cid)
			if exist {
				club.PresidentID = uid
				err = dao.UpdateClubById(cid, club)
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
					Data: serializer.BuildClub(club),
					Msg:  e.GetMsg(code),
				}
			} else {
				code = e.ERROR
				return serializer.Response{
					Code: code,
					Msg:  e.GetMsg(code),
					Data: "该用户不在此社团",
				}
			}
		} else {
			code = e.ERROR
			return serializer.Response{
				Code: code,
				Msg:  e.GetMsg(code),
				Data: "该用户不存在",
			}
		}
	} else {
		code = e.ERROR
		return serializer.Response{
			Code: code,
			Msg:  e.GetMsg(code),
			Data: "你不是该社团的会长，不能修改",
		}
	}
}

// 我加入的社团
func (service ClubService) GetMyClubInfo(c *gin.Context, myid uint) serializer.Response {
	code := e.SUCCESS
	club, count, err := dao.GetMyClubInfo(myid)
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
			"list":  club,
			"count": count,
		},
	}
}
