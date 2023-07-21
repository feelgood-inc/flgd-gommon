package middleware

import (
	"net/http"

	"github.com/feelgood-inc/flgd-gommon/models"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func SetUserInfo(withKey string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			// The header is set at the proxy level, so we can trust it
			authHeader := ctx.Request().Header.Get("fg-token")
			if authHeader == "" {
				return ctx.JSON(http.StatusUnauthorized, "Unauthorized")
			}

			decodedToken, _ := jwt.ParseWithClaims(authHeader, &models.FeelgoodJWTClaims{}, nil)
			user := models.User{
				UID:   decodedToken.Claims.(*models.FeelgoodJWTClaims).User.UID,
				Email: decodedToken.Claims.(*models.FeelgoodJWTClaims).User.Email,
				Type:  decodedToken.Claims.(*models.FeelgoodJWTClaims).User.Type,
			}
			ctx.Set(withKey, user)

			return next(ctx)
		}
	}
}
