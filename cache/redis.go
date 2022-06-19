package cache

import (
    "context"

    "github.com/go-redis/redis/v8"
    "api-go/logs"
)

var Redis *redis.Client

func InitRedis(url string) {
    ctx := context.Background()
    opt, err := redis.ParseURL(url)
    if err != nil {
        logs.PanicKvs("parse redis url error", err)
    }
    client := redis.NewClient(opt)

    _, err = client.Ping(ctx).Result()
    if err != nil {
        logs.PanicKvs("connect to redis error", err)
    }

    Redis = client
}
