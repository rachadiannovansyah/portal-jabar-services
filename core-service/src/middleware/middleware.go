package middleware

import (
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/config"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/helpers"
	"github.com/newrelic/go-agent/v3/newrelic"
)

// GoMiddleware represent the data-struct for middleware
type GoMiddleware struct {
	// another stuff , may be needed by middleware
	JWTKey []byte
	Apm    *newrelic.Application
	Logger helpers.Logger
}

// InitMiddleware initialize the middleware
func InitMiddleware(cfg *config.Config, apm *newrelic.Application, logger helpers.Logger) *GoMiddleware {
	return &GoMiddleware{
		JWTKey: []byte(cfg.JWT.AccessSecret),
		Apm:    apm,
		Logger: logger,
	}
}
