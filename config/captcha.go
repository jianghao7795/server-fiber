package config

// Captcha 验证码配置
type Captcha struct {
	KeyLong   int     `mapstructure:"key-long" json:"key-long" yaml:"key-long"`       // 验证码长度
	ImgWidth  int     `mapstructure:"img-width" json:"img-width" yaml:"img-width"`    // 验证码宽度
	ImgHeight int     `mapstructure:"img-height" json:"img-height" yaml:"img-height"` // 验证码高度
	DotCount  int     `mapstructure:"dot-count" json:"dot-count" yaml:"dot-count"`    // 干扰点数量
	MaxSkew   float64 `mapstructure:"max-skew" json:"max-skew" yaml:"max-skew"`       // 最大倾斜度
}

// 图片验证码配置
type ImageCaptcha struct {
}
