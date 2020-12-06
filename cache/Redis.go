package cache

import (
	"strconv"

	"github.com/go-redis/redis"
)

// Redis缓存客户端单例
var RedisClient *redis.Client

// 初始化redis连接
func Redis(addr string, password string, db string) {
	dbNum, _ := strconv.Atoi(db)
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       dbNum,
	})

	_, err := client.Ping().Result()

	if err != nil {
		panic("redis连接异常: " + err.Error())
	}

	RedisClient = client
}
