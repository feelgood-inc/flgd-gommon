package middleware

import (
	"github.com/labstack/echo/v4"
)

func SetAuthorizationHeaderInCtx() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			authHeader := ctx.Request().Header.Get("Authorization")
			if authHeader == "" {
				return next(ctx)
			}

			ctx.Set("UserAuthToken", authHeader)

			return next(ctx)
		}
	}
}
