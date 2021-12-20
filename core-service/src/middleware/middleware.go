package middleware

import "github.com/jabardigitalservice/portal-jabar-services/core-service/src/config"

// GoMiddleware represent the data-struct for middleware
type GoMiddleware struct {
	// another stuff , may be needed by middleware
	JWTKey []byte
}

// InitMiddleware initialize the middleware
func InitMiddleware(cfg *config.Config) *GoMiddleware {
	return &GoMiddleware{
		JWTKey: []byte(cfg.JWT.AccessSecret),
	}
}
