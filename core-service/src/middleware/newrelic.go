package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/newrelic/go-agent/v3/newrelic"
)

func (m *GoMiddleware) NewRelic(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		txn := newrelic.FromContext(c.Request().Context())
		defer txn.End()
		return next(c)
	}
}
