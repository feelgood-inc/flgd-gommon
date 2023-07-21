package middleware

import (
	"testing"

	"github.com/feelgood-inc/flgd-gommon/models"
	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
)

func TestDecodesTokenCorrectly(t *testing.T) {
	const token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjbGFpbXMiOnsiZXhwIjoxNjg0NzMwMTQ3LCJyb2xlcyI6WyJhZG1pbiIsImludGVybmFsIl0sInNlcnZpY2VfbmFtZSI6InRlc3QiLCJ1c2VyX3R5cGUiOiJhZG1pbiJ9fQ.LmLpBfGFTJdtyNJucAaD2fikU6SlSQj90flGctXDKJs"

	decodedToken, err := jwt.ParseWithClaims(token, &models.FeelgoodJWTClaims{}, nil)

	print(decodedToken.Claims.(*models.FeelgoodJWTClaims).User.Roles)

	assert.Error(t, err)
	assert.Contains(t, decodedToken.Claims.(*models.FeelgoodJWTClaims).User.Roles, "admin")
}
