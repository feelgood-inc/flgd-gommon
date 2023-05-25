package utils

import (
	"github.com/dgraph-io/ristretto"
	"net/http"
)

func CheckIfRequestIsAuthenticated(req *http.Request) (bool, *string) {
	isAuthenticated := false

	// Check first if there is a token in the header
	authHeader := req.Header.Get("Authorization")
	if authHeader != "" {
		isAuthenticated = true
		println("Token in header")
		println(authHeader)
		return isAuthenticated, &authHeader
	}

	// Check if there is a token in the cookie
	cookie, err := req.Cookie("Authorization")
	if err != nil {
		return isAuthenticated, nil
	}
	if cookie.String() != "" {
		isAuthenticated = true
		println("Token in cookie")
		println(cookie.String())
		return isAuthenticated, &cookie.Value
	}

	return isAuthenticated, nil
}

func SetInternalAuthTokenInLocalCache(localCache *ristretto.Cache, token string) {
	localCache.Set("internal_auth_token", token, 1)
}
