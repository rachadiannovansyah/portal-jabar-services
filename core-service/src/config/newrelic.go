package config

import (
	newrelic "github.com/newrelic/go-agent"
	"github.com/spf13/viper"
)

// LoadNewRelicConfig loads the newrelic configuration
func LoadNewRelicConfig() newrelic.Config {
	return newrelic.Config{
		AppName: viper.GetString("NEWRELIC_APP_NAME"),
		License: viper.GetString("NEWRELIC_LICENSE_KEY"),
		Enabled: viper.GetBool("NEWRELIC_ENABLED"),
	}
}
