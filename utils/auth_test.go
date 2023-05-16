package utils_test

import (
	"github.com/feelgood-inc/flgd-gommon/utils"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckIfRequestIsAuthenticated_HeaderToken(t *testing.T) {
	// Create a new request with an Authorization header
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("Authorization", "Bearer TOKEN123")

	// Call the function
	isAuthenticated, token := utils.CheckIfRequestIsAuthenticated(req)

	// Assert that the request is authenticated and the token is correct
	assert.True(t, isAuthenticated)
	assert.NotNil(t, token)
	assert.Equal(t, "Bearer TOKEN123", *token)
}

func TestCheckIfRequestIsAuthenticated_CookieToken(t *testing.T) {
	// Create a new request with an Authorization cookie
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	cookie := &http.Cookie{
		Name:  "Authorization",
		Value: "TOKEN123",
	}
	req.AddCookie(cookie)

	// Call the function
	isAuthenticated, token := utils.CheckIfRequestIsAuthenticated(req)

	// Assert that the request is authenticated and the token is correct
	assert.True(t, isAuthenticated)
	assert.NotNil(t, token)
	assert.Equal(t, "TOKEN123", *token)
}

func TestCheckIfRequestIsAuthenticated_NoToken(t *testing.T) {
	// Create a new request without an Authorization header or cookie
	req := httptest.NewRequest(http.MethodGet, "/", nil)

	// Call the function
	isAuthenticated, token := utils.CheckIfRequestIsAuthenticated(req)

	// Assert that the request is not authenticated and the token is nil
	assert.False(t, isAuthenticated)
	assert.Nil(t, token)
}
