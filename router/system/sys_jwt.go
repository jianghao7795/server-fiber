package system

import (
	v1 "server-fiber/api/v1"

	"github.com/gofiber/fiber/v2"
)

type JwtRouter struct{}

func (s *JwtRouter) InitJwtRouter(Router fiber.Router) {
	jwtRouter := Router.Group("jwt")
	jwtApi := v1.ApiGroupApp.SystemApiGroup.JwtApi
	{
		jwtRouter.Post("jsonInBlacklist", jwtApi.JsonInBlacklist) // jwt加入黑名单
	}
}
