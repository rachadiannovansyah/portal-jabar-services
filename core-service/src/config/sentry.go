package config

import "os"

// SentryConfig is the sentry configuration
type SentryConfig struct {
	DSN              string
	TracesSampleRate float64
}

// LoadSentryConfig loads the sentry configuration
func LoadSentryConfig() SentryConfig {
	return SentryConfig{
		DSN:              os.Getenv("SENTRY_DSN"),
		TracesSampleRate: 1.0,
	}
}
