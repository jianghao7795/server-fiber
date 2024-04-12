package config

import "github.com/gofiber/fiber/v2"

type FiberConfig struct {
	Prefork               bool   `mapstructure:"prefork" json:"prefork" yaml:"prefork"`
	ServerHeader          string `mapstructure:"server-header" json:"server-header" yaml:"server-header"`
	StrictRouting         bool   `mapstructure:"strict-routing" json:"strict-routing" yaml:"strict-routing"`
	CaseSensitive         bool   `mapstructure:"case-sensitive" json:"case-sensitive" yaml:"case-sensitive"`
	BodyLimit             int    `mapstructure:"body-limit" json:"body-limit" yaml:"body-limit"`
	AppName               string `mapstructure:"app-name" json:"app-name" yaml:"app-name"`
	Concurrency           int    `mapstructure:"concurrency" json:"concurrency" yaml:"concurrency"`
	DisableStartupMessage bool   `mapstructure:"DisableStartupMessage" json:"DisableStartupMessage" yaml:"DisableStartupMessage"`
	JSONEncoder           func(val interface{}) ([]byte, error)
	JSONDecoder           func(buf []byte, val interface{}) error
	ErrorHandler          func(ctx *fiber.Ctx, err error) error
}

type FiberLogger struct {
	Done       func(c *fiber.Ctx, logString []byte)
	Format     string `mapstructure:"format" json:"format" yaml:"format"`
	TimeFormat string `mapstructure:"time_format" json:"time_format" yaml:"time_format"`
	TimeZone   string `mapstructure:"time_zone" json:"time_zone" yaml:"time_zone"`
	IsOpen     bool   `mapstructure:"is_open" json:"is_open" yaml:"is_open"`
}
