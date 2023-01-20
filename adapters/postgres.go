package adapters

import (
	"context"
	"fmt"
	"github.com/uptrace/opentelemetry-go-extra/otelgorm"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresConfig struct {
	URL      string
	Host     string
	Port     int
	DB       string
	User     string
	Password string
}

func CreatePostgresClient(ctx context.Context, postgresConfig PostgresConfig) *gorm.DB {
	if postgresConfig.URL != "" {
		postgresDB, err := gorm.Open(postgres.Open(postgresConfig.URL), &gorm.Config{
			PrepareStmt: false,
		})
		if err != nil {
			panic(err)
		}
		if err := postgresDB.Use(otelgorm.NewPlugin()); err != nil {
			panic(err)
		}
		return postgresDB
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		postgresConfig.Host,
		postgresConfig.User,
		postgresConfig.Password,
		postgresConfig.DB,
		postgresConfig.Port,
	)
	postgresDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	if err := postgresDB.Use(otelgorm.NewPlugin()); err != nil {
		panic(err)
	}

	return postgresDB
}
