package system

import (
	v1 "server-fiber/api/v1/system"
	"server-fiber/middleware"

	"github.com/gofiber/fiber/v2"
)

type DictionaryDetailRouter struct{}

func (s *DictionaryDetailRouter) InitSysDictionaryDetailRouter(Router fiber.Router) {
	dictionaryDetailRouter := Router.Group("sysDictionaryDetail")
	sysDictionaryDetailApi := new(v1.DictionaryDetailApi)

	dictionaryDetailRouter.Post("createSysDictionaryDetail", middleware.OperationRecord, sysDictionaryDetailApi.CreateSysDictionaryDetail)   // 新建SysDictionaryDetail
	dictionaryDetailRouter.Delete("deleteSysDictionaryDetail", middleware.OperationRecord, sysDictionaryDetailApi.DeleteSysDictionaryDetail) // 删除SysDictionaryDetail
	dictionaryDetailRouter.Put("updateSysDictionaryDetail", middleware.OperationRecord, sysDictionaryDetailApi.UpdateSysDictionaryDetail)    // 更新SysDictionaryDetail

	dictionaryDetailRouter.Get("findSysDictionaryDetail", sysDictionaryDetailApi.FindSysDictionaryDetail)       // 根据ID获取SysDictionaryDetail
	dictionaryDetailRouter.Get("getSysDictionaryDetailList", sysDictionaryDetailApi.GetSysDictionaryDetailList) // 获取SysDictionaryDetail列表

}
