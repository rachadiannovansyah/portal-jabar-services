package config

import "github.com/spf13/viper"

type RedisConfig struct {
	Host string
	Port int
}

func LoadRedisConfig() RedisConfig {
	return RedisConfig{
		Host: viper.GetString("REDIS_HOST"),
		Port: viper.GetInt("REDIS_PORT"),
	}
}
