package config

type Local struct {
	Path string `mapstructure:"path" json:"path" yaml:"path"` // 本地文件路径
	Size int64  `mapstructure:"size" json:"size" yaml:"size"` // 本地文件大小限制
}
