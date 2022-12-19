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
			var jwtToken string
			authCookie, err := ctx.Request().Cookie("Authorization")
			if err != nil {
				return ctx.JSON(http.StatusUnauthorized, "Unauthorized")
			} else {
				jwtToken = authCookie.Value
			}

			authHeader := ctx.Request().Header.Get("Authorization")
			if authHeader == "" {
				return ctx.JSON(http.StatusUnauthorized, "Unauthorized")
			} else {
				jwtToken = authHeader
			}

			decodedToken, _ := jwt.ParseWithClaims(jwtToken, &models.FeelgoodJWTClaims{}, nil)
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
