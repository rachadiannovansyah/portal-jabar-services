package utils

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jabardigitalservice/portal-jabar-services/service-worker/src/config"
)

// Conn ...
type Conn struct {
	Mysql *sql.DB
}

// NewDBConn ...
func NewDBConn(cfg *config.Config) *Conn {
	return &Conn{
		Mysql: Initialize(cfg),
	}
}

// Initialize a new connection to the database
func Initialize(cfg *config.Config) *sql.DB {
	db, err := sql.Open("mysql", cfg.DB.DSN)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return db
}
