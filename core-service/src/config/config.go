package config

import (
	"log"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/labstack/echo/v4/middleware"
	newrelic "github.com/newrelic/go-agent"

	"github.com/spf13/viper"
)

// Config is the struct for the config
type Config struct {
	DB       DBConfig
	JWT      JWTConfig
	Sentry   SentryConfig
	Cors     middleware.CORSConfig
	ELastic  elasticsearch.Config
	AWS      AwsConfig
	Redis    RedisConfig
	NewRelic newrelic.Config
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
		DB:       LoadDBConfig(),
		JWT:      LoadJWTConfig(),
		Sentry:   LoadSentryConfig(),
		Cors:     LoadCorsConfig(),
		ELastic:  LoadElasticConfig(),
		AWS:      LoadAwsConfig(),
		Redis:    LoadRedisConfig(),
		NewRelic: LoadNewRelicConfig(),
	}
}
