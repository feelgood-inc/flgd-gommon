package adapters

import (
	"context"
	"github.com/go-redis/redis/extra/redisotel/v9"
	"github.com/go-redis/redis/v9"
)

type RedisConfig struct {
	Addr string
	DB   int
	URL  string
}

func CreateRedisClient(ctx context.Context, redisConfig RedisConfig) *redis.Client {
	var redisClient *redis.Client

	if redisConfig.Addr != "" {
		opt, _ := redis.ParseURL(redisConfig.URL)
		redisClient = redis.NewClient(opt)
	} else {
		redisClient = redis.NewClient(&redis.Options{
			Addr: redisConfig.Addr,
			DB:   redisConfig.DB, // use default DB
		})
	}

	redisClient.AddHook(redisotel.NewTracingHook())

	return redisClient
}
