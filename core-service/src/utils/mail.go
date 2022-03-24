package utils

import (
	"strconv"

	"github.com/spf13/viper"
	"gopkg.in/gomail.v2"
)

func InitMail() *gomail.Dialer {
	port, _ := strconv.Atoi(viper.GetString("SMTP_PORT"))

	dialer := gomail.NewDialer(
		viper.GetString("SMTP_HOST"),
		port,
		viper.GetString("AUTH_EMAIL"),
		viper.GetString("AUTH_PASSWORD"),
	)

	return dialer
}
