package config

import (
	"github.com/spf13/viper"
)

// SentryConfig is the sentry configuration
type SentryConfig struct {
	DSN              string
	TracesSampleRate float64
	Environment      string
}

// LoadSentryConfig loads the sentry configuration
func LoadSentryConfig() SentryConfig {
	return SentryConfig{
		DSN:              viper.GetString("SENTRY_DSN"),
		TracesSampleRate: viper.GetFloat64("SENTRY_TRACES_SAMPLE_RATE"),
		Environment:      viper.GetString("SENTRY_ENVIRONMENT"),
	}
}
