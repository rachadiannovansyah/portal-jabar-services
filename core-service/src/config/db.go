package config

import (
	"fmt"
	"os"
)

// DBConfig represents DB configuration.
type DBConfig struct {
	DSN string
}

// LoadDBConfig loads DB configuration from file.
func LoadDBConfig() DBConfig {
	return DBConfig{
		DSN: fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_NAME"),
		),
	}
}
