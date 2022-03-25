package helpers

import (
	"log"
	"strings"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/utils"
	"github.com/spf13/viper"
	"gopkg.in/gomail.v2"
)

func SendMail(user domain.User, template domain.Template) (err error) {
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", viper.GetString("SENDER_NAME"))
	mailer.SetHeader("To", user.Email)
	mailer.SetHeader("Subject", template.Subject)
	messg := template.Body
	messg = strings.ReplaceAll(messg, "{name}", user.Name)
	messg = strings.ReplaceAll(messg, "{unitName}", user.UnitName)
	mailer.SetBody("text/html", messg)

	dialer := utils.InitMail()

	err = dialer.DialAndSend(mailer)
	if err != nil {
		return
	}

	log.Println("Mail sent!")

	return
}
