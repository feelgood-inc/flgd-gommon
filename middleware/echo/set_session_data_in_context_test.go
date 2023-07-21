package middleware

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/feelgood-inc/flgd-gommon/models"
	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
)

func TestDecodesTokenCorrectly(t *testing.T) {
	const token = "eyJhbGciOiJSUzI1NiIsImtpZCI6ImIyZGZmNzhhMGJkZDVhMDIyMTIwNjM0OTlkNzdlZjRkZWVkMWY2NWIiLCJ0eXAiOiJKV1QifQ.eyJ1c2VyIjp7ImVtYWlsIjoic2Ficmlub2xlZUBnbWFpbC5jb20iLCJmaXJzdF9uYW1lIjoiU2FicmlubyIsImxhc3RfbmFtZSI6IlZhbiBkZSBLYW1wIiwibmF0aW9uYWxfaWQiOiIxNzk1OTcyNjgiLCJuYXRpb25hbGl0eSI6IkNMIiwicm9sZXMiOm51bGwsInR5cGUiOiJwcmFjdGl0aW9uZXIiLCJ1aWQiOiJmdkkxWTl2bVN5VEE3NkdkYVhvRlBFQmZwbG4xIn0sImlzcyI6Imh0dHBzOi8vc2VjdXJldG9rZW4uZ29vZ2xlLmNvbS9mbGdkLWRldiIsImF1ZCI6ImZsZ2QtZGV2IiwiYXV0aF90aW1lIjoxNjg5OTA0NjE1LCJ1c2VyX2lkIjoiZnZJMVk5dm1TeVRBNzZHZGFYb0ZQRUJmcGxuMSIsInN1YiI6ImZ2STFZOXZtU3lUQTc2R2RhWG9GUEVCZnBsbjEiLCJpYXQiOjE2ODk5MDQ2MTUsImV4cCI6MTY4OTkwODIxNSwiZW1haWwiOiJzYWJyaW5vbGVlQGdtYWlsLmNvbSIsImVtYWlsX3ZlcmlmaWVkIjpmYWxzZSwiZmlyZWJhc2UiOnsiaWRlbnRpdGllcyI6eyJlbWFpbCI6WyJzYWJyaW5vbGVlQGdtYWlsLmNvbSJdfSwic2lnbl9pbl9wcm92aWRlciI6ImN1c3RvbSJ9fQ.vGnInSJ9AdEHwHj6NpZvcQm2OaCtA5numq_keSkizq80J_qvjmkketl3-gsiNK0gk4YvXVREWmxOacy6IoE3upK6HsXcHv7j4M-T15jUd1EYtdnyauM1JqejERalAKDm1EP1PB3wySP0oA2B-_2-wpDqU5fwj9Sm3mR6irl-9N1dRLusX12fIiZ-TZwq75XsRBUkYWlzh8doCQbAIwcoqAdG6m_c3_ArEHr_qNxa5S8Q8PFQtEKH0jCurwOcs4X2V9vOa64fUvmYM5APtxtPtJ-C0cQ6AeGskB2-XBN_URZFqzxSDPXg0lf9Pkvds_HmZ4scEyioUWijCvmMazgiww"

	decodedToken, err := jwt.ParseWithClaims(token, &models.FeelgoodJWTClaims{}, nil)

	print(decodedToken.Claims.(*models.FeelgoodJWTClaims).User.Email)

	assert.Error(t, err)
	assert.Equal(t, *decodedToken.Claims.(*models.FeelgoodJWTClaims).User.Email, "sabrinolee@gmail.com")
}

func TestDecodesAuthTokenCorrectly(t *testing.T) {
	const token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyIjp7ImV4cCI6MTY5MDAwMTQ0NCwicm9sZXMiOlsiYWRtaW4iLCJpbnRlcm5hbCJdLCJzZXJ2aWNlX25hbWUiOiJmZy1jb25zdWx0YXRpb25zUGF5bWVudHMtbXMiLCJ1c2VyX3R5cGUiOiJhZG1pbiJ9fQ.i5fANpFHTlbnRt6koLigIygUZTQX-sxuNDaUDDShVNA"

	decodedToken, err := jwt.ParseWithClaims(token, &models.FeelgoodJWTClaims{}, nil)

	print(decodedToken.Claims.(*models.FeelgoodJWTClaims).User.Roles)

	assert.Error(t, err)
	assert.Contains(t, decodedToken.Claims.(*models.FeelgoodJWTClaims).User.Roles, "admin")
	assert.Contains(t, decodedToken.Claims.(*models.FeelgoodJWTClaims).User.Roles, "internal")
}

func TestSetSessionDataInContext_ValidToken(t *testing.T) {
	// Create a new Echo context for testing
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("fg-token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyIjp7ImV4cCI6MTY5MDA1ODUyNSwicm9sZXMiOlsiYWRtaW4iLCJpbnRlcm5hbCJdLCJzZXJ2aWNlX25hbWUiOiJ0ZXN0IiwidHlwZSI6ImFkbWluIn19.Fhkgj5liiPKkTSx38OVzNKNkA3HhUOw-QHBfhNnJGsA") // Replace with a valid JWT token
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	// Call the middleware
	middleware := SetSessionDataInContext()
	err := middleware(func(c echo.Context) error {
		// Check if the session data is set in the context as expected
		sessionData := c.Get(sessionDataKey).(models.SessionData)
		assert.Equal(t, "", sessionData.UID)                                                                                                                                                                                                                // Replace with the expected UID
		assert.Equal(t, "", sessionData.Email)                                                                                                                                                                                                              // Replace with the expected email
		assert.Equal(t, "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyIjp7ImV4cCI6MTY5MDA1ODUyNSwicm9sZXMiOlsiYWRtaW4iLCJpbnRlcm5hbCJdLCJzZXJ2aWNlX25hbWUiOiJ0ZXN0IiwidHlwZSI6ImFkbWluIn19.Fhkgj5liiPKkTSx38OVzNKNkA3HhUOw-QHBfhNnJGsA", sessionData.Token) // Replace with the expected token
		assert.Equal(t, "admin", sessionData.UserType)                                                                                                                                                                                                      // Replace with the expected user type
		assert.Equal(t, []string{"admin", "internal"}, sessionData.UserRoles)                                                                                                                                                                               // Replace with the expected user roles
		return nil
	})(ctx)

	// Check if the middleware executed without errors
	assert.NoError(t, err)
}
