package database

import (
	"database/sql"
	"github.com/elastic/go-elasticsearch/v8"
	"log"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/config"
)

type DBConn struct {
	Mysql   *sql.DB
	Elastic *elasticsearch.Client
}

func NewDBConn(cfg *config.Config) *DBConn {
	return &DBConn{
		Mysql:   initMysql(cfg),
		Elastic: initELastClient(cfg),
	}
}

// InitDB a new connection to the database
func initMysql(cfg *config.Config) *sql.DB {
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

func initELastClient(cfg *config.Config) *elasticsearch.Client {
	es, err := elasticsearch.NewClient(cfg.ELastic)
	if err != nil {
		log.Fatal(err)
	}

	// 1. Get cluster info
	res, err := es.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	return es
}
