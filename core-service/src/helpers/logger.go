package helpers

import (
	"os"

	"github.com/sirupsen/logrus"
)

type Logger struct {
	logger *logrus.Logger
}

func InitLogger() *Logger {
	// set formatter logs
	logger := logrus.New()
	formatter := &logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			"msg": "message",
		},
	}
	logger.Formatter = formatter
	logger.SetOutput(os.Stdout)

	return &Logger{
		logger: logger,
	}
}

func (l *Logger) Error(logsField logrus.Fields, message interface{}) {
	l.logger.WithFields(logsField).Error(message)
}

func (l *Logger) Info(logsField logrus.Fields, message interface{}) {
	l.logger.WithFields(logsField).Info(message)
}
