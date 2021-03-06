package sms

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"net/url"
	"strings"
	"yangyj/backend/pkg/config"
)

type aliyun struct {
	countryNo string
}

func (s *aliyun) send(request *dysmsapi.SendSmsRequest) (err error) {
	cfg := config.Config.Sms.Aliyun
	client, err := dysmsapi.NewClientWithAccessKey("cn-hangzhou", cfg.Appid, cfg.Appsecret)
	if err != nil {
		return
	}

	response, err := client.SendSms(request)
	if err != nil {
		return
	}
	if strings.ToUpper(response.Code) != "OK" {
		return errors.New(response.Message)
	}
	return
}

func (s *aliyun) buildRequest(signName, phone, tplCode, tplParam string) *dysmsapi.SendSmsRequest {
	if s.countryNo != "86" {
		phone = url.QueryEscape(fmt.Sprintf("%v%v", s.countryNo, phone))
	}

	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"
	request.PhoneNumbers = phone
	request.SignName = signName
	request.TemplateCode = tplCode
	request.TemplateParam = tplParam

	return request
}

func (s *aliyun) CaptchaCode(phone, code string) (err error) {
	params := struct {
		Code string `json:"code"`
	}{
		Code: code,
	}
	byteSlice, err := json.Marshal(params)
	if err != nil {
		return
	}
	tpl := ALIYUN_CAPTCHACODE_TPL
	if s.countryNo != "86" {
		tpl = ALIYUN_CAPTCHACODE_EN_TPL

	}
	tpls := strings.Split(tpl, ",")
	singName := tpls[0]
	tplCode := tpls[1]
	request := s.buildRequest(singName, phone, tplCode, string(byteSlice))
	return s.send(request)
}
