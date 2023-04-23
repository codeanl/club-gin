package router

import (
	"github.com/gin-gonic/gin"
	"school-bbs/api"
	"school-bbs/conf"
	"school-bbs/middleware"
)

// 路由配置
func InitRouter() *gin.Engine {
	gin.SetMode(conf.AppMode)
	r := gin.Default()
	r.Use(middleware.Cors())
	r.GET("hello", func(context *gin.Context) {
		context.JSON(200, gin.H{"msg": "hello"})
	})

	v1 := r.Group("/")
	{
		v1.POST("/send-code", api.SendCode)       // 发送验证码
		v1.POST("/register", api.Register)        // 用户注册
		v1.POST("/login", api.Login)              // 用户登录
		v1.GET("/user-list", api.GetUserList)     // 获取用户列表
		v1.GET("/user-search", api.GetUserDetail) // 查询用户
		v1.POST("/user-retrieve", api.Retrieve)   // 找回账号密码

		v1.GET("/club-list", api.GetClubList)   // 获取社团列表和总数
		v1.GET("/club-search", api.GetClubInfo) // 查询社团

		v1.GET("/active-list", api.GetActiveList)            // 获取活动列表
		v1.GET("/activecclub-list", api.GetActiveListByClub) // 获取某个社团的活动列表

		v1.GET("/active-search", api.GetActiveDetail) // 搜索活动

		v1.GET("/article-list", api.GetArticleList)     // 获取文章列表
		v1.GET("/article-search", api.GetArticleDetail) // 搜索文章

		v1.GET("/tag-list", api.GetTagList) //获取标签列表

		v1.GET("/comment-list", api.GetCommentList) //获取某个文章的评论和次数

	}

	auth := r.Group("/") //需要登陆保护
	auth.Use(middleware.JWT())
	{
		auth.POST("/user-uploadavatar", api.UploadAvatar)      //更新头像
		auth.PUT("/update-user", api.UserUpdate)               //用户信息修改
		auth.PUT("/update-UpdatePassword", api.UpdatePassword) //用户修改密码
		auth.DELETE("/delete-myuser", api.UserMyDelete)        // 用户注销
		auth.GET("/user-myinfo", api.GetMyInfo)                // 获取自己的详细信息

		auth.POST("/userlike", api.UserLike)     //关注
		auth.POST("/userunlike", api.UserUnLike) //取消关注

		auth.POST("/add-club", api.ClubAdd)                   //创建社团
		auth.GET("/myclub", api.UserClubInfo)                 // 我创建的社团
		auth.PUT("/update-club", api.ClubUpdate)              // 更新社团信息
		auth.DELETE("/delete-myclub", api.ClubDelete)         // 社团注销
		auth.POST("/club-uploadavatar", api.UploadClubAvatar) //更新头像
		auth.POST("/club-transfer", api.TransferClub)         //转让社长
		auth.GET("/myclub-list", api.GetMyClubInfo)           // 获取我加入的社团列表和总数

		auth.POST("/joidclub", api.JoinClub) // 学生加入社团
		auth.POST("/quitclub", api.QuitClub) // 学生退出社团

		auth.POST("/active-add", api.ActiveAdd)         // 创建活动
		auth.DELETE("/delete-active", api.ActiveDelete) // 删除活动
		auth.PUT("/update-active", api.ActiveUpdate)    // 更新活动信息

		auth.POST("/article-add", api.ArticleAdd)         //创建文章
		auth.DELETE("/delete-article", api.ArticleDelete) // 删除文章
		auth.PUT("/update-article", api.ArticleUpdate)    // 修改文章信息
		auth.GET("/myarticle-list", api.GetMyArticle)     // 查询我的文章

		auth.POST("/tag-add", api.TagAdd) //创建标签

		auth.POST("/comment-add", api.CommentAdd)         //发表评论
		auth.DELETE("/delete-comment", api.CommentDelete) // 删除评论

		auth.POST("/commentlike", api.CommentLike)     //点赞评论
		auth.POST("/commentunlike", api.CommentUnLike) //取消点赞

		auth.POST("/articlelike", api.ArticleLike)     //点赞文章
		auth.POST("/articleunlike", api.ArticleUnLike) //取消点赞文章

		auth.POST("/notice-add", api.MessageAdd)         //发布通知
		auth.DELETE("/delete-notice", api.MessageDelete) // 删除消息
		auth.PUT("/update-notice", api.MessageUpdate)    // 更新通知

		auth.GET("/noticeclub-list", api.GetClubMessageList) // 获取某个社团消息列表
		auth.GET("/mynoticeclub-list", api.GetMyMessage)     // 查询我的社团消息

		auth.POST("/joinfavorite", api.JoinFavorite)        //加入收藏
		auth.POST("/quitfavoorite", api.QuitFavorite)       //移除收藏
		auth.GET("/myfavoritteclub", api.GetMyFavoriteClub) // 查询我收藏的社团

	}

	r.Run(conf.HttpPort)
	return r
}
