package config

import (
	"github.com/spf13/viper"
	"time"
)

// JWTConfig represents JWT configuration.
type JWTConfig struct {
	AccessSecret       string
	RefreshSecret      string
	ExpireCount        time.Duration
	ExpireRefreshCount time.Duration
}

// LoadJWTConfig loads JWT configuration.
func LoadJWTConfig() JWTConfig {
	return JWTConfig{
		AccessSecret:       viper.GetString("JWT_ACCESS_SECRET"),
		RefreshSecret:      viper.GetString("JWT_REFRESH_SECRET"),
		ExpireCount:        2,
		ExpireRefreshCount: 168,
	}
}
