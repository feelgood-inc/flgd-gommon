package middleware

import (
	"context"
	"github.com/feelgood-inc/flgd-gommon/enums"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/feelgood-inc/flgd-gommon/models"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestAllowUserRoles(t *testing.T) {
	// Create a user with the role "admin"
	user := models.User{
		Roles: []string{"admin"},
	}

	// Create an echo context with the user
	e := echo.New()

	e.Use(AllowUserRoles([]string{string(enums.RoleAdmin)}))

	e.GET("/foo", func(c echo.Context) error {
		// Do some task
		return c.JSON(http.StatusOK, echo.Map{"foo": "bar"})
	})

	req := httptest.NewRequest(http.MethodGet, "/foo", nil)
	res := httptest.NewRecorder()

	c := context.WithValue(req.Context(), "user", user)
	req = req.WithContext(c)

	e.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)
	assert.Equal(t, "{\"foo\":\"bar\"}\n", res.Body.String())
}

func TestAllowUserRolesUnauthorized(t *testing.T) {
	// Create a user with no roles
	user := models.User{}

	// Create an echo context with the user
	e := echo.New()
	ctx := e.NewContext(nil, nil)
	ctx.Set("user", user)

	e.Use(AllowUserRoles([]string{"admin"}))

	e.GET("/foo", func(c echo.Context) error {
		// Do some task
		return c.JSON(http.StatusOK, echo.Map{"foo": "bar"})
	})

	req := httptest.NewRequest(http.MethodGet, "/foo", nil)
	res := httptest.NewRecorder()

	c := context.WithValue(req.Context(), "user", user)
	req = req.WithContext(c)

	e.ServeHTTP(res, req)

	assert.Equal(t, http.StatusUnauthorized, res.Code)
	assert.Equal(t, "\"Unauthorized\"\n", res.Body.String())
}
