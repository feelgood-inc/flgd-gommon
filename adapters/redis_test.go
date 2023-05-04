package adapters

import (
	"context"
	"github.com/redis/go-redis/v9"
	"testing"

	"github.com/go-redis/redismock/v9"
	"github.com/stretchr/testify/assert"
)

func TestCreateRedisClient(t *testing.T) {
	ctx := context.Background()

	redisConfig := RedisConfig{
		Addr: "localhost:6379",
		DB:   0,
		URL:  "",
	}

	// create a mock redis client
	client, mock := redismock.NewClientMock()
	defer client.Close()

	// set up expectations
	mock.ExpectPing()

	// call CreateRedisClient with the mock client
	rdb := CreateRedisClient(ctx, redisConfig)
	defer func(rdb *redis.Client) {
		err := rdb.Close()
		if err != nil {
			t.Errorf("error closing redis client: %s", err)
		}
	}(rdb)

	// assert that ping succeeded
	err := mock.ExpectationsWereMet()
	assert.NoError(t, err)
}
