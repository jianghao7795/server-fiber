package config

// 任务 model

type Timer struct {
	Start  bool     `mapstructure:"start" json:"start" yaml:"start"`    // 是否启用
	Spec   string   `mapstructure:"spec" json:"spec" yaml:"spec"`       // CRON表达式
	Detail []Detail `mapstructure:"detail" json:"detail" yaml:"detail"` // 任务: 清除数据库
	Tasks  []Task   `mapstructure:"tasks" json:"tasks" yaml:"tasks"`    // 任务： 执行任务
}

type Detail struct {
	TableName    string `mapstructure:"tableName" json:"tableName" yaml:"tableName"`          // 需要清理的表名
	CompareField string `mapstructure:"compareField" json:"compareField" yaml:"compareField"` // 需要比较时间的字段
	Interval     string `mapstructure:"interval" json:"interval" yaml:"interval"`             // 时间间隔
}

type Task struct {
	TaskName string `mapstructure:"taskName" json:"taskName" yaml:"taskName"` // 执行任务名
	Output   string `mapstructure:"output" json:"output" yaml:"output"`       // 输出
	Interval string `mapstructure:"interval" json:"interval" yaml:"interval"` // 时间间隔 秒
}
