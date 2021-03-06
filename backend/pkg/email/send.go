package email

import (
	"fmt"
	"gopkg.in/gomail.v2"
	"yangyj/backend/pkg/config"
)

func SendCaptchaCode(address, code string) (err error) {
	e := &email{}

	m := gomail.NewMessage()
	m.SetHeader("To", address)
	m.SetHeader("Subject", "【YANGYJ】邮箱验证码")
	expired := config.Config.Captcha.Expired
	body := fmt.Sprintf("验证码: <span style='color:red'>%v</span>，%v分钟内有效，请勿泄露给他人！", code, expired)
	m.SetBody("text/html", body)

	return e.send(m)
}
