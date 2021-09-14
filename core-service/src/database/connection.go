package database

import (
	"database/sql"
	"log"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/config"
)

// InitDB a new connection to the database
func InitDB(cfg *config.Config) *sql.DB {
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
