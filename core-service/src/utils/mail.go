package utils

import (
	"strconv"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/config"
	"gopkg.in/gomail.v2"
)

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
