package adapters

import (
	"fmt"
	"github.com/getsentry/sentry-go"
)

type SentryConfig struct {
	DSN              string
	Env              string
	TracesSampleRate float64
	TracesSampler    sentry.TracesSampler
	EnableTracing    bool
}

func CreateSentry(sentryConfig SentryConfig) error {
	if err := sentry.Init(sentry.ClientOptions{
		Dsn:              sentryConfig.DSN,
		Environment:      sentryConfig.Env,
		TracesSampleRate: sentryConfig.TracesSampleRate,
		EnableTracing:    sentryConfig.EnableTracing,
	}); err != nil {
		fmt.Printf("Sentry initialization failed: %v\n", err)
		return err
	}

	return nil
}
