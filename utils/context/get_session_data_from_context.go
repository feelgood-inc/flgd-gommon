package context

import (
	"github.com/feelgood-inc/flgd-gommon/models"
	"github.com/labstack/echo/v4"
)

func GetSessionDataFormContext(c echo.Context) models.SessionData {
	sessionData := c.Get("session_data")
	if sessionData == nil {
		return models.SessionData{}
	}

	return sessionData.(models.SessionData)
}
