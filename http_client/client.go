package httpclient

import (
	"context"
	"fmt"
	"github.com/feelgood-inc/flgd-gommon/models"
	"os"
	"time"

	"github.com/feelgood-inc/flgd-gommon/config"
	resty "github.com/go-resty/resty/v2"
)

type InternalClientConfig struct {
	RetryCount       int
	RetryWaitTime    time.Duration
	ServiceName      string
	Debug            bool
	AuthMSHost       string
	WithInternalAuth bool
}

func Client(clientConfig *ClientConfig) *resty.Client {
	return resty.New().SetBaseURL(clientConfig.Host + ":" + clientConfig.Port).SetHeaders(map[string]string{
		"Content-Type": "application/json",
		"Accept":       "application/json",
		"User-Agent":   "flgd-resty-client",
	}).SetRetryCount(2).SetRetryWaitTime(2 * time.Second)
}

func Default() *resty.Client {
	internalDNSURL := os.Getenv("FG_INTERNAL_URL")
	if internalDNSURL == "" {
		panic("FG_INTERNAL_URL is not set")
	}
	return resty.New().
		SetBaseURL(internalDNSURL).
		SetHeaders(map[string]string{
			"Content-Type":     "application/json",
			"Accept":           "application/json",
			"User-Agent":       "flgd-resty-client",
			"X-Application-ID": os.Getenv("PROJECT_NAME"),
		}).
		SetRetryCount(3).
		SetRetryWaitTime(1 * time.Second)
}

func Internal(ctx context.Context, cfg *InternalClientConfig) (*resty.Client, error) {
	customClient := resty.New().
		SetHeaders(map[string]string{
			"Content-Type":     "application/json",
			"Accept":           "application/json",
			"User-Agent":       "flgd-resty-client",
			"X-Application-ID": cfg.ServiceName,
		}).
		SetRetryCount(cfg.RetryCount).
		SetRetryWaitTime(cfg.RetryWaitTime)

	if cfg.Debug {
		customClient.SetDebug(true)
	}

	if !cfg.WithInternalAuth {
		return customClient, nil
	}

	// Obtain the token from the auth microservice
	resp, err := customClient.
		R().
		SetContext(ctx).
		SetBody(map[string]string{
			"service_name": cfg.ServiceName,
		}).
		Post("/login-internal-ms")

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, fmt.Errorf("failed with status code: %d and body: %s", resp.StatusCode(), resp.Body())
	}

	var loginData models.LoginData
	err = FeelgoodResponseToStruct(resp, &loginData)
	if err != nil {
		return nil, err
	}

	// Assign the token to the client
	customClient.SetHeader("Authorization", loginData.Token)

	return customClient, err
}

func External(cfg *config.Config) *resty.Client {
	return resty.New().
		// SetTransport(otelhttp.NewTransport(http.DefaultTransport)).
		SetHeaders(map[string]string{
			"Content-Type": "application/json",
			"Accept":       "application/json",
			"User-Agent":   "flgd-resty-client",
		}).
		SetRetryCount(3).
		SetRetryWaitTime(1 * time.Second)
}
