package dao

import (
	"github.com/redis/go-redis/v9"
	"school-bbs/conf"
)

var RDB = InitRedisDB()

func InitRedisDB() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     conf.RedisAddr,
		Password: "",
		DB:       conf.RedisDbName,
	})
}
