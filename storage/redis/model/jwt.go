package model

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"

	localRedis "api-go/storage/redis"
)

type JwtDao struct {
	redisClient *redis.Client
}

func NewJwtDao() *JwtDao {
	return &JwtDao{
		redisClient: localRedis.GetClient(),
	}
}

func makeJwtBanedKey(token string) string {
	return "token-baned:" + token
}

func (d *JwtDao) BanToken(ctx context.Context, token string, expireTime time.Duration) error {
	key := makeJwtBanedKey(token)
	return d.redisClient.Set(ctx, key, nil, expireTime).Err()
}

func (d *JwtDao) IsBanedToken(ctx context.Context, token string) bool {
	key := makeJwtBanedKey(token)
	if count := d.redisClient.Exists(ctx, key).Val(); count > 0 {
		return true
	}

	return false
}
