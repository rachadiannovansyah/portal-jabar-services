package main

import (
	"log"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/utils"

	_ "github.com/go-sql-driver/mysql"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/cmd/server"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/config"
)

func main() {
	cfg := config.NewConfig()
	apm := utils.NewApm(cfg)
	conn := utils.NewDBConn(cfg)
	logrus := utils.NewLogrus()

	defer func() {
		err := conn.Mysql.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	// init repo category repo
	mysqlRepos := server.NewRepository(conn, cfg, logrus)
	usecases := server.NewUcase(cfg, conn, mysqlRepos, cfg.App.ContextTimeout)
	server.NewHandler(cfg, apm, usecases, *logrus)
}
