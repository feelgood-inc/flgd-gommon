package middleware

import (
	"github.com/feelgood-inc/flgd-gommon/utils"
	"github.com/labstack/echo/v4"
)

func SetAuthHeaders() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// check if request is authenticated
			isAuthenticated, token := utils.CheckIfRequestIsAuthenticated(c.Request())

			// Enrich the request with the token and authentication information
			if isAuthenticated {
				c.Request().Header.Set("fg-authenticated", "true")
				c.Request().Header.Set("fg-token", *token)
				c.Request().Header.Set("fg-uid", "test")
			} else {
				c.Request().Header.Set("fg-authenticated", "false")
			}

			return next(c)
		}
	}
}
