package config

import (
	"time"

	"github.com/spf13/viper"
)

// JWTConfig represents JWT configuration.
type JWTConfig struct {
	AccessSecret  string
	RefreshSecret string
	TTL           time.Duration
	RefreshTTL    time.Duration
}

// LoadJWTConfig loads JWT configuration.
func LoadJWTConfig() JWTConfig {
	return JWTConfig{
		AccessSecret:  viper.GetString("JWT_ACCESS_SECRET"),
		RefreshSecret: viper.GetString("JWT_REFRESH_SECRET"),
		TTL:           viper.GetDuration("JWT_TTL") * time.Second,
		RefreshTTL:    viper.GetDuration("JWT_REFRESH_TTL") * time.Second,
	}
}
