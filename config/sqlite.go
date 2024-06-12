package config

type Sqlite struct {
	Path    string `mapstructure:"path" json:"path" yaml:"path"`             // 文件路径
	LogMode string `mapstructure:"log-mode" json:"log-mode" yaml:"log-mode"` // 是否开启Gorm全局日志
}

func (m *Sqlite) GetLogMode() string {
	return m.LogMode
}
