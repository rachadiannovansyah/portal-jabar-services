package config

import (
	"time"

	"github.com/spf13/viper"
)

// AppConfig represents app configuration.
type AppConfig struct {
	Name           string
	Version        string
	CmsUrl         string
	PortalUrl      string
	Env            string
	Domain         string
	ContextTimeout time.Duration
}

// LoadAppConfig loads app configuration from file.
func LoadAppConfig() AppConfig {
	return AppConfig{
		Name:           viper.GetString("APP_NAME"),
		Version:        viper.GetString("APP_VERSION"),
		CmsUrl:         viper.GetString("PORTAL_JABAR_CMS_URL"),
		PortalUrl:      viper.GetString("PORTAL_JABAR_LANDING_PAGE_URL"),
		Env:            viper.GetString("APP_ENV"),
		Domain:         viper.GetString("APP_DOMAIN"),
		ContextTimeout: viper.GetDuration("APP_TIMEOUT") * time.Second,
	}
}
