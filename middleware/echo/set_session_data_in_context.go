package middleware

import (
	"github.com/feelgood-inc/flgd-gommon/models"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

const (
	sessionDataKey = "session_data"
)

// SetSessionDataInContext sets the user in the request context
// It does not fail or return an error in case the token is not found
func SetSessionDataInContext() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			// The header is set at the proxy level, so we can trust it
			authHeader := ctx.Request().Header.Get("fg-token")

			// In case there is no header, we just skip the middleware and return to the main flow
			if authHeader == "" {
				return next(ctx)
			}

			decodedToken, _ := jwt.ParseWithClaims(authHeader, &models.FeelgoodJWTClaims{}, nil)
			var email, userType string
			if decodedToken.Claims.(*models.FeelgoodJWTClaims).User.Email == nil {
				email = ""
			} else {
				email = *decodedToken.Claims.(*models.FeelgoodJWTClaims).User.Email
			}
			if decodedToken.Claims.(*models.FeelgoodJWTClaims).User.Type == nil {
				userType = ""
			} else {
				userType = *decodedToken.Claims.(*models.FeelgoodJWTClaims).User.Type
			}
			sessionData := models.SessionData{
				UID:       decodedToken.Claims.(*models.FeelgoodJWTClaims).User.UID,
				Email:     email,
				Token:     authHeader,
				UserType:  userType,
				UserRoles: decodedToken.Claims.(*models.FeelgoodJWTClaims).User.Roles,
			}
			ctx.Set(sessionDataKey, sessionData)

			return next(ctx)
		}
	}
}
