package system

import (
	v1 "server-fiber/api/v1"

	"github.com/gofiber/fiber/v2"
)

type ProblemRouter struct{}

func (*ProblemRouter) InitProblemRouter(Router fiber.Router) {
	problemRouter := Router.Group("problem")
	var problemApi = v1.ApiGroupApp.SystemApiGroup.UserProblem
	{
		problemRouter.Get("getProblemList/:id", problemApi.GetProblemSetting)
		problemRouter.Put("updateProblem", problemApi.UpdateProblemSetting)

		problemRouter.Get("getIsSetting/:uid", problemApi.HasSetting)
		problemRouter.Post("verifyAnswer", problemApi.VerifyAnswer)
	}
}
