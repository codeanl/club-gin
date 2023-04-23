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
	"time"
)

// UserService 管理用户服务
type UserService struct {
	Email         string `gorm:""  json:"email" form:"email"`                      //邮箱
	Username      string `gorm:""  json:"username" form:"username"`                //用户名
	Password      string `gorm:""  json:"password" form:"password"`                //密码
	NickName      string `gorm:"column:nickname"  json:"nickname" form:"nickname"` //昵称
	Gender        int    `gorm:"default:0"  json:"gender" form:"gender"`           //性别:1-男/2-女/0-未知
	Age           int    `gorm:"default:0"  json:"age" form:"age"`                 //年龄
	Phone         string `gorm:""  json:"phone" form:"phone"`                      //手机号码
	Avatar        string `gorm:""  json:"avatar" form:"avatar"`                    //头像
	City          string `gorm:"" json:"city" form:"city"`                         //城市
	Job           string `gorm:"" json:"job" form:"job"`                           //职业
	Desc          string `gorm:""  json:"desc" form:"desc"`                        //简介
	IsAdmin       int    `gorm:"default:0"  json:"is_admin" form:"is_admin"`       //0-普通用户 1-管理员 2-社团管理员
	WechatUnionID string `gorm:"" json:"wechat_union_id"`                          //微信
	ThreadsCnt    int    `gorm:"default:0" json:"threads_cnt"`                     //发帖数
	PostsCnt      int    `gorm:"default:0" json:"posts_cnt"`                       //回帖数
	FavouriteCnt  int    `gorm:"default:0" json:"favourite_cnt"`                   //收藏社团数
	AttentionCnt  int    `gorm:"default:0" json:"attention_cnt"`                   //关注数
	Fans          int    `gorm:"default:0" json:"fans"`                            //粉丝数
	UserCode      string `gorm:"" json:"usercode" form:"usercode"`                 //验证码
}

// 用户注册
func (service UserService) Register(c *gin.Context) serializer.Response {
	code := e.SUCCESS
	//输入昵称 用户名 邮箱 密码进行注册
	if service.NickName == "" || service.UserCode == "" || service.Username == "" || service.Email == "" || service.Password == "" {
		code = e.InvalidParams
		return serializer.Response{
			Code: code,
			Msg:  e.GetMsg(code),
		}
	}
	// 验证验证码是否正确
	sysCode, err := dao.RDB.Get(c, service.Email).Result()
	if err != nil {
		code = e.ERROR
		return serializer.Response{
			Code: code,
			Msg:  e.GetMsg(code),
			Data: "验证码不正确，请重新获取验证码",
		}
	}
	if sysCode != service.UserCode {
		code = e.ERROR
		return serializer.Response{
			Code: code,
			Msg:  e.GetMsg(code),
			Data: "验证码不正确",
		}
	}
	// 判断邮箱是否已存在
	_, exist, err := dao.GetUserEmail(service.Email)
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
			Data: "该邮箱已被注册",
		}
	}

	// 数据的插入
	user := &model.User{
		Username: service.Username,
		Password: service.Password,
		NickName: service.NickName,
		Email:    service.Email,
	}
	//加密密码
	if err = user.SetPassword(service.Password); err != nil {
		code = e.ERROR
		return serializer.Response{
			Code: code,
			Msg:  e.GetMsg(code),
		}
	}
	user.Avatar = "https://an23.co/upload/2022/07/%E5%BE%AE%E4%BF%A1%E5%9B%BE%E7%89%87_20220722020131.jpg"
	err = dao.CreateUser(user)
	if err != nil {
		code = e.ERROR
		return serializer.Response{
			Code: code,
			Msg:  e.GetMsg(code),
		}
	}
	return serializer.Response{
		Code: code,
		Data: user,
		Msg:  e.GetMsg(code),
	}
}

// 发送验证码
func (service UserService) SendCode(c *gin.Context) serializer.Response {
	code := e.SUCCESS
	if service.Email == "" {
		code = e.InvalidParams
		return serializer.Response{
			Code: code,
			Msg:  e.GetMsg(code),
		}
	}
	authcode := util.GetRand()
	dao.RDB.Set(c, service.Email, authcode, time.Second*300)
	err := util.SendCode(service.Email, authcode)
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
		Data: "验证码已发送",
	}
}

// 用户登陆函数
func (service UserService) Login(c *gin.Context) serializer.Response {
	code := e.SUCCESS
	user, exist, err := dao.ExistOrNotByUserName(service.Username)
	if !exist { //如果查询不到，返回相应的错误
		code = e.ErrorNotExistUser
		return serializer.Response{
			Code: code,
			Msg:  e.GetMsg(code),
		}
	}
	if user.CheckPassword(service.Password) == false {
		code = e.ErrorNotCompare
		return serializer.Response{
			Code: code,
			Msg:  e.GetMsg(code),
		}
	}
	token, err := util.GenerateToken(user.ID, service.Username, 0, user.IsAdmin)
	if err != nil {
		code = e.ERROR
		return serializer.Response{
			Code: code,
			Msg:  e.GetMsg(code),
		}
	}
	return serializer.Response{
		Code: code,
		Data: serializer.TokenData{User: serializer.BuildUser(user), Token: token},
		Msg:  e.GetMsg(code),
	}
}

// 更新用户头像
func (service *UserService) UploadAvatar(c *gin.Context, uId uint, file multipart.File, fileSize int64) serializer.Response {
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
	user, err := dao.GetUserById(uId)
	if err != nil {
		code = e.ERROR
		return serializer.Response{
			Code:  code,
			Msg:   e.GetMsg(code),
			Error: err.Error(),
		}
	}
	user.Avatar = path
	err = dao.UpdateUserById(uId, user)
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
		Data: serializer.BuildUser(user),
		Msg:  e.GetMsg(code),
	}
}

// 用户修改密码
func (service *UserService) UpdatePassword(c *gin.Context, uId uint, oldpassword string, newpassword string) serializer.Response {
	code := e.SUCCESS
	if oldpassword == "" || newpassword == "" {
		code = e.InvalidParams
		return serializer.Response{
			Code: code,
			Msg:  e.GetMsg(code),
		}
	}
	user, _ := dao.GetUserById(uId)
	if user.CheckPassword(oldpassword) == false {
		code = e.ERROR
		return serializer.Response{
			Code: code,
			Msg:  e.GetMsg(code),
			Data: "旧密码错误",
		}
	} else {
		user.SetPassword(newpassword)
		err := dao.UpdateUserById(uId, user)
		if err != nil {
			code = e.ERROR
			return serializer.Response{
				Code: code,
				Msg:  e.GetMsg(code),
			}
		}
		return serializer.Response{
			Code: code,
			Data: serializer.BuildUser(user),
			Msg:  e.GetMsg(code),
		}
	}

}

// 用户修改信息
func (service UserService) Update(c *gin.Context, uId uint) serializer.Response {
	var err error
	code := e.SUCCESS
	user, err := dao.GetUserById(uId)
	if service.NickName != "" {
		user.NickName = service.NickName
	}
	if service.Gender != 0 {
		user.Gender = service.Gender
	}
	if service.Age != 0 {
		user.Age = service.Age
	}
	if service.Phone != "" {
		user.Phone = service.Phone
	}
	if service.City != "" {
		user.City = service.City
	}
	if service.Job != "" {
		user.Job = service.Job
	}
	if service.Desc != "" {
		user.Desc = service.Desc
	}
	if service.WechatUnionID != "" {
		user.WechatUnionID = service.WechatUnionID
	}
	err = dao.UpdateUserById(uId, user)
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
		Data: serializer.BuildUser(user),
		Msg:  e.GetMsg(code),
	}
}

// 获取用户列表
func (service UserService) GetUserList(c *gin.Context) serializer.Response {
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
	list, count, err := dao.GetUserList(page, size)
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

// 用户注销
func (service UserService) UserMyDelete(c *gin.Context, uid uint) serializer.Response {
	code := e.SUCCESS
	err := dao.DeleteUser(uid)
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
}

// 获取自己的详细信息
func (service UserService) GetMyInfo(c *gin.Context, uid uint) serializer.Response {
	code := e.SUCCESS
	data, err := dao.GetUserById(uid)
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

// 查询用户
func (service UserService) GetUserDetail(c *gin.Context) serializer.Response {
	code := e.SUCCESS
	user, err := dao.GetUserByNickname(service.NickName)
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
		Data: user,
	}
}

// 找回账号密码
func (service UserService) Retrieve(c *gin.Context) serializer.Response {
	code := e.SUCCESS
	if service.UserCode == "" || service.Email == "" {
		code = e.InvalidParams
		return serializer.Response{
			Code: code,
			Msg:  e.GetMsg(code),
		}
	}
	// 验证验证码是否正确
	sysCode, err := dao.RDB.Get(c, service.Email).Result()
	if err != nil {
		code = e.ERROR
		return serializer.Response{
			Code: code,
			Msg:  e.GetMsg(code),
			Data: "验证码不正确，请重新获取验证码",
		}
	}
	if sysCode != service.UserCode {
		code = e.ERROR
		return serializer.Response{
			Code: code,
			Msg:  e.GetMsg(code),
			Data: "验证码不正确",
		}
	}
	user, err := dao.RetrieveUser(service.Email)
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
		Data: user,
	}
}
