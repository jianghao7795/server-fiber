package system

import (
	"errors"
	"fmt"

	"server-fiber/global"
	"server-fiber/model/common/request"
	"server-fiber/model/system"
	systemReq "server-fiber/model/system/request"

	"gorm.io/gorm"
)

//
//@function: CreateApi
//@description: 新增基础api
//@param: api model.SysApi
//@return: err error

type ApiService struct{}

var ApiServiceApp = new(ApiService)

func (apiService *ApiService) CreateApi(api *system.SysApi) (err error) {
	if !errors.Is(global.DB.Where("path = ? AND method = ?", api.Path, api.Method).First(&system.SysApi{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在相同api")
	}
	return global.DB.Create(api).Error
}

//
//@function: DeleteApi
//@description: 删除基础api
//@param: api model.SysApi
//@return: err error

func (apiService *ApiService) DeleteApi(api system.SysApi) (err error) {
	err = global.DB.Delete(&api).Error
	CasbinServiceApp.ClearCasbin(1, api.Path, api.Method)
	return err
}

//
//@function: GetAPIInfoList
//@description: 分页获取数据,
//@param: api model.SysApi, info request.PageInfo, order string, desc bool
//@return: err error

func (apiService *ApiService) GetAPIInfoList(info systemReq.SearchApiParams) (list []system.SysApi, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.DB.Model(&system.SysApi{})
	var apiList []system.SysApi

	if info.Path != "" {
		db = db.Where("path LIKE ?", "%"+info.Path+"%")
	}

	if info.Description != "" {
		db = db.Where("description LIKE ?", "%"+info.Description+"%")
	}

	if info.Method != "" {
		db = db.Where("method = ?", info.Method)
	}

	if info.ApiGroup != "" {
		db = db.Where("api_group = ?", info.ApiGroup)
	}
	err = db.Count(&total).Error

	if err != nil {
		return apiList, total, err
	} else {
		db = db.Limit(limit).Offset(offset)
		var sort = "api_group"
		if info.OrderKey != "" {
			if info.Desc == "true" {
				sort = fmt.Sprintf("%s %s", info.OrderKey, "desc")
			} else {
				sort = info.OrderKey
			}

		}
		err = db.Order(sort).Find(&apiList).Error
		// if order != "" {
		// 	var OrderStr string
		// 	// 设置有效排序key 防止sql注入
		// 	// 感谢 Tom4t0 提交漏洞信息
		// 	orderMap := make(map[string]bool, 5)
		// 	orderMap["id"] = true
		// 	orderMap["path"] = true
		// 	orderMap["api_group"] = true
		// 	orderMap["description"] = true
		// 	orderMap["method"] = true
		// 	if orderMap[order] {
		// 		if desc {
		// 			OrderStr = order + " desc"
		// 		} else {
		// 			OrderStr = order
		// 		}

		// 	} else { // didn't matched any order key in `orderMap`
		// 		err = fmt.Errorf("非法的排序字段: %v", order)
		// 		return apiList, total, err
		// 	}
		// err = db.Order(OrderStr).Find(&apiList).Error
		// } else {
		// 	// err = db.Order("api_group").Find(&apiList).Error
		// }
	}
	return apiList, total, err
}

//
//@function: GetAllApis
//@description: 获取所有的api
//@return: err error, apis []model.SysApi

func (apiService *ApiService) GetAllApis() (apis []system.SysApi, err error) {
	err = global.DB.Find(&apis).Error
	return
}

//
//@function: GetApiById
//@description: 根据id获取api
//@param: id float64
//@return: err error, api model.SysApi

func (apiService *ApiService) GetApiById(id int) (api system.SysApi, err error) {
	err = global.DB.Where("id = ?", id).First(&api).Error
	return
}

//
//@function: UpdateApi
//@description: 根据id更新api
//@param: api model.SysApi
//@return: err error

func (apiService *ApiService) UpdateApi(api *system.SysApi) (err error) {
	var oldA system.SysApi
	err = global.DB.Where("id = ?", api.ID).First(&oldA).Error
	if err != nil {
		return err
	}
	if oldA.Path != api.Path || oldA.Method != api.Method {
		if !errors.Is(global.DB.Where("path = ? AND method = ?", api.Path, api.Method).First(&system.SysApi{}).Error, gorm.ErrRecordNotFound) {
			return errors.New("存在相同api路径")
		}
	}
	return global.DB.Save(api).Error
}

//
//@function: DeleteApis
//@description: 删除选中API
//@param: apis []model.SysApi
//@return: err error

func (apiService *ApiService) DeleteApisByIds(ids request.IdsReq) (err error) {
	err = global.DB.Delete(&[]system.SysApi{}, "id in ?", ids.Ids).Error
	return err
}

func (apiService *ApiService) DeleteApiByIds(ids []string) (err error) {
	return global.DB.Delete(&system.SysApi{}, "id in ?", ids).Error
}
