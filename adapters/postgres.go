package adapters

import (
	"context"
	"fmt"
	"github.com/feelgood-inc/flgd-gommon/config"
	"github.com/uptrace/opentelemetry-go-extra/otelgorm"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func CreatePostgresClient(ctx context.Context, cfg *config.Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		cfg.Postgres.Host,
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		cfg.Postgres.DB,
		cfg.Postgres.Port,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	if err := db.Use(otelgorm.NewPlugin()); err != nil {
		panic(err)
	}

	return db
}
