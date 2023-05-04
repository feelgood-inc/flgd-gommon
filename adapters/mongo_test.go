package adapters

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestCreateMongoClient(t *testing.T) {
	ctx := context.Background()

	mongoConfig := MongoConfig{
		URI:         "mongodb://localhost:27017",
		MaxPoolSize: 100,
		MinPoolSize: 10,
		RetryWrites: true,
	}

	client := CreateMongoClient(ctx, mongoConfig)
	defer func(client *mongo.Client, ctx context.Context) {
		err := client.Disconnect(ctx)
		if err != nil {
			panic(err)
		}
	}(client, ctx)

	err := client.Ping(ctx, nil)
	assert.NoError(t, err)

	dbNames, err := client.ListDatabaseNames(ctx, nil, options.ListDatabases().SetNameOnly(true))
	assert.NoError(t, err)
	assert.NotNil(t, dbNames)
}
