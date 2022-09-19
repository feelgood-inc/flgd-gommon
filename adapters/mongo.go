package adapters

import (
	"context"
	"github.com/feelgood-inc/flgd-gommon/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"time"
)

func CreateMongoClient(ctx context.Context, cfg *config.Config) *mongo.Client {
	clientOptions := options.Client()

	if &cfg.MongoDB.URI != nil {
		clientOptions.ApplyURI(os.Getenv("MONGODB_URI"))
	}
	clientOptions.MaxPoolSize = &cfg.MongoDB.MaxPoolSize
	clientOptions.MinPoolSize = &cfg.MongoDB.MinPoolSize
	clientOptions.RetryWrites = &cfg.MongoDB.RetryWrites

	ctxTimeout, cancel := context.WithTimeout(ctx, 20000*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctxTimeout, clientOptions)
	if err != nil {
		println(err.Error())
		panic(err)
	}

	return client
}
