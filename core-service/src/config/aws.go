package config

import (
	"github.com/spf13/viper"
)

// SentryConfig is the sentry configuration
type AwsConfig struct {
	AccessKey       string
	SecretAccessKey string
	Region          string
	Bucket          string
}

// LoadAwsConfig loads the sentry configuration
func LoadAwsConfig() AwsConfig {
	return AwsConfig{
		AccessKey:       viper.GetString("AWS_ACCESS_KEY_ID"),
		SecretAccessKey: viper.GetString("AWS_SECRET_ACCESS_KEY"),
		Region:          viper.GetString("AWS_DEFAULT_REGION"),
		Bucket:          viper.GetString("AWS_BUCKET"),
	}
}
