package system

import (
	v1 "server-fiber/api/v1"

	"github.com/gofiber/fiber/v2"
)

type InitRouter struct{}

func (s *InitRouter) InitInitRouter(Router fiber.Router) {
	initRouter := Router.Group("init")
	dbApi := v1.ApiGroupApp.SystemApiGroup.DBApi
	{
		initRouter.Post("initdb", dbApi.InitDB)  // 创建Api
		initRouter.Get("checkdb", dbApi.CheckDB) // 检查是否初始化
	}
}
