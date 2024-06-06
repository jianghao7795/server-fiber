package system

import (
	v1 "server-fiber/api/v1/system"
	"server-fiber/middleware"

	"github.com/gofiber/fiber/v2"
)

type DictionaryRouter struct{}

func (s *DictionaryRouter) InitSysDictionaryRouter(Router fiber.Router) {
	sysDictionaryRouter := Router.Group("sysDictionary")
	sysDictionaryApi := new(v1.DictionaryApi)

	sysDictionaryRouter.Post("createSysDictionary", middleware.OperationRecord, sysDictionaryApi.CreateSysDictionary)       // 新建SysDictionary
	sysDictionaryRouter.Delete("deleteSysDictionary/:id", middleware.OperationRecord, sysDictionaryApi.DeleteSysDictionary) // 删除SysDictionary
	sysDictionaryRouter.Put("updateSysDictionary/:id", middleware.OperationRecord, sysDictionaryApi.UpdateSysDictionary)    // 更新SysDictionary

	sysDictionaryRouter.Get("findSysDictionary/:id", sysDictionaryApi.FindSysDictionary)   // 根据ID获取SysDictionary
	sysDictionaryRouter.Get("getSysDictionaryList", sysDictionaryApi.GetSysDictionaryList) // 获取SysDictionary列表
}
