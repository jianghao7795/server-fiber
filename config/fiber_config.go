package config

import "github.com/gofiber/fiber/v2"

type FiberConfig struct {
	Prefork               bool   `mapstructure:"prefork" json:"prefork" yaml:"prefork"`
	ServerHeader          string `mapstructure:"server_header" json:"server_header" yaml:"server_header"`
	StrictRouting         bool   `mapstructure:"strict_routing" json:"strict_routing" yaml:"strict_routing"`
	CaseSensitive         bool   `mapstructure:"case_sensitive" json:"case_sensitive" yaml:"case_sensitive"`
	BodyLimit             int    `mapstructure:"body_limit" json:"body_limit" yaml:"body_limit"`
	AppName               string `mapstructure:"app_name" json:"app_name" yaml:"app_name"`
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
