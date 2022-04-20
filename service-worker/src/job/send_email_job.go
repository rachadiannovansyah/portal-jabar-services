package job

import (
	"github.com/jabardigitalservice/portal-jabar-services/service-worker/src/config"
	"github.com/jabardigitalservice/portal-jabar-services/service-worker/src/utils"
	"gopkg.in/gomail.v2"
)

func SendEmailJob(to string, subject string, body string) (err error) {
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", config.LoadMailConfig().SenderName)
	mailer.SetHeader("To", to)
	mailer.SetHeader("Subject", subject)
	mailer.SetBody("text/html", body)

	dialer := utils.InitMail()

	err = dialer.DialAndSend(mailer)
	if err != nil {
		return
	}

	return
}
