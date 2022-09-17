package adapters

import (
	"context"
	"github.com/feelgood-inc/flgd-gommon/config"
	"github.com/go-redis/redis/extra/redisotel/v9"
	"github.com/go-redis/redis/v9"
)

func CreateRedisClient(ctx context.Context, cfg *config.Config) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.RedisAddr,
		Password: cfg.Redis.RedisPassword, // no password set
		DB:       cfg.Redis.DB,            // use default DB
	})
	rdb.AddHook(redisotel.NewTracingHook())

	return rdb
}
