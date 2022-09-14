package adapters

import (
	"fmt"
	"github.com/feelgood-inc/flgd-gommon/config"
	"github.com/getsentry/sentry-go"
)

func CreateSentry(cfg *config.Config) error {
	if err := sentry.Init(sentry.ClientOptions{
		Dsn:         cfg.Sentry.SentryDSN,
		Environment: cfg.Env,
	}); err != nil {
		fmt.Printf("Sentry initialization failed: %v\n", err)
		return err
	}

	return nil
}
