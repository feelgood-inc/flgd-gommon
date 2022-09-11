package middleware

import (
	"github.com/feelgood-inc/flgd-gommon/models"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func SetUserInfo(withKey string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			authHeader := ctx.Request().Header.Get("Authorization")
			if authHeader == "" {
				return ctx.JSON(http.StatusUnauthorized, "Unauthorized")
			}

			decodedToken, _ := jwt.ParseWithClaims(authHeader, &models.FeelgoodJWTClaims{}, nil)
			user := models.User{
				UID:   decodedToken.Claims.(*models.FeelgoodJWTClaims).Claims.UID,
				Email: &decodedToken.Claims.(*models.FeelgoodJWTClaims).Claims.Email,
				Type:  &decodedToken.Claims.(*models.FeelgoodJWTClaims).Claims.Type,
			}
			ctx.Set(withKey, user)

			return next(ctx)
		}
	}
}
