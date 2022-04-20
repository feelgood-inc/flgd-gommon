package authorization

import "github.com/vale-app/vale-common/models"
import "github.com/thoas/go-funk"

func ExtractPayloadFromJWT(token *string) (payload models.JWTClaims, err error) {
	payload, err = ExtractPayloadFromToken(token)
	if err != nil {
		return
	}
	return
}

func ExtractRolesForUser(payload *models.JWTPayload) []string {
	return payload.Roles
}

func DoCurrentRolesContainAnyDesiredRoles(currentRoles, toCheckIfHaveRoles *[]string) bool {
	return funk.Contains(&currentRoles, &toCheckIfHaveRoles)
}

func DoCurrentRolesContainAllDesiredRoles(currentRoles, toCheckIfHaveRoles *[]string) bool {
	iOne, iTwo := funk.Difference(&currentRoles, &toCheckIfHaveRoles)
	if len(iOne) == 0 && len(iTwo) == 0 {
		return true
	}
}
