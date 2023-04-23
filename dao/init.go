package dao

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"school-bbs/conf"
	"school-bbs/model"
	"time"
)

var (
	DB *gorm.DB
)

func init() {
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.DbUser,
		conf.DbPassWord,
		conf.DbHost,
		conf.DbPort,
		conf.DbName,
	)
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{
		// gorm日志模式：silent
		Logger: logger.Default.LogMode(logger.Silent),
		// 外键约束
		DisableForeignKeyConstraintWhenMigrating: true,
		// 禁用默认事务（提高运行速度）
		SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			// 使用单数表名，启用该选项，此时，`User` 的表名应该是 `user`
			SingularTable: true,
		},
	})

	if err != nil {
		log.Println("连接数据库失败，请检查参数：", err)
	} else {
		log.Printf("连接数据库成功")

	}

	sqlDB, _ := db.DB()
	// SetMaxIdleCons 设置连接池中的最大闲置连接数。
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenCons 设置数据库的最大连接数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetiment 设置连接的最大可复用时间。
	sqlDB.SetConnMaxLifetime(10 * time.Second)
	DB = db
	// 迁移数据表，在没有数据表结构变更时候，建议注释不执行
	err = DB.AutoMigrate(&model.User{},
		&model.Club{},
		&model.Article{},
		&model.Active{},
		&model.ActiveImg{},
		&model.Tag{},
		&model.ArticleTag{},
		&model.ArticleImg{},
		&model.Comment{},
		&model.ClubMembers{},
		&model.Notice{},
		&model.FavoriteClub{},
		&model.ClubImg{},
		&model.CommentLike{},
		&model.UserFans{},
		&model.ArticleLike{},
	)
	if err != nil {
		log.Printf("迁移数据表失败")
	} else {
		log.Printf("迁移数据表成功")

	}

}
