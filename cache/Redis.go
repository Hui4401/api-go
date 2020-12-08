package cache

import (
	"context"

	"github.com/go-redis/redis/v8"
)

// Redis缓存客户端单例
var RedisClient *redis.Client

var Context = context.Background()

// 初始化redis连接
func Redis(url string) {
	opt, err := redis.ParseURL(url)
	if err != nil {
		panic("解析redis地址失败: " + err.Error())
	}
	client := redis.NewClient(opt)

	_, err = client.Ping(Context).Result()

	if err != nil {
		panic("redis连接异常: " + err.Error())
	}

	RedisClient = client
}
