package system

import (
	v1 "server-fiber/api/v1"

	"github.com/gofiber/fiber/v2"
)

type GithubRouter struct{}

func (g *GithubRouter) InitGithubRouter(Router fiber.Router) {
	githubRouter := Router.Group("github") //.Use(middleware.OperationRecord)
	githubRouterRecord := Router.Group("github")

	githubRouterApi := v1.ApiGroupApp.SystemApiGroup.SystemGithubApi
	{
		githubRouter.Get("createGithub", githubRouterApi.CreateGithub) // 创建github
	}
	{
		githubRouterRecord.Get("getGithubList", githubRouterApi.GetGithubList)
	}
}
