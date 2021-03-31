package email

import (
	"fmt"
	"gopkg.in/gomail.v2"
	"math/rand"
	"mime"
	"yangyj/pkg/config"
)

type Email struct {}

func (e *Email) send(mails ...*gomail.Message) (err error) {
	mailConfig := config.Config.Email
	index := rand.Intn(len(mailConfig))
	cfg := mailConfig[index]

	var m = make([]*gomail.Message, len(mails))
	for i, v := range mails {
		from := cfg.Username
		if cfg.Name != "" {
			from = fmt.Sprintf("%v<%v>", mime.QEncoding.Encode("utf-8", cfg.Name), cfg.Username)
		}
		v.SetHeader("From", from)
		m[i] = v
	}

	d := gomail.NewDialer(cfg.Host, cfg.Port, cfg.Username, cfg.Password)
	return d.DialAndSend(m...)
}

func New() *Email {
	return &Email{}
}

func (e *Email) CaptchaCode(address, code string) (err error) {
	m := gomail.NewMessage()
	m.SetHeader("To", address)
	subject := fmt.Sprintf("验证码 %v", code)
	m.SetHeader("Subject", subject)
	expired := config.Config.Captcha.Expired
	body := fmt.Sprintf("%v， %v分钟内有效，请勿泄露给他人！", subject, expired)
	m.SetBody("text/html", body)

	return e.send(m)
}