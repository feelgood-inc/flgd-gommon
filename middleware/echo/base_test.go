package middleware

import (
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
	"net/http/httptest"
)

type MiddlewareCase struct {
	Req     *http.Request
	Res     *httptest.ResponseRecorder
	context echo.Context
}

func SetupMiddlewareCase(method string, url string, body io.Reader, headers map[string]string) MiddlewareCase {
	e := echo.New()
	req := httptest.NewRequest(method, url, body)

	if len(headers) > 0 {
		for key, value := range headers {
			req.Header.Set(key, value)
		}
	}

	res := httptest.NewRecorder()
	c := e.NewContext(req, res)

	return MiddlewareCase{req, res, c}
}

func (c MiddlewareCase) Run(cb echo.HandlerFunc, m echo.MiddlewareFunc) error {
	h := m(cb)

	return h(c.context)
}
