package email

import (
	"gopkg.in/gomail.v2"
	"math/rand"
	"yangyj/backend/pkg/config"
)

type email struct {}

func (e *email) send(mails ...*gomail.Message) (err error) {
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
