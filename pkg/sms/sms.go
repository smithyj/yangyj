package sms

import (
	"errors"
	"yangyj/pkg/config"
	"yangyj/pkg/helper"
)

type Sms interface {
	CaptchaCode(phone, code string) (err error)
}

func New(countryCode string) (s Sms, err error) {
	countryCode = helper.FilterCountryCode(countryCode)

	kind := config.Config.Sms.Kind
	switch kind {
	case "smsbao":
		s = &smsbao{
			countryCode: countryCode,
		}
	case "aliyun":
		s = &aliyun{
			countryCode: countryCode,
		}
	default:
		err = errors.New("Unknown sms kind")
	}
	return
}
