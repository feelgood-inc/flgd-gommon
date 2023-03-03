package middleware

import (
	"github.com/feelgood-inc/flgd-gommon/models"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

// SetUserInContext sets the user in the request context
// It does not fail or return an error in case the token is not found
func SetUserInContext(withKey string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			// The header is set at the proxy level, so we can trust it
			authHeader := ctx.Request().Header.Get("fg-token")

			// In case there is no header, we just skip the middleware and return to the main flow
			if authHeader == "" {
				return next(ctx)
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
