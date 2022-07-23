package adapters

import (
	"context"
	"github.com/feelgood-inc/flgd-gommon/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func CreateMongoClient(ctx context.Context, dbOptions *models.DBOptions) *mongo.Client {
	clientOptions := options.Client()

	if dbOptions.URI != nil {
		clientOptions.ApplyURI(*dbOptions.URI)
	}
	clientOptions.MaxPoolSize = dbOptions.MaxPoolSize
	clientOptions.MinPoolSize = dbOptions.MinPoolSize
	clientOptions.RetryWrites = dbOptions.RetryWrites
	clientOptions.RetryWrites = dbOptions.RetryWrites

	ctxTimeout, cancel := context.WithTimeout(ctx, 20000*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctxTimeout, clientOptions)
	if err != nil {
		println(err.Error())
		panic(err)
	}

	return client
}
