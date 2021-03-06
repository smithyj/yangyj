package email

import (
	"fmt"
	"gopkg.in/gomail.v2"
	"math/rand"
	"yangyj/backend/pkg/config"
)

type Email struct {}

func (e *Email) send(mails ...*gomail.Message) (err error) {
	mailConfig := config.Config.Email
	index := rand.Intn(len(mailConfig))
	cfg := mailConfig[index]

	var m = make([]*gomail.Message, len(mails))
	for i, v := range mails {
		v.SetHeader("From", cfg.Username)
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
	m.SetHeader("Subject", "【YANGYJ】邮箱验证码")
	expired := config.Config.Captcha.Expired
	body := fmt.Sprintf("验证码: <span style='color:red'>%v</span>，%v分钟内有效，请勿泄露给他人！", code, expired)
	m.SetBody("text/html", body)

	return e.send(m)
}