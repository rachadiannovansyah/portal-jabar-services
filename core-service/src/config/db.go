package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// DBConfig represents DB configuration.
type DBConfig struct {
	DSN             string
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxIdleTime int
	ConnMaxLifetime int
}

// LoadDBConfig loads DB configuration from file.
func LoadDBConfig() DBConfig {
	return DBConfig{
		DSN: fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
			viper.GetString("DB_USER"),
			viper.GetString("DB_PASSWORD"),
			viper.GetString("DB_HOST"),
			viper.GetString("DB_PORT"),
			viper.GetString("DB_NAME"),
		),
		MaxIdleConns:    viper.GetInt("DB_MAX_IDLE_CONNS"),
		MaxOpenConns:    viper.GetInt("DB_MAX_OPEN_CONNS"),
		ConnMaxIdleTime: viper.GetInt("DB_CONN_MAX_IDLE_TIME"),
		ConnMaxLifetime: viper.GetInt("DB_CONN_MAX_LIFETIME"),
	}
}
