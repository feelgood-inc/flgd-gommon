package middleware

import (
	"github.com/feelgood-inc/flgd-gommon/models"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func SetUserInfo(ctx echo.Context) {
	authHeader := ctx.Request().Header.Get("Authorization")
	var claims models.FeelgoodJWTCustomClaims
	decodedToken, err := jwt.ParseWithClaims(authHeader, claims, nil)
	if err != nil {
		return
	}
	user := models.User{
		UID:      decodedToken.Claims.(*models.FeelgoodJWTCustomClaims).UID,
		Email:    decodedToken.Claims.(*models.FeelgoodJWTCustomClaims).Email,
		Provider: decodedToken.Claims.(*models.FeelgoodJWTCustomClaims).Provider,
		Type:     decodedToken.Claims.(*models.FeelgoodJWTCustomClaims).Type,
	}
	ctx.Set("user_info", user)
}
