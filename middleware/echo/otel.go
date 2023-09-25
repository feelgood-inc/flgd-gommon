package middleware

import "github.com/labstack/echo/v4"

func PropagateTraceFromRequest() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			// Check if request has a trace
			// If it does, propagate it
			panic("not implemented")
		}
	}
}
