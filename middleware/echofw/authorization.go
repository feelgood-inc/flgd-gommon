package echofw

import (
	"github.com/labstack/echo/v4"
	"github.com/vale-app/vale-common/utils/authorization"
	"github.com/vale-app/vale-common/utils/jwt"
)

func UserContainsAllRoles(roles *[]string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			decodedToken, err := jwt.DecodeJWT(&authHeader)
			if err != nil {
				return c.JSON(401, map[string]string{"error": "Invalid token"})
			}

			payload, err := authorization.ExtractPayloadFromJWT(&decodedToken)
			if err != nil {
				return c.JSON(401, map[string]string{"error": "Invalid token"})
			}

			if !authorization.DoCurrentRolesContainAllDesiredRoles(&payload, roles) {
				return c.JSON(401, map[string]string{"error": "Invalid token"})
			}

			return next(c)
		}
	}
}

func UserContainsAnyRole(roles *[]string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			decodedToken, err := jwt.DecodeJWT(&authHeader)
			if err != nil {
				return c.JSON(401, map[string]string{"error": "Invalid token"})
			}

			payload, err := authorization.ExtractPayloadFromJWT(&decodedToken)
			if err != nil {
				return c.JSON(401, map[string]string{"error": "Invalid token"})
			}

			if !authorization.DoCurrentRolesContainAnyDesiredRoles(&payload, roles) {
				return c.JSON(401, map[string]string{"error": "Invalid token"})
			}

			return next(c)
		}
	}
}
