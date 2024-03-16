package mobile

import (
	"server-fiber/global"
	"server-fiber/model/common/response"
	"server-fiber/model/mobile"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type RegisterMobile struct{}

func (*RegisterMobile) Register(c *fiber.Ctx) error {
	var data mobile.Register
	_ = c.BodyParser(&data)
	if err := mobileRegisterService.Register(data); err != nil {
		global.LOG.Error("注册失败!", zap.Error(err))
		return response.FailWithMessage400("注册失败，请重试", c)
	} else {
		return response.OkWithDetailed("", "注册成功", c)
	}
}
