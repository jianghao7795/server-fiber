package system

import (
	v1 "server-fiber/api/v1"
	"server-fiber/middleware"

	"github.com/gofiber/fiber/v2"
)

type DictionaryDetailRouter struct{}

func (s *DictionaryDetailRouter) InitSysDictionaryDetailRouter(Router fiber.Router) {
	dictionaryDetailRouter := Router.Group("sysDictionaryDetail").Use(middleware.OperationRecord)
	dictionaryDetailRouterWithoutRecord := Router.Group("sysDictionaryDetail")
	sysDictionaryDetailApi := v1.ApiGroupApp.SystemApiGroup.DictionaryDetailApi
	{
		dictionaryDetailRouter.Post("createSysDictionaryDetail", sysDictionaryDetailApi.CreateSysDictionaryDetail)   // 新建SysDictionaryDetail
		dictionaryDetailRouter.Delete("deleteSysDictionaryDetail", sysDictionaryDetailApi.DeleteSysDictionaryDetail) // 删除SysDictionaryDetail
		dictionaryDetailRouter.Put("updateSysDictionaryDetail", sysDictionaryDetailApi.UpdateSysDictionaryDetail)    // 更新SysDictionaryDetail
	}
	{
		dictionaryDetailRouterWithoutRecord.Get("findSysDictionaryDetail", sysDictionaryDetailApi.FindSysDictionaryDetail)       // 根据ID获取SysDictionaryDetail
		dictionaryDetailRouterWithoutRecord.Get("getSysDictionaryDetailList", sysDictionaryDetailApi.GetSysDictionaryDetailList) // 获取SysDictionaryDetail列表
	}
}
