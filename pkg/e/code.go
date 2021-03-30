package e

const (
	SUCCESS         = 0
	ERROR           = 10001
	PARAMS_INVALID  = 10002
	CAPTCHA_INVALID = 10003
)

var codeMap = map[int]string{
	SUCCESS:         "app.success",
	ERROR:           "app.error",
	PARAMS_INVALID:  "app.params.invalid",
	CAPTCHA_INVALID: "captcha.invalid",
}