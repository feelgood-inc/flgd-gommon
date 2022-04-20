package authorization

import (
	"github.com/stretchr/testify/assert"
	"github.com/vale-app/vale-common/models"
	"testing"
)

func TestDoCurrentRolesContainAllDesiredRoles(t *testing.T) {
	result := DoCurrentRolesContainAllDesiredRoles(&[]string{"foo", "bar"}, &[]string{"foo", "bar"})
	assert.NotNil(t, result)
	assert.Truef(t, result, "DoCurrentRolesContainAnyDesiredRoles should return true")
}

func TestDoCurrentRolesContainAnyDesiredRoles(t *testing.T) {
	result := DoCurrentRolesContainAnyDesiredRoles(&[]string{"foo", "bar"}, &[]string{"foo", "bar", "gom"})
	assert.NotNil(t, result)
	assert.Truef(t, result, "DoCurrentRolesContainAnyDesiredRoles should return true")
}

func TestExtractRolesForUser(t *testing.T) {
	payload := models.JWTPayload{
		UserID:    "12345",
		Roles:     []string{"foo", "bar"},
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john.doe@gmail.com",
	}
	roles := ExtractRolesForUser(&payload)

	assert.NotNil(t, roles)
	assert.Equal(t, "foo", roles[0])
}
