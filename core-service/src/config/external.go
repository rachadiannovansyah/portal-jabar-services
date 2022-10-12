package config

import (
	"github.com/spf13/viper"
)

// ExternalConfig represents app configuration.
type ExternalConfig struct {
	CoreDataUrl string
}

// LoadExternalConfig loads app configuration from file.
func LoadExternalConfig() ExternalConfig {
	return ExternalConfig{
		CoreDataUrl: viper.GetString("EXTERNAL_CORE_DATA_URL"),
	}
}
