package system

import (
	v1 "server-fiber/api/v1"
	"server-fiber/middleware"

	"github.com/gofiber/fiber/v2"
)

type BaseRouter struct{}

func (s *BaseRouter) InitBaseRouter(Router fiber.Router) {
	baseRouter := Router.Group("base")
	baseApi := v1.ApiGroupApp.SystemApiGroup.BaseApi
	{
		baseRouter.Post("login", baseApi.Login)
		baseRouter.Get("captcha", middleware.NeedInit, baseApi.Captcha)
	}
}
