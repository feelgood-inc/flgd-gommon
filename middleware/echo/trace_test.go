package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestProcess(t *testing.T) {
	t.Run("trace middleware add response header to success request", func(t *testing.T) {
		e := echo.New()
		e.Use(Trace)

		e.GET("/foo", func(c echo.Context) error {
			// Do some task
			return c.JSON(http.StatusOK, echo.Map{"foo": "bar"})
		})

		req := httptest.NewRequest(http.MethodGet, "/foo", nil)
		res := httptest.NewRecorder()
		e.ServeHTTP(res, req)

		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, "{\"foo\":\"bar\"}\n", res.Body.String())

		xTraceID := res.Result().Header.Get(headerTraceID)
		assert.NotEmpty(t, xTraceID, "Should NOT be empty, but was \"%v\"", xTraceID)
	})

	t.Run("trace middleware add response header to non successful request", func(t *testing.T) {
		e := echo.New()
		e.Use(Trace)

		e.GET("/foo", func(c echo.Context) error {
			// Do some task and it failed
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": "haber nacido"})
		})

		req := httptest.NewRequest(http.MethodGet, "/foo", nil)
		res := httptest.NewRecorder()
		e.ServeHTTP(res, req)

		assert.Equal(t, http.StatusInternalServerError, res.Code)
		assert.Equal(t, "{\"error\":\"haber nacido\"}\n", res.Body.String())

		xTraceID := res.Result().Header.Get(headerTraceID)
		assert.NotEmpty(t, xTraceID, "Should NOT be empty, but was \"%v\"", xTraceID)
	})
}
