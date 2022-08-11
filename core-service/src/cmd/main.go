package main

import (
	"log"
	"time"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/utils"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/cmd/server"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/config"
)

func main() {
	cfg := config.NewConfig()
	apm := utils.NewApm(cfg)
	conn := utils.NewDBConn(cfg)
	defer func() {
		err := conn.Mysql.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	timeoutContext := time.Duration(viper.GetInt("APP_TIMEOUT")) * time.Second

	// init repo category repo
	mysqlRepos := server.NewRepository(conn)
	usecases := server.NewUcase(cfg, conn, mysqlRepos, timeoutContext)
	server.NewHandler(cfg, apm, usecases)
}
