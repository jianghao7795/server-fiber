package example

import (
	global "server-fiber/model"
	"server-fiber/model/system"
)

type ExaCustomer struct {
	global.MODEL
	CustomerName       string         `json:"customerName" form:"customerName" gorm:"comment:客户名"`                // 客户名
	CustomerPhoneData  string         `json:"customerPhoneData" form:"customerPhoneData" gorm:"comment:客户手机号"`    // 客户手机号
	SysUserID          uint           `json:"sysUserId" form:"sysUserId" gorm:"comment:管理ID"`                     // 管理ID
	SysUserAuthorityID string         `json:"sysUserAuthorityID" form:"sysUserAuthorityID" gorm:"comment:管理角色ID"` // 管理角色ID
	SysUser            system.SysUser `json:"sysUser" form:"sysUser" gorm:"comment:管理详情"`                         // 管理详情
}

func (ExaCustomer) TableName() string {
	return "exa_customers"
}
