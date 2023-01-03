package middleware

import (
	"github.com/sirupsen/logrus"

	"github.com/labstack/echo/v4"
)

func (m *GoMiddleware) Logging(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		m.Logger.Info(logrus.Fields(logrus.Fields{
			"user_agent": c.Request().Header.Get("User-Agent"),
			"method":     c.Request().Method,
			"uri":        c.Request().URL.String(),
			"ip":         c.Request().RemoteAddr,
			"host":       c.Request().URL.Host,
		}), "incoming an request")
		return next(c)
	}
}
