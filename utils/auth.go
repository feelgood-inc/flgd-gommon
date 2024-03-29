package utils

import (
	"context"
	"fmt"
	"github.com/dgraph-io/ristretto"
	"github.com/feelgood-inc/flgd-gommon/models"
	"github.com/go-resty/resty/v2"
	"github.com/spf13/viper"
	"net/http"
	"time"
)

type GetInternalTokenConfig struct {
	RetryCount    int
	RetryWaitTime time.Duration
	ServiceName   string
	Debug         bool
	AuthMSHost    string
}

func CheckIfRequestIsAuthenticated(req *http.Request) (bool, *string) {
	isAuthenticated := false

	// Check first if there is a token in the header
	authHeader := req.Header.Get("Authorization")
	if authHeader != "" {
		isAuthenticated = true
		return isAuthenticated, &authHeader
	}

	// Check if there is a token in the cookie
	cookie, err := req.Cookie("Authorization")
	if err != nil {
		return isAuthenticated, nil
	}
	if cookie.String() != "" {
		isAuthenticated = true
		return isAuthenticated, &cookie.Value
	}

	return isAuthenticated, nil
}

func SetInternalAuthTokenInLocalCache(localCache *ristretto.Cache, token string) {
	localCache.Set("internal_auth_token", token, 1)
}

func GetInternalAuthTokenFromLocalCache(localCache *ristretto.Cache) string {
	token, found := localCache.Get("internal_auth_token")
	if !found {
		return ""
	}
	return token.(string)
}

func GetInternalToken(ctx context.Context, config *GetInternalTokenConfig) (string, error) {
	// Get username and password from environment variables
	username := viper.GetString("INTERNAL_USER")
	if username == "" {
		panic("INTERNAL_USER environment variable not set")
	}

	password := viper.GetString("INTERNAL_PASSWORD")
	if password == "" {
		panic("INTERNAL_PASSWORD environment variable not set")
	}

	customClient := resty.New().
		SetHeaders(map[string]string{
			"Content-Type":     "application/json",
			"Accept":           "application/json",
			"User-Agent":       "flgd-resty-client",
			"X-Application-ID": config.ServiceName,
		}).
		SetRetryCount(config.RetryCount).
		SetRetryWaitTime(config.RetryWaitTime)

	if config.Debug {
		customClient.SetDebug(true)
	}

	// Obtain the token from the auth microservice
	resp, err := customClient.
		R().
		SetContext(ctx).
		SetBody(map[string]string{
			"service_name": config.ServiceName,
			"user":         username,
			"password":     password,
		}).
		Post(config.AuthMSHost + "/auth/login-internal-ms")

	if err != nil {
		return "", err
	}

	if resp.IsError() {
		return "", fmt.Errorf("failed with status code: %d and body: %s", resp.StatusCode(), resp.Body())
	}

	var loginData models.LoginData
	err = FeelgoodResponseToStruct(resp, &loginData)
	if err != nil {
		return "", err
	}

	return loginData.Token, nil
}
