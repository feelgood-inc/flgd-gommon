package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

const token = "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJmaXJlYmFzZS1hZG1pbnNkay1waXkwN0BmbGdkLTU3NDkyLmlhbS5nc2VydmljZWFjY291bnQuY29tIiwiYXVkIjoiaHR0cHM6Ly9pZGVudGl0eXRvb2xraXQuZ29vZ2xlYXBpcy5jb20vZ29vZ2xlLmlkZW50aXR5LmlkZW50aXR5dG9vbGtpdC52MS5JZGVudGl0eVRvb2xraXQiLCJleHAiOjE2NTY3ODUxNzgsImlhdCI6MTY1Njc4MTU3OCwic3ViIjoiZmlyZWJhc2UtYWRtaW5zZGstcGl5MDdAZmxnZC01NzQ5Mi5pYW0uZ3NlcnZpY2VhY2NvdW50LmNvbSIsInVpZCI6IlJyV1BuaTYxcXZReHRYQUVEZ2t2cWpRS0FidTEiLCJjbGFpbXMiOnsiZW1haWwiOiJjaGFuY2hvQGdtYWlsLmNvbSIsInByb3ZpZGVyIjoicGFzc3dvcmQiLCJ0eXBlIjoicHJhY3RpdGlvbmVyIiwidWlkIjoiUnJXUG5pNjFxdlF4dFhBRURna3ZxalFLQWJ1MSJ9fQ.q41BgjBfW1dwbRVDOyzTb_-eUBWWxS1DSqqIVLXU0iTPyQbJxuPccC7RxrSn2fW3Nric19pEpb52HXfDaasisemwHwpR3A9CCW2nMk_OUguyAwk2lF2t59t_LnChFjQoVzdn5-UczSBCJ2Qr9FV3nLBj53FZSQEuNcQQ7XgdmhnzlbmZck-3taQlffMA8zofcPZAY0CSO688EREAZRnhSsPIuUs2ct88Ud5dGUIJ9egtbpSbl7GLboWor2YiK_523Bp6GnL1bk_cpMG19jWs2IQVgT1zi9i4Unn5udUdw-CPHMHnOFU-NA9aARLNoJuThUmfopDcXHY2fzkXSKCoLA"

func TestSetUserInfoOK(t *testing.T) {
	cb := func(c echo.Context) error {
		ctx := c.Request().Context()

		err := ctx.Err()

		assert.Nil(t, err)

		return c.String(http.StatusOK, "OK")
	}
	c := SetupMiddlewareCase(http.MethodGet, "/", nil, map[string]string{
		"Authorization": token,
	})

	assert.NoError(t, c.Run(cb, SetUserInfo("user")))
}
