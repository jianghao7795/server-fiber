package router

import (
	"server-fiber/router/app"
	"server-fiber/router/example"
	"server-fiber/router/frontend"
	"server-fiber/router/mobile"
	"server-fiber/router/system"
)

type appGroup struct {
	app.ArticleRouter
	app.CommentRouter
	app.BaseMessageRouter
	app.UserRouter
	app.TaskRouter
	app.TagRouter
}

var AppRouter = new(appGroup)

type exampleGroup struct {
	example.CustomerRouter
	example.ExcelRouter
	example.FileUploadAndDownloadRouter
}

var ExampleRouter = new(exampleGroup)

type frontendGroup struct {
	frontend.FrontendRouter
}

var FrontendRouter = new(frontendGroup)

type mobileGroup struct {
	mobile.MobileLoginRouter
	mobile.MobileUserRouter
}

var MobileRouter = new(mobileGroup)

type systemGroup struct {
	system.ApiRouter
	system.GithubRouter
	system.AuthorityBtnRouter
	system.AuthorityRouter
	system.AutoCodeHistoryRouter
	system.AutoCodeRouter
	system.BaseRouter
	system.CasbinRouter
	system.DictionaryDetailRouter
	system.DictionaryRouter
	system.InitRouter
	system.JwtRouter
	system.MenuRouter
	system.OperationRecordRouter
	system.ProblemRouter
	system.SysRouter
	system.UserRouter
}

var SystemRouter = new(systemGroup)
