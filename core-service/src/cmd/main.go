package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"time"

	sentryecho "github.com/getsentry/sentry-go/echo"

	"github.com/getsentry/sentry-go"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"

	httpDelivery "github.com/jabardigitalservice/portal-jabar-services/core-service/src/delivery/http"
	middl "github.com/jabardigitalservice/portal-jabar-services/core-service/src/delivery/http/middleware"
	repo "github.com/jabardigitalservice/portal-jabar-services/core-service/src/repositories/mysql"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/usecases"
)

func init() {
	viper.SetConfigFile(`.env`)
	viper.AutomaticEnv()
	viper.ReadInConfig()

	if viper.GetBool(`DEBUG`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func main() {
	dbHost := viper.GetString(`DB_HOST`)
	dbPort := viper.GetString(`DB_PORT`)
	dbUser := viper.GetString(`DB_USER`)
	dbPass := viper.GetString(`DB_PASSWORD`)
	dbName := viper.GetString(`DB_NAME`)
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	log.Println(connection)
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Jakarta")
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())
	dbConn, err := sql.Open(`mysql`, dsn)

	if err != nil {
		log.Fatal("A", err)
	}
	err = dbConn.Ping()
	if err != nil {
		log.Fatal("B", err)
	}

	defer func() {
		err := dbConn.Close()
		if err != nil {
			log.Fatal("C", err)
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
		Dsn:              viper.GetString(`SENTRY_DSN`),
		TracesSampleRate: 1.0,
	}); err != nil {
		fmt.Printf("Sentry initialization failed: %v\n", err)
	}

	e.Use(sentryecho.New(sentryecho.Options{
		Repanic: true,
	}))

	timeoutContext := time.Duration(viper.GetInt("APP_TIMEOUT")) * time.Second

	// init repo category repo
	nr := repo.NewMysqlNewsRepository(dbConn)
	ctg := repo.NewMysqlCategoriesRepository(dbConn)
	ir := repo.NewMysqlInformationsRepository(dbConn)

	// news handler
	nu := usecases.NewNewsUsecase(nr, ctg, timeoutContext)
	httpDelivery.NewContentHandler(v1, r, nu)

	// informations handler
	iu := usecases.NewInformationUcase(ir, ctg, timeoutContext)
	httpDelivery.NewInformationHandler(e, r, iu)

	log.Fatal(e.Start(viper.GetString("APP_ADDRESS")))
}
