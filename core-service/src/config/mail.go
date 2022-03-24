package config

import (
	"github.com/spf13/viper"
)

// SentryConfig is the sentry configuration
type MailConfig struct {
	SMTPHost     string
	SMTPPort     string
	SenderName   string
	ReceiverName string
	AuthEmail    string
	AuthPassword string
}

// LoadMailConfig loads the sentry configuration
func LoadMailConfig() MailConfig {
	return MailConfig{
		SMTPHost:     viper.GetString("SMTP_HOST"),
		SMTPPort:     viper.GetString("SMTP_PORT"),
		SenderName:   viper.GetString("SENDER_NAME"),
		ReceiverName: viper.GetString("RECEIVER_NAME"),
		AuthEmail:    viper.GetString("AUTH_EMAIL"),
		AuthPassword: viper.GetString("AUTH_PASSWORD"),
	}
}
