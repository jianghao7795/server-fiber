package app

import (
	v1 "server-fiber/api/v1"
	"server-fiber/middleware"

	"github.com/gofiber/fiber/v2"
)

type TaskRouter struct{}

func (t *TaskRouter) InitTaskRouter(Router fiber.Router) {
	taskRouterWithoutRecord := Router.Group("tasking")
	var taskApi = v1.ApiGroupApp.AppApiGroup.TaskNameApi
	{
		taskRouterWithoutRecord.Get("start", middleware.OperationRecord, taskApi.StartTasking) // 开启任务
	}
}
