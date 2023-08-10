package middleware

import (
	"github.com/feelgood-inc/flgd-gommon/models"
	"github.com/labstack/echo/v4"
	"github.com/samber/lo"
	"net/http"
)

func AllowUserRoles(roles []string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			user, ok := ctx.Request().Context().Value("user").(models.User)
			if !ok {
				return ctx.JSON(http.StatusUnauthorized, "Unauthorized")
			}

			for _, role := range roles {
				// Check if the user roles contains one of the allowed roles
				_, ok := lo.Find(user.Roles, func(s string) bool {
					return s == role
				})

				if ok {
					return next(ctx)
				}
			}

			return ctx.JSON(http.StatusUnauthorized, "Unauthorized")
		}
	}
}

func AllowRoles(roles []string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			sessionData, ok := ctx.Get("session_data").(models.SessionData)
			if !ok {
				return ctx.JSON(http.StatusUnauthorized, "Unauthorized")
			}

			for _, role := range roles {
				// Check if the user roles contains one of the allowed roles
				_, ok := lo.Find(sessionData.UserRoles, func(s string) bool {
					return s == role
				})

				if ok {
					return next(ctx)
				}
			}

			return ctx.JSON(http.StatusUnauthorized, "Unauthorized")
		}
	}
}
