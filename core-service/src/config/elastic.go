package config

import (
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/spf13/viper"
)

type ElasticConfig struct {
	ElasticConfig *elasticsearch.Config
	IndexContent  string
	IndexSize     int
}

// LoadElasticConfig loads the elasticsearch configuration
func LoadElasticConfig() ElasticConfig {
	return ElasticConfig{
		ElasticConfig: &elasticsearch.Config{
			CloudID: viper.GetString("ELASTIC_CLOUD_ID"),
			APIKey:  viper.GetString("ELASTIC_API_KEY"),
		},
		IndexContent: viper.GetString("ELASTIC_CONTENT_INDEX"),
		IndexSize:    viper.GetInt("ELASTIC_INDEX_SIZE"),
	}
}
