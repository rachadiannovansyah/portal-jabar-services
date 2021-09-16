package main

import (
	"fmt"
	"log"
	"time"

	"github.com/getsentry/sentry-go"
	sentryecho "github.com/getsentry/sentry-go/echo"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/config"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/database"
	httpDelivery "github.com/jabardigitalservice/portal-jabar-services/core-service/src/delivery/http"
	middl "github.com/jabardigitalservice/portal-jabar-services/core-service/src/delivery/http/middleware"
	repo "github.com/jabardigitalservice/portal-jabar-services/core-service/src/repositories/mysql"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/usecases"
)

func main() {
	cfg := config.NewConfig()
	dbConn := database.InitDB(cfg)
	defer func() {
		err := dbConn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	e := echo.New()
	middL := middl.InitMiddleware()
	e.Use(middL.CORS)
	e.Use(middL.SENTRY)
	e.Use(middleware.Logger())

	// api v1
	v1 := e.Group("/v1")

	// restricted group
	r := v1.Group("")
	r.Use(middL.JWT)

	if err := sentry.Init(sentry.ClientOptions{
		Dsn:              cfg.Sentry.DSN,
		TracesSampleRate: cfg.Sentry.TracesSampleRate,
	}); err != nil {
		fmt.Printf("Sentry initialization failed: %v\n", err)
	}

	e.Use(sentryecho.New(sentryecho.Options{
		Repanic: true,
	}))

	timeoutContext := time.Duration(viper.GetInt("APP_TIMEOUT")) * time.Second

	// init repo category repo
	mysqlRepos := repo.NewMysqlRepositories(dbConn)
	usecases := usecases.NewUcase(mysqlRepos, timeoutContext)
	httpDelivery.NewHandler(v1, r, usecases)

	log.Fatal(e.Start(viper.GetString("APP_ADDRESS")))
}
