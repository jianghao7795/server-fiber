package system

import (
	v1 "server-fiber/api/v1"

	"github.com/gofiber/fiber/v2"
)

type AuthorityBtnRouter struct{}

func (s *AuthorityBtnRouter) InitAuthorityBtnRouterRouter(Router fiber.Router) {
	authorityRouter := Router.Group("authorityBtn") //.Use(middleware.OperationRecord)
	authorityRouterWithoutRecord := Router.Group("authorityBtn")
	authorityBtnApi := v1.ApiGroupApp.SystemApiGroup.AuthorityBtnApi
	{
		authorityRouter.Post("setAuthorityBtn", authorityBtnApi.SetAuthorityBtn)
		authorityRouter.Delete("canRemoveAuthorityBtn/:id", authorityBtnApi.CanRemoveAuthorityBtn)
	}
	{
		authorityRouterWithoutRecord.Get("getAuthorityBtn", authorityBtnApi.GetAuthorityBtn)

	}
}
