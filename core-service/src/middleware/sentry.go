package middleware

import (
	"github.com/getsentry/sentry-go"
	"github.com/labstack/echo/v4"
)

// SENTRY will handle the SENTRY middleware
func (m *GoMiddleware) SENTRY(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		span := sentry.StartSpan(
			c.Request().Context(), "", sentry.TransactionName(c.Request().URL.Path),
		)
		span.Finish()
		return next(c)
	}
}
