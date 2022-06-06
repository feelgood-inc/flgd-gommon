package middleware

import (
	"github.com/feelgood-inc/flgd-gommon/models"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func SetUserInfo(withKey string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			authHeader := ctx.Request().Header.Get("Authorization")
			decodedToken, err := jwt.ParseWithClaims(authHeader, &models.FeelgoodJWTCustomClaims{}, nil)
			if err != nil {
				return ctx.JSON(401, map[string]string{"error": "Invalid token"})
			}
			user := models.User{
				UID:      decodedToken.Claims.(*models.FeelgoodJWTCustomClaims).UID,
				Email:    decodedToken.Claims.(*models.FeelgoodJWTCustomClaims).Email,
				Provider: decodedToken.Claims.(*models.FeelgoodJWTCustomClaims).Provider,
				Type:     decodedToken.Claims.(*models.FeelgoodJWTCustomClaims).Type,
			}
			ctx.Set(withKey, user)

			return next(ctx)
		}
	}
}
