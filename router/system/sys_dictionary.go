package system

import (
	v1 "server-fiber/api/v1"
	"server-fiber/middleware"

	"github.com/gofiber/fiber/v2"
)

type DictionaryRouter struct{}

func (s *DictionaryRouter) InitSysDictionaryRouter(Router fiber.Router) {
	sysDictionaryRouter := Router.Group("sysDictionary").Use(middleware.OperationRecord)
	sysDictionaryRouterWithoutRecord := Router.Group("sysDictionary")
	sysDictionaryApi := v1.ApiGroupApp.SystemApiGroup.DictionaryApi
	{
		sysDictionaryRouter.Post("createSysDictionary", sysDictionaryApi.CreateSysDictionary)   // 新建SysDictionary
		sysDictionaryRouter.Delete("deleteSysDictionary", sysDictionaryApi.DeleteSysDictionary) // 删除SysDictionary
		sysDictionaryRouter.Put("updateSysDictionary", sysDictionaryApi.UpdateSysDictionary)    // 更新SysDictionary
	}
	{
		sysDictionaryRouterWithoutRecord.Get("findSysDictionary", sysDictionaryApi.FindSysDictionary)       // 根据ID获取SysDictionary
		sysDictionaryRouterWithoutRecord.Get("getSysDictionaryList", sysDictionaryApi.GetSysDictionaryList) // 获取SysDictionary列表
	}
}
