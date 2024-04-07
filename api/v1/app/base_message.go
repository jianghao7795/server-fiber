package app

import (
	"errors"

	"server-fiber/global"
	"server-fiber/model/app"
	"server-fiber/model/common/response"
	appService "server-fiber/service/app"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type BaseMessageApi struct{}

var baseMessageService = new(appService.BaseMessageService)

// CreateBaseMessage 创建base_message
func (a *BaseMessageApi) CreateBaseMessage(c *fiber.Ctx) error {
	var baseMessage app.BaseMessage
	err := c.BodyParser(&baseMessage)
	if err != nil {
		global.LOG.Error("获取数据失败!", zap.Error(err))
		return response.FailWithMessage("获取数据失败", c)
	}
	if err := baseMessageService.CreateBaseMessage(&baseMessage); err != nil {
		global.LOG.Error("创建失败!", zap.Error(err))
		return response.FailWithMessage("创建失败", c)
	} else {
		return response.OkWithId("创建成功", baseMessage.ID, c)
	}
}

/**
 * @description: update base message
 * @param {*fiber.Ctx} c
 * @return {*}
 */

func (a *BaseMessageApi) UpdateBaseMessage(c *fiber.Ctx) error {
	var baseMessage app.BaseMessage
	err := c.BodyParser(&baseMessage)
	if err != nil {
		global.LOG.Error("获取数据失败!", zap.Error(err))
		return response.FailWithMessage("获取数据失败", c)
	}
	if err = baseMessageService.UpdateBaseMessage(&baseMessage); err != nil {
		global.LOG.Error("更新失败!", zap.Error(err))
		return response.FailWithMessage("更新失败", c)
	} else {
		return response.OkWithMessage("更新成功", c)
	}
}

/**
 * @description: find base message
 * @param {*fiber.Ctx} c
 * @return nil
 */
func (a *BaseMessageApi) FindBaseMessage(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	if responseBaseMessage, err := baseMessageService.FindBaseMessage(uint(id)); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// respBaseMessage := baseMessageNotFound{message: "not found"}
			str := "not found"
			global.LOG.Error("查询失败!", zap.Error(errors.New(str)))
			return response.OkWithData(fiber.Map{"error": str}, c)
		} else {
			global.LOG.Error("查询失败!", zap.Error(err))
			return response.FailWithMessage("查询失败", c)
		}

	} else {
		return response.OkWithData(fiber.Map{"baseMessage": responseBaseMessage}, c)
	}
}
