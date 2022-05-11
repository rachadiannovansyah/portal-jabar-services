package utils

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jabardigitalservice/portal-jabar-services/service-worker/src/config"
	"gopkg.in/gomail.v2"
)

// Conn ...
type Conn struct {
	Mysql   *sql.DB
	Redis   *redis.Client
	Elastic *elasticsearch.Client
}

// NewDBConn ...
func NewDBConn(cfg *config.Config) *Conn {
	return &Conn{
		Mysql:   Initialize(cfg),
		Redis:   initRedisClient(cfg),
		Elastic: initELastClient(cfg),
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

// initRedisClient ...
func initRedisClient(cfg *config.Config) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", cfg.Redis.Host, cfg.Redis.Port),
	})

	_, err := rdb.Ping(rdb.Context()).Result()
	if err != nil {
		log.Fatal(err)
	}

	return rdb
}

func InitMail() *gomail.Dialer {
	port, _ := strconv.Atoi(config.LoadMailConfig().SMTPPort)

	dialer := gomail.NewDialer(
		config.LoadMailConfig().SMTPHost,
		port,
		config.LoadMailConfig().AuthEmail,
		config.LoadMailConfig().AuthPassword,
	)

	return dialer
}

func initELastClient(cfg *config.Config) *elasticsearch.Client {
	es, err := elasticsearch.NewClient(*cfg.ELastic.ElasticConfig)
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
