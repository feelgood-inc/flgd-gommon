package authorization

import "github.com/vale-app/vale-common/models"
import "github.com/thoas/go-funk"

func ExtractRolesForUser(payload *models.JWTPayload) []string {
	return payload.Roles
}

func DoCurrentRolesContainAnyDesiredRoles(currentRoles, toCheckIfHaveRoles *[]string) bool {
	return funk.Contains(&currentRoles, &toCheckIfHaveRoles)
}
