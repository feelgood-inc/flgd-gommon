package adapters

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type MongoConfig struct {
	URI         string
	MaxPoolSize uint64
	MinPoolSize uint64
	RetryWrites bool
}

func CreateMongoClient(ctx context.Context, mongoConfig MongoConfig) *mongo.Client {
	clientOptions := options.Client()

	if &cfg.MongoDB.URI != nil {
		clientOptions.ApplyURI(mongoConfig.URI)
	}
	clientOptions.MaxPoolSize = &mongoConfig.MaxPoolSize
	clientOptions.MinPoolSize = &mongoConfig.MinPoolSize
	clientOptions.RetryWrites = &mongoConfig.RetryWrites

	ctxTimeout, cancel := context.WithTimeout(ctx, 20000*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctxTimeout, clientOptions)
	if err != nil {
		println(err.Error())
		panic(err)
	}

	return client
}
