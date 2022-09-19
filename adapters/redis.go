package adapters

import (
	"context"
	"github.com/go-redis/redis/extra/redisotel/v9"
	"github.com/go-redis/redis/v9"
)

type RedisConfig struct {
	Addr string
	DB   int
}

func CreateRedisClient(ctx context.Context, redisConfig RedisConfig) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: redisConfig.Addr,
		DB:   redisConfig.DB, // use default DB
	})
	rdb.AddHook(redisotel.NewTracingHook())

	return rdb
}
