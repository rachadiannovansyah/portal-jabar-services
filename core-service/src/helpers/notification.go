package helpers

import (
	"fmt"
	"strings"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/utils"
	"github.com/spf13/viper"
	"gopkg.in/gomail.v2"
)

func ReplaceBodyParams(body string, params []string) string {

	for i := 0; i < len(params); i++ {
		param, value := fmt.Sprintf("{param%v}", i+1), params[i]
		body = strings.ReplaceAll(body, param, value)
	}

	return body
}

func SendEmail(to string, template domain.Template, params []string) (err error) {
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", viper.GetString("SENDER_NAME"))
	mailer.SetHeader("To", to)
	mailer.SetHeader("Subject", template.Subject)
	messg := ReplaceBodyParams(template.Body, params)
	mailer.SetBody("text/html", messg)

	dialer := utils.InitMail()

	err = dialer.DialAndSend(mailer)
	if err != nil {
		return
	}

	return
}
