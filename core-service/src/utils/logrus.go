package utils

import (
	"os"

	_ "github.com/newrelic/go-agent/v3/integrations/nrmysql"
	"github.com/sirupsen/logrus"
)

type Logrus struct {
	logger *logrus.Logger
}

func NewLogrus() *Logrus {
	// set formatter logs
	log := logrus.New()
	formatter := &logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			"msg": "message",
		},
	}
	log.Formatter = formatter
	log.SetOutput(os.Stdout)

	return &Logrus{
		logger: log,
	}
}

func (l *Logrus) Error(logsField logrus.Fields, message interface{}) {
	l.logger.WithFields(logsField).Error(message)
}

func (l *Logrus) Info(logsField logrus.Fields, message interface{}) {
	l.logger.WithFields(logsField).Info(message)
}
