package main

import (
	"fmt"
	"runtime"

	"github.com/jabardigitalservice/portal-jabar-services/service-worker/src/config"
	"github.com/jabardigitalservice/portal-jabar-services/service-worker/src/job"
	"github.com/jabardigitalservice/portal-jabar-services/service-worker/src/utils"
	"github.com/robfig/cron"
)

func main() {
	cfg := config.NewConfig()
	conn := utils.NewDBConn(cfg)

	c := cron.New()
	// run cron @daily
	c.AddFunc("@daily", func() { job.NewsArchiveJob(conn) })
	c.AddFunc("@daily", func() { job.NewsPublishingJob(conn) })

	c.Start()

	fmt.Println("service-worker is running")
	runtime.Goexit()
}
