package config

type FiberConfig struct {
	Prefork               bool   `mapstructure:"prefork" json:"prefork" yaml:"prefork"`
	ServerHeader          string `mapstructure:"server-header" json:"server-header" yaml:"server-header"`
	StrictRouting         bool   `mapstructure:"strict-routing" json:"strict-routing" yaml:"strict-routing"`
	CaseSensitive         bool   `mapstructure:"case-sensitive" json:"case-sensitive" yaml:"case-sensitive"`
	BodyLimit             int    `mapstructure:"body-limit" json:"body-limit" yaml:"body-limit"`
	AppName               string `mapstructure:"app-name" json:"app-name" yaml:"app-name"`
	Concurrency           int    `mapstructure:"concurrency" json:"concurrency" yaml:"concurrency"`
	DisableStartupMessage bool   `mapstructure:"DisableStartupMessage" json:"DisableStartupMessage" yaml:"DisableStartupMessage"`
}
