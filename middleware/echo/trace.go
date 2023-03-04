package middleware

import (
	"github.com/labstack/echo/v4"
	"go.opentelemetry.io/otel/trace"
)

const headerTraceID = "X-Trace-Id"

func Trace(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		rootSpan := trace.SpanFromContext(c.Request().Context())
		defer rootSpan.End()

		c.Response().Header().Set(headerTraceID, rootSpan.SpanContext().TraceID().String())

		return next(c)
	}
}
