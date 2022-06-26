package model

import (
    "context"

    "github.com/go-redis/redis/v8"

    localRedis "api-go/storage/redis"
)

type jwtDao struct {
    redisClient *redis.Client
}

func NewJwtDao() *jwtDao {
    return &jwtDao{
        redisClient: localRedis.GetClient(),
    }
}

func makeJwtBanedKey() string {
    return "jwt:baned"
}

func (d *jwtDao) BanToken(ctx context.Context, token string) error {
    key := makeJwtBanedKey()
    return d.redisClient.SAdd(ctx, key, token).Err()
}

func (d *jwtDao) IsBanedToken(ctx context.Context, token string) bool {
    key := makeJwtBanedKey()
    return d.redisClient.SIsMember(ctx, key, token).Val()
}
