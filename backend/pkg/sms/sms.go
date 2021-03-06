package sms

import (
	"errors"
	"yangyj/backend/pkg/config"
	"yangyj/backend/pkg/helper"
)

type Sms interface {
	CaptchaCode(phone, code string) (err error)
}

func New(countryNo string) (s Sms, err error) {
	countryNo = helper.FilterCountryNo(countryNo)

	kind := config.Config.Sms.Kind
	switch kind {
	case "smsbao":
		s = &smsbao{
			countryNo: countryNo,
		}
	case "aliyun":
		s = &aliyun{
			countryNo: countryNo,
		}
	default:
		err = errors.New("Unknown sms kind")
	}
	return
}
