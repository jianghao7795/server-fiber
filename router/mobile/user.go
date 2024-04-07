package mobile

import (
	v1 "server-fiber/api/v1/mobile"
	"server-fiber/middleware"

	"github.com/gofiber/fiber/v2"
)

type MobileUserRouter struct{}

func (m *MobileUserRouter) InitMobileRouter(Router fiber.Router) {
	moblieUserRouterWithoutRecord := Router.Group("mobile")
	var mobileUserApi = new(v1.MobileUserApi)

	moblieUserRouterWithoutRecord.Post("createMobileUser", middleware.OperationRecord, mobileUserApi.CreateMoblieUser)             // 新建MoblieUser
	moblieUserRouterWithoutRecord.Delete("deleteMobileUser/:id", middleware.OperationRecord, mobileUserApi.DeleteMoblieUser)       // 删除MoblieUser
	moblieUserRouterWithoutRecord.Delete("deleteMobileUserByIds", middleware.OperationRecord, mobileUserApi.DeleteMoblieUserByIds) // 批量删除MoblieUser
	moblieUserRouterWithoutRecord.Put("updateMobileUser/:id", middleware.OperationRecord, mobileUserApi.UpdateMoblieUser)          // 更新MoblieUser

	moblieUserRouterWithoutRecord.Get("findMobileUser/:id", mobileUserApi.FindMoblieUser)   // 根据ID获取MoblieUser
	moblieUserRouterWithoutRecord.Get("getMobileUserList", mobileUserApi.GetMoblieUserList) // 获取MoblieUser列表

}
