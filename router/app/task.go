package app

import (
	v1 "server-fiber/api/v1"

	"github.com/gofiber/fiber/v2"
)

type TaskRouter struct{}

func (t *TaskRouter) InitTaskRouter(Router fiber.Router) {
	// articleRouter := Router.Group("article").Use(middleware.OperationRecord())
	taskRouterWithoutRecord := Router.Group("tasking")
	var taskApi = v1.ApiGroupApp.AppApiGroup.TaskNameApi
	{
		taskRouterWithoutRecord.Get("start", taskApi.StartTasking) // 开启任务
	}
}
