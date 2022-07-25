package httpclient

import (
	resty "github.com/go-resty/resty/v2"
	"os"
	"time"
)

func Client(clientConfig *ClientConfig) *resty.Client {
	return resty.New().SetBaseURL(clientConfig.Host + ":" + clientConfig.Port).SetHeaders(map[string]string{
		"Content-Type": "application/json",
		"Accept":       "application/json",
		"User-Agent":   "flgd-resty-client",
	}).SetRetryCount(2).SetRetryWaitTime(2 * time.Second)
}

func Default() *resty.Client {
	internalDNSURL := os.Getenv("VALE_INTERNAL_URL")
	if internalDNSURL == "" {
		panic("VALE_INTERNAL_URL is not set")
	}
	return resty.New().
		SetBaseURL(internalDNSURL).
		SetHeaders(map[string]string{
			"Content-Type": "application/json",
			"Accept":       "application/json",
			"User-Agent":   "flgd-resty-client",
		}).
		SetRetryCount(3).
		SetRetryWaitTime(1 * time.Second)
}
