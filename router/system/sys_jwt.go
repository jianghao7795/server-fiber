package system

import (
	v1 "server-fiber/api/v1/system"
	"server-fiber/middleware"

	"github.com/gofiber/fiber/v2"
)

type JwtRouter struct{}

func (s *JwtRouter) InitJwtRouter(Router fiber.Router) {
	jwtRouter := Router.Group("jwt")
	jwtApi := new(v1.JwtApi)

	jwtRouter.Post("jsonInBlacklist", middleware.OperationRecord, jwtApi.JsonInBlacklist) // jwt加入黑名单

}
