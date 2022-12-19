package config

import (
	"time"

	"github.com/spf13/viper"
)

type RedisConfig struct {
	Host string
	Port int
	TTL  time.Duration
}

func LoadRedisConfig() RedisConfig {
	return RedisConfig{
		Host: viper.GetString("REDIS_HOST"),
		Port: viper.GetInt("REDIS_PORT"),
		TTL:  time.Duration(viper.GetInt("REDIS_TTL")) * time.Minute,
	}
}
