package config

import (
	"log"

	"github.com/labstack/echo/v4/middleware"
	newrelic "github.com/newrelic/go-agent"

	"github.com/spf13/viper"
)

// Config is the struct for the config
type Config struct {
	App      AppConfig
	DB       DBConfig
	JWT      JWTConfig
	Sentry   SentryConfig
	Cors     middleware.CORSConfig
	ELastic  ElasticConfig
	AWS      AwsConfig
	Redis    RedisConfig
	NewRelic newrelic.Config
	Mail     MailConfig
	External ExternalConfig
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
		App:      LoadAppConfig(),
		DB:       LoadDBConfig(),
		JWT:      LoadJWTConfig(),
		Sentry:   LoadSentryConfig(),
		Cors:     LoadCorsConfig(),
		ELastic:  LoadElasticConfig(),
		AWS:      LoadAwsConfig(),
		Redis:    LoadRedisConfig(),
		NewRelic: LoadNewRelicConfig(),
		Mail:     LoadMailConfig(),
		External: LoadExternalConfig(),
	}
}
