package httpclient

import (
	resty "github.com/go-resty/resty/v2"
	"time"
)

func Client(clientConfig *ClientConfig) *resty.Client {
	return resty.New().SetBaseURL(clientConfig.Host + ":" + clientConfig.Port).SetHeaders(map[string]string{
		"Content-Type": "application/json",
		"Accept":       "application/json",
		"User-Agent":   "vale-resty-client",
	}).SetRetryCount(2).SetRetryWaitTime(2 * time.Second)
}
