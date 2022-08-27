package adapters

import (
	"fmt"
	"github.com/feelgood-inc/flgd-gommon/config"
	"github.com/getsentry/sentry-go"
	"os"
)

func CreateSentry(cfg *config.Config) error {
	if err := sentry.Init(sentry.ClientOptions{
		Dsn:         os.Getenv("SENTRY_DSN"),
		Environment: os.Getenv("ENV"),
	}); err != nil {
		fmt.Printf("Sentry initialization failed: %v\n", err)
		return err
	}

	return nil
}
