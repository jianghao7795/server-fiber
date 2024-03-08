package system

import (
	"server-fiber/global"
)

type SysUserProblem struct {
	global.MODEL
	SysUserId int    `json:"sys_user_id" binding:"required,gte=1" form:"sys_user_id" gorm:"column:sys_user_id;comment:用户的ID"`
	Problem   string `json:"problem" binding:"required,max=255,min=1" form:"problem" gorm:"column:problem;comment:问题"`
	Answer    string `json:"answer" binding:"required,max=255,min=1" form:"answer" gorm:"column:answer;comment:答案"`
}

// TableName Comment 表名
func (SysUserProblem) TableName() string {
	return "sys_user_problem"
}
