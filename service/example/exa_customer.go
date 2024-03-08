package example

import (
	"server-fiber/global"
	"server-fiber/model/example"
	"server-fiber/model/example/request"
	"server-fiber/model/system"
	systemService "server-fiber/service/system"
	"strings"
)

type CustomerService struct{}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateExaCustomer
//@description: 创建客户
//@param: e model.ExaCustomer
//@return: err error

func (exa *CustomerService) CreateExaCustomer(e example.ExaCustomer) (err error) {
	err = global.DB.Create(&e).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteFileChunk
//@description: 删除客户
//@param: e model.ExaCustomer
//@return: err error

func (exa *CustomerService) DeleteExaCustomer(e example.ExaCustomer) (err error) {
	err = global.DB.Delete(&e).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateExaCustomer
//@description: 更新客户
//@param: e *model.ExaCustomer
//@return: err error

func (exa *CustomerService) UpdateExaCustomer(e *example.ExaCustomer) (err error) {
	err = global.DB.Save(e).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetExaCustomer
//@description: 获取客户信息
//@param: id uint
//@return: err error, customer model.ExaCustomer

func (exa *CustomerService) GetExaCustomer(id uint) (customer example.ExaCustomer, err error) {
	err = global.DB.Where("id = ?", id).First(&customer).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetCustomerInfoList
//@description: 分页获取客户列表
//@param: sysUserAuthorityID string, info request.PageInfo
//@return: err error, list interface{}, total int64

func (exa *CustomerService) GetCustomerInfoList(sysUserAuthorityID string, info request.SearchCustomerParams) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.DB.Model(&example.ExaCustomer{})
	var a system.SysAuthority
	a.AuthorityId = sysUserAuthorityID
	auth, err := systemService.AuthorityServiceApp.GetAuthorityInfo(a)
	if err != nil {
		return
	}
	var dataId []string
	for _, v := range auth.DataAuthorityId {
		dataId = append(dataId, v.AuthorityId)
	}
	var CustomerList []example.ExaCustomer
	if info.CustomerName != "" {
		var build strings.Builder
		build.WriteString("%")
		build.WriteString(info.CustomerName)
		build.WriteString("%")
		var customername string = build.String()
		db = db.Where("customer_name like ?", customername)
	}
	if info.CustomerPhoneData != "" {
		db = db.Where("customer_phone_data = ?", info.CustomerPhoneData)
	}
	err = db.Where("sys_user_authority_id in ?", dataId).Count(&total).Error
	if err != nil {
		return CustomerList, total, err
	} else {
		err = db.Order("id desc").Limit(limit).Offset(offset).Preload("SysUser").Where("sys_user_authority_id in ?", dataId).Find(&CustomerList).Error
	}
	return CustomerList, total, err
}
