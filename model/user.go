package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email         string `gorm:""  json:"email" form:"email"`                  //邮箱
	Username      string `gorm:""  json:"username" form:"username"`            //用户名
	Password      string `gorm:""  json:"password" form:"password"`            //密码
	NickName      string `gorm:"" form:"nickname" json:"nickname" form:"name"` //昵称
	Gender        int    `gorm:"default:0"  json:"gender" form:"gender"`       //性别:1-男/2-女/0-未知
	Age           int    `gorm:"default:0"  json:"gender" form:"gender"`       //年龄
	Phone         string `gorm:""  json:"phone" form:"phone"`                  //手机号码
	Avatar        string `gorm:""  json:"avatar" form:"avatar"`                //头像
	City          string `gorm:"" json:"city" form:"city"`                     //城市
	Job           string `gorm:"" json:"job" form:"job"`                       //职业
	Desc          string `gorm:""  json:"desc" form:"desc"`                    //简介
	IsAdmin       int    `gorm:"default:0"  json:"is_admin" form:"is_admin"`   //0-普通用户 1-管理员
	WechatUnionID string `gorm:"" json:"wechat_union_id"`                      //微信
	ThreadsCnt    int    `gorm:"default:0" json:"threads_cnt"`                 //发帖数
	PostsCnt      int    `gorm:"default:0" json:"posts_cnt"`                   //回帖数
	FavouriteCnt  int    `gorm:"default:0" json:"favourite_cnt"`               //收藏社团数
	AttentionCnt  int    `gorm:"default:0" json:"attention_cnt"`               //关注数
	Fans          int    `gorm:"default:0" json:"fans"`                        //粉丝数
}

const (
	PassWordCost = 12 //密码加密难度
)

// SetPassword 设置密码
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

// CheckPassword 校验密码
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}
