package system

import (
	v1 "server-fiber/api/v1/system"
	"server-fiber/middleware"

	"github.com/gofiber/fiber/v2"
)

type CasbinRouter struct{}

func (s *CasbinRouter) InitCasbinRouter(Router fiber.Router) {
	casbinRouter := Router.Group("casbin")
	casbinApi := new(v1.CasbinApi)

	casbinRouter.Post("updateCasbin", middleware.OperationRecord, casbinApi.UpdateCasbin)

	casbinRouter.Get("getPolicyPathByAuthorityId/:id", casbinApi.GetPolicyPathByAuthorityId)

}
