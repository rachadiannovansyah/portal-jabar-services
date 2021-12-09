package config

import (
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/spf13/viper"
)

// LoadElasticConfig loads the elasticsearch configuration
func LoadElasticConfig() elasticsearch.Config {
	return elasticsearch.Config{
		CloudID: viper.GetString("ELASTIC_CLOUD_ID"),
		APIKey:  viper.GetString("ELASTIC_API_KEY"),
	}
}
