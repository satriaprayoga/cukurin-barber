package email

import (
	"net/mail"

	"github.com/satriaprayoga/cukurin-barber/pkg/settings"
	"gopkg.in/gomail.v2"
)

func SendEmail(to string, subject string, htmlBody string) error {
	smtp := settings.AppConfigSetting.SMTP

	from := mail.Address{
		Name:    smtp.Identity,
		Address: smtp.Sender,
	}

	m := gomail.NewMessage()
	m.Reset()
	m.SetHeader("From", from.String())
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)

	m.SetBody("text/html", htmlBody)

	d := gomail.NewDialer(smtp.Server, smtp.Port, smtp.User, smtp.Passwd)
	return d.DialAndSend(m)
}
