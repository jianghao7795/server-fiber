package app

import (
	"errors"
	"strconv"

	"server-fiber/global"
	"server-fiber/model/app"
	"server-fiber/model/common/response"
	"server-fiber/service"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type BaseMessageApi struct{}

var baseMessageService = service.ServiceGroupApp.AppServiceGroup.BaseMessageService

// CreateBaseMessage 创建base_message
func (a *BaseMessageApi) CreateBaseMessage(c *fiber.Ctx) error {
	var baseMessage app.BaseMessage
	_ = c.BodyParser(&baseMessage)
	if err := baseMessageService.CreateBaseMessage(baseMessage); err != nil {
		global.LOG.Error("创建失败!", zap.Error(err))
		return response.FailWithMessage("创建失败", c)
	} else {
		return response.OkWithMessage("创建成功", c)
	}
}

/**
 * @description: update base message
 * @param {*fiber.Ctx} c
 * @return {*}
 */

func (a *BaseMessageApi) UpdateBaseMessage(c *fiber.Ctx) error {
	var baseMessage app.BaseMessage
	_ = c.BodyParser(&baseMessage)
	if err := baseMessageService.UpdateBaseMessage(baseMessage); err != nil {
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
	id := c.Params("id")
	ids, _ := strconv.Atoi(id)
	if responseBaseMessage, err := baseMessageService.FindBaseMessage(uint(ids)); err != nil {
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
