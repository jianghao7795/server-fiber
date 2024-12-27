package config

// 缓存时间
type Cache struct {
	Time int64 `mapstructure:"time" json:"time" yaml:"time"` // 时间
}
