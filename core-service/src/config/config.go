package config

import (
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/labstack/echo/v4/middleware"
	"log"

	"github.com/spf13/viper"
)

// Config is the struct for the config
type Config struct {
	DB      DBConfig
	Sentry  SentryConfig
	Cors    middleware.CORSConfig
	ELastic elasticsearch.Config
}

// NewConfig creates a new Config struct
func NewConfig() *Config {
	viper.SetConfigFile(`.env`)
	viper.AutomaticEnv()
	viper.ReadInConfig()

	if viper.GetBool(`DEBUG`) {
		log.Println("Service RUN on DEBUG mode")
	}

	return &Config{
		DB:      LoadDBConfig(),
		Sentry:  LoadSentryConfig(),
		Cors:    LoadCorsConfig(),
		ELastic: LoadElasticConfig(),
	}
}
