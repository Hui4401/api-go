package redis

import (
    "context"

    "github.com/Hui4401/gopkg/logs"
    "github.com/go-redis/redis/v8"
)

var client *redis.Client

func InitRedis(url string) {
	ctx := context.Background()
	opt, err := redis.ParseURL(url)
	if err != nil {
		logs.PanicKvs("parse redis url error", err)
	}
	c := redis.NewClient(opt)

	_, err = c.Ping(ctx).Result()
	if err != nil {
		logs.PanicKvs("connect to redis error", err)
	}

	client = c
}

func GetClient() *redis.Client {
	return client
}
