package middleware

import (
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/config"
	"github.com/newrelic/go-agent/v3/newrelic"
)

// GoMiddleware represent the data-struct for middleware
type GoMiddleware struct {
	// another stuff , may be needed by middleware
	JWTKey []byte
	Apm    *newrelic.Application
}

// InitMiddleware initialize the middleware
func InitMiddleware(cfg *config.Config, apm *newrelic.Application) *GoMiddleware {
	return &GoMiddleware{
		JWTKey: []byte(cfg.JWT.AccessSecret),
		Apm:    apm,
	}
}
