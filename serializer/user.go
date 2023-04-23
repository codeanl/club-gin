package serializer

import "school-bbs/model"

type User struct {
	ID            uint
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
	IsAdmin       int    `gorm:"default:0"  json:"is_admin" form:"is_admin"`   //0-普通用户 1-管理员 2-社团管理员
	WechatUnionID string `gorm:"" json:"wechat_union_id"`                      //微信
	ThreadsCnt    int    `gorm:"default:0" json:"threads_cnt"`                 //发帖数
	PostsCnt      int    `gorm:"default:0" json:"posts_cnt"`                   //回帖数
	FavouriteCnt  int    `gorm:"default:0" json:"favourite_cnt"`               //收藏文章数
	AttentionCnt  int    `gorm:"default:0" json:"attention_cnt"`               //关注数
	Fans          int    `gorm:"default:0" json:"fans"`                        //粉丝数
}

// BuildUser 序列化用户
func BuildUser(user *model.User) User {
	return User{
		ID:            user.ID,
		Email:         user.Email,
		Username:      user.Username,
		NickName:      user.NickName,
		Gender:        user.Gender,
		Age:           user.Age,
		Phone:         user.Phone,
		Avatar:        user.Avatar,
		City:          user.City,
		Job:           user.Job,
		Desc:          user.Desc,
		IsAdmin:       user.IsAdmin,
		WechatUnionID: user.WechatUnionID,
		ThreadsCnt:    user.ThreadsCnt,
		PostsCnt:      user.PostsCnt,
		FavouriteCnt:  user.FavouriteCnt,
		AttentionCnt:  user.AttentionCnt,
		Fans:          user.Fans,
	}
}
