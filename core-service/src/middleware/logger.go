package middleware

import (
	"time"

	"github.com/sirupsen/logrus"

	"github.com/labstack/echo/v4"
)

func (m *GoMiddleware) Logging(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		m.Logger.Info(logrus.Fields(logrus.Fields{
			"timestamps": time.Now().Format("2006-01-02 15:04:05"),
			"method":     c.Request().Method,
			"uri":        c.Request().URL.String(),
			"ip":         c.Request().RemoteAddr,
		}), "incoming an request")
		return next(c)
	}
}
