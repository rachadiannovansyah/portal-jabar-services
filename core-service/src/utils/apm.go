package utils

import (
	"log"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/config"
	"github.com/newrelic/go-agent/v3/newrelic"
)

type Apm struct {
	NewRelic *newrelic.Application
}

func NewApm(cfg *config.Config) *Apm {
	return &Apm{
		NewRelic: initNewRelic(cfg),
	}
}

func initNewRelic(cfg *config.Config) *newrelic.Application {
	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName(cfg.NewRelic.AppName),
		newrelic.ConfigLicense(cfg.NewRelic.License),
		newrelic.ConfigDistributedTracerEnabled(cfg.NewRelic.Enabled),
	)

	if err != nil {
		log.Fatal(err)
	}

	return app
}
