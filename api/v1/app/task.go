package app

import (
	"server-fiber/global"
	"server-fiber/model/common/response"

	"github.com/gofiber/fiber/v2"
)

// StartTasking start Tasking
// @Tags start Tasking
// @Summary start Tasking
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"开启成功"}"
// @Router /tasking/start [get]
func (*TaskNameApi) StartTasking(c *fiber.Ctx) error {
	tasking := c.Query("task")
	if tasking == "" {
		global.LOG.Error("请传入任务名!")
		return response.FailWithMessage("请传入任务名", c)
	}
	_, status := global.Timer.FindCron(tasking)

	// log.Println(task)

	if !status {
		global.LOG.Error("开启失败!")
		return response.FailWithMessage("开启失败，没有这个任务", c)
	} else {
		global.Timer.StartTask(tasking)
		return response.OkWithMessage("开启成功", c)
	}
}

func (*TaskNameApi) StopTasking(c *fiber.Ctx) error {
	tasking := c.Query("task")
	if tasking == "" {
		global.LOG.Error("请传入任务名!")
		return response.FailWithMessage("请传入任务名", c)
	}
	_, status := global.Timer.FindCron(tasking)

	// log.Println(task)

	if !status {
		global.LOG.Error("关闭失败!")
		return response.FailWithMessage("关闭失败，没有这个任务", c)
	} else {
		global.Timer.StopTask(tasking)
		return response.OkWithMessage("关闭成功", c)
	}
}
