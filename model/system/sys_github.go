package system

import (
	"server-fiber/global"
)

type SysGithub struct {
	global.MODEL
	Author     string `json:"author" form:"author" gorm:"comment:提交代码用户"`
	Message    string `json:"message" form:"message" gorm:"comment:提交代码的commit"`
	CommitTime string `json:"commit_time" form:"commit_time" gorm:"comment:提交代码时间"`
}

func (SysGithub) TableName() string {
	return "githubs"
}
