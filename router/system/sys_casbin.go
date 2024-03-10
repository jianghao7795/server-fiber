package system

import (
	v1 "server-fiber/api/v1"
	"server-fiber/middleware"

	"github.com/gofiber/fiber/v2"
)

type CasbinRouter struct{}

func (s *CasbinRouter) InitCasbinRouter(Router fiber.Router) {
	casbinRouter := Router.Group("casbin").Use(middleware.OperationRecord)
	// casbinRouterWithoutRecord := Router.Group("casbin")
	casbinApi := v1.ApiGroupApp.SystemApiGroup.CasbinApi
	{
		casbinRouter.Post("updateCasbin", casbinApi.UpdateCasbin)
	}
	{
		casbinRouter.Get("getPolicyPathByAuthorityId/:id", casbinApi.GetPolicyPathByAuthorityId)
	}
}
