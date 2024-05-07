package response

import "github.com/wenlng/go-captcha/captcha"

type SysCaptchaResponse struct {
	CaptchaId     string `json:"captchaId"`
	PicPath       string `json:"picPath"`
	CaptchaLength int    `json:"captchaLength"`
}

type SysCaptchaImgResponse struct {
	Dots       map[int]captcha.CharDot `json:"dots"`
	Images     string                  `json:"images"`
	ThemeImage string                  `json:"themeImage"`
	CaptchaKey string                  `json:"captcha_key"`
}
