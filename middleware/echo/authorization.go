package middleware

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func AllowToRoles(roles ...string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// get the session data from the context
			sessionData := c.Get("sessionData").(map[string]interface{})
			// get the role from the session data
			role := sessionData["role"].(string)
			// check if the role is allowed
			for _, allowedRole := range roles {
				if role == allowedRole {
					return next(c)
				}
			}
			// return an error if the role is not allowed
			return echo.NewHTTPError(http.StatusForbidden, "You are not allowed to access this resource")
		}
	}
}
