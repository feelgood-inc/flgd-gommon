package adapters

import (
	"context"
	"github.com/feelgood-inc/flgd-gommon/config"
	"github.com/go-redis/redis/extra/redisotel/v9"
	"github.com/go-redis/redis/v9"
	"os"
)

func CreateRedisClient(ctx context.Context, cfg *config.Config) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_ADDR"),
		DB:   cfg.Redis.DB, // use default DB
	})
	rdb.AddHook(redisotel.NewTracingHook())

	return rdb
}
