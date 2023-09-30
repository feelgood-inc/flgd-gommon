package middleware

import (
	"context"
	"encoding/base64"
	firebase "firebase.google.com/go/v4"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

type AuthenticationMiddleware struct {
	firebase *firebase.App
}

func NewAuthenticationMiddleware(firebase *firebase.App) *AuthenticationMiddleware {
	return &AuthenticationMiddleware{
		firebase: firebase,
	}
}

func (m *AuthenticationMiddleware) ValidateJWT(ctx context.Context, IDToken string) error {
	client, err := m.firebase.Auth(ctx)
	if err != nil {
		return err
	}
	_, err = client.VerifyIDToken(ctx, IDToken)
	if err != nil {
		return err
	}

	return nil
}

// BasicAuth is a custom middleware for basic authentication.
func BasicAuth(username, password string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Get the "Authorization" header from the request
			auth := c.Request().Header.Get("Authorization")

			// Check if the header is empty
			if auth == "" {
				return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
			}

			// Parse the "Authorization" header
			// It should be in the format "Basic <base64-encoded-username:password>"
			// Extract the base64-encoded part and decode it
			credentials := auth[len("Basic "):]
			decodedCredentials, err := decodeBase64(credentials)
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
			}

			// Split the decoded credentials into username and password
			authParts := strings.Split(decodedCredentials, ":")
			if len(authParts) != 2 {
				return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
			}

			// Check if the provided username and password match the expected values
			if authParts[0] == username && authParts[1] == password {
				return next(c) // Authentication successful, continue to the next middleware or handler
			}

			return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
		}
	}
}

// Decode base64-encoded string
func decodeBase64(encodedStr string) (string, error) {
	decodedBytes, err := base64.StdEncoding.DecodeString(encodedStr)
	if err != nil {
		return "", err
	}
	return string(decodedBytes), nil
}
