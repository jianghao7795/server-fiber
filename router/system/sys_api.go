package system

import (
	v1 "server-fiber/api/v1"
	"server-fiber/middleware"

	"github.com/gofiber/fiber/v2"
)

type ApiRouter struct{}

func (s *ApiRouter) InitApiRouter(Router fiber.Router) {
	apiRouter := Router.Group("api").Use(middleware.OperationRecord)
	apiRouterWithoutRecord := Router.Group("api")
	apiRouterApi := v1.ApiGroupApp.SystemApiGroup.SystemApiApi
	{
		apiRouter.Post("createApi", apiRouterApi.CreateApi)               // 创建Api
		apiRouter.Delete("DeleteApi/:id", apiRouterApi.DeleteApi)         // 删除Api
		apiRouter.Put("updateApi/:id", apiRouterApi.UpdateApi)            // 更新api
		apiRouter.Delete("DeleteApisByIds", apiRouterApi.DeleteApisByIds) // 删除选中api
	}
	{
		apiRouterWithoutRecord.Get("getApiById/:id", apiRouterApi.GetApiById) // 获取单条Api消息
		apiRouterWithoutRecord.Get("getAllApis", apiRouterApi.GetAllApis)     // 获取所有api
		apiRouterWithoutRecord.Get("getApiList", apiRouterApi.GetApiList)     // 获取Api列表
	}
}
