package config

import (
	"log"

	"github.com/spf13/viper"
)

// Config is the struct for the config
type Config struct {
	DB      DBConfig
	Redis   RedisConfig
	ELastic ElasticConfig
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
		Redis:   LoadRedisConfig(),
		ELastic: LoadElasticConfig(),
	}
}
