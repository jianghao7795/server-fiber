package system

import (
	v1 "server-fiber/api/v1/system"

	"github.com/gofiber/fiber/v2"
)

type OperationRecordRouter struct{}

func (s *OperationRecordRouter) InitSysOperationRecordRouter(Router fiber.Router) {
	operationRecordRouter := Router.Group("sysOperationRecord")
	authorityMenuApi := new(v1.OperationRecordApi)

	operationRecordRouter.Post("createSysOperationRecord", authorityMenuApi.CreateSysOperationRecord)             // 新建SysOperationRecord
	operationRecordRouter.Delete("deleteSysOperationRecord", authorityMenuApi.DeleteSysOperationRecord)           // 删除SysOperationRecord
	operationRecordRouter.Delete("deleteSysOperationRecordByIds", authorityMenuApi.DeleteSysOperationRecordByIds) // 批量删除SysOperationRecord
	operationRecordRouter.Get("findSysOperationRecord", authorityMenuApi.FindSysOperationRecord)                  // 根据ID获取SysOperationRecord
	operationRecordRouter.Get("getSysOperationRecordList", authorityMenuApi.GetSysOperationRecordList)            // 获取SysOperationRecord列表
}
