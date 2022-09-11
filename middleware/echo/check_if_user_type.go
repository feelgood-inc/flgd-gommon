package middleware

import (
	"github.com/feelgood-inc/flgd-gommon/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

// CheckIfUserIsOfType checks if the user is of the given type
// It depends on the user info to be set in the context
// TODO: eventually, change for https://golangtutorial.dev/tips/golang-slice-contains-method/
func CheckIfUserIsOfType(userTypes ...string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			user, ok := ctx.Get("user").(models.User)
			if !ok {
				return ctx.JSON(http.StatusUnauthorized, "Unauthorized")
			}

			for _, userType := range userTypes {
				if userType == *user.Type {
					return next(ctx)
				}
			}
			return ctx.JSON(http.StatusUnauthorized, "Unauthorized")
		}
	}
}
