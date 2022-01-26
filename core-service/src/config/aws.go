package config

import (
	"github.com/spf13/viper"
)

// SentryConfig is the sentry configuration
type AwsConfig struct {
	Env             string
	AccessKey       string
	SecretAccessKey string
	Region          string
	Bucket          string
	Cloudfront      string
}

// LoadAwsConfig loads the sentry configuration
func LoadAwsConfig() AwsConfig {
	return AwsConfig{
		Env:             viper.GetString("AWS_ENV"),
		AccessKey:       viper.GetString("AWS_ACCESS_KEY_ID"),
		SecretAccessKey: viper.GetString("AWS_SECRET_ACCESS_KEY"),
		Region:          viper.GetString("AWS_DEFAULT_REGION"),
		Bucket:          viper.GetString("AWS_BUCKET"),
		Cloudfront:      viper.GetString("AWS_CLOUDFRONT"),
	}
}
