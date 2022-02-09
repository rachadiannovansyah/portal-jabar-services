package main

import (
	"fmt"
	"log"
	"runtime"

	"github.com/jabardigitalservice/portal-jabar-services/service-worker/src/config"
	"github.com/jabardigitalservice/portal-jabar-services/service-worker/src/job"
	"github.com/jabardigitalservice/portal-jabar-services/service-worker/src/utils"
	"github.com/robfig/cron"
)

func main() {
	log.Println("Service RUN on DEBUG mode")
	cfg := config.NewConfig()
	conn := utils.NewDBConn(cfg)
	// defer func() {
	// 	err := conn.Mysql.Close()
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// }()

	c := cron.New()
	// run cron every minute
	c.AddFunc("@every 1m", func() { job.NewsArchiveJob(conn) })

	c.Start()

	fmt.Println("service-worker is running")
	runtime.Goexit()
}
