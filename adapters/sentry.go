package adapters

import (
	"fmt"
	"github.com/getsentry/sentry-go"
)

type SentryConfig struct {
	DSN                 string
	Env                 string
	TracesSampleRate    float64
	TracesSampler       sentry.TracesSampler
	EnableTracing       bool
	EnableProfiling     bool
	ProfilingSampleRate float64
}

func CreateSentry(sentryConfig SentryConfig) error {
	options := sentry.ClientOptions{}
	options.Dsn = sentryConfig.DSN
	options.Environment = sentryConfig.Env
	options.TracesSampleRate = sentryConfig.TracesSampleRate

	if sentryConfig.EnableTracing {
		options.EnableTracing = true
		options.TracesSampleRate = sentryConfig.TracesSampleRate
	}

	if sentryConfig.EnableProfiling {
		options.ProfilesSampleRate = sentryConfig.ProfilingSampleRate
	}

	if err := sentry.Init(options); err != nil {
		fmt.Printf("Sentry initialization failed: %v\n", err)
		return err
	}

	return nil
}
