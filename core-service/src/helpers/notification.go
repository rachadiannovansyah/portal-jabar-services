package helpers

import (
	"log"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/utils"
	"github.com/spf13/viper"
	"gopkg.in/gomail.v2"
)

func SendMail(res domain.User, template domain.Mail) (err error) {
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", viper.GetString("SENDER_NAME"))
	mailer.SetHeader("To", viper.GetString("RECEIVER_NAME"))
	mailer.SetHeader("Subject", template.Subject)
	mailer.SetBody("text/html", template.Body)

	dialer := utils.InitMail()

	err = dialer.DialAndSend(mailer)
	if err != nil {
		return
	}

	log.Println("Mail sent!")

	return
}
