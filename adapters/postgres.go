package adapters

import (
	"context"
	"fmt"
	"github.com/feelgood-inc/flgd-gommon/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func CreatePostgresClient(ctx context.Context, dbOptions *models.DBOptions) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=%s",
		*dbOptions.Host,
		*dbOptions.User,
		*dbOptions.Password,
		dbOptions.DBName,
		dbOptions.Port,
		*dbOptions.TimeZone,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}
