package app

import (
	"errors"

	"server-fiber/global"
	"server-fiber/model/app"
	"server-fiber/model/common/response"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// CreateBaseMessage 创建base_message
// @Tags BaseMessage
// @Summary 创建BaseMessage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body app.BaseMessage true "创建base_message"
// @Success 200 {object} response.Response{msg=string,data=object,code=number} "创建base_message"
// @Router /base_message/createBaseMessage [post]
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

// @Tags BaseMessage
// @Summary 更新 base_message
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body app.BaseMessage true "创建base_message"
// @Success 200 {object} response.Response{msg=string,code=number} "查找base message"
// @Router /base_message/updateBaseMessage/:id [put]
func (a *BaseMessageApi) UpdateBaseMessage(c *fiber.Ctx) error {
	var baseMessage app.BaseMessage
	id, err := c.ParamsInt("id", 0)
	if err != nil {
		global.LOG.Error("获取id失败", zap.Error(err))
		return response.FailWithMessage("获取id失败", c)
	}
	err = c.BodyParser(&baseMessage)
	if err != nil {
		global.LOG.Error("获取数据失败!", zap.Error(err))
		return response.FailWithMessage("获取数据失败", c)
	}
	if err = baseMessageService.UpdateBaseMessage(id, &baseMessage); err != nil {
		global.LOG.Error("更新失败!", zap.Error(err))
		return response.FailWithMessage("更新失败", c)
	} else {
		return response.OkWithMessage("更新成功", c)
	}
}

// @Tags BaseMessage
// @Summary 查找base message
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param id path number true "查找base message"
// @Success 200 {object} response.Response{msg=string,data=app.BaseMessage,code=number} "查找base message"
// @Router /base_message/getBaseMessage/:id [get]
func (a *BaseMessageApi) FindBaseMessage(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		global.LOG.Error("获取id失败!", zap.Error(err))
		return response.FailWithMessage("获取id失败", c)
	}
	if responseBaseMessage, err := baseMessageService.FindBaseMessage(uint(id)); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// respBaseMessage := baseMessageNotFound{message: "not found"}
			str := "base message not found"
			global.LOG.Error("查询失败!", zap.Error(errors.New(str)))
			return response.OkWithData(str, c)
		} else {
			global.LOG.Error("查询失败!", zap.Error(err))
			return response.FailWithMessage("查询失败", c)
		}
	} else {
		return response.OkWithDetailed(responseBaseMessage, "查询成功", c)
	}
}
