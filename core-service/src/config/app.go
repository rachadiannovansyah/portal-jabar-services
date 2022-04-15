package config

import (
	"github.com/spf13/viper"
)

// AppConfig represents app configuration.
type AppConfig struct {
	Name    string
	Version string
	CmsUrl  string
}

// LoadAppConfig loads app configuration from file.
func LoadAppConfig() AppConfig {
	return AppConfig{
		Name:    viper.GetString("APP_NAME"),
		Version: viper.GetString("APP_VERSION"),
		CmsUrl:  viper.GetString("PORTAL_JABAR_CMS_URL"),
	}
}
