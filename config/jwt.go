package config

import "crypto/rsa"

type JWT struct {
	PublicKey   *rsa.PublicKey  `mapstructure:"public-key" json:"public-key"`                         // 公钥
	PrivateKey  *rsa.PrivateKey `mapstructure:"private-key" json:"private-key"`                       // 私钥
	ExpiresTime int64           `mapstructure:"expires-time" json:"expires-time" yaml:"expires-time"` // 过期时间
	BufferTime  int64           `mapstructure:"buffer-time" json:"buffer-time" yaml:"buffer-time"`    // 缓冲时间
	Issuer      string          `mapstructure:"issuer" json:"issuer" yaml:"issuer"`                   // 签发者
}
