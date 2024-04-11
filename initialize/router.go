/*
 * @Author: jianghao
 * @Date: 2022-07-29 09:48:24
 * @LastEditors: jianghao
 * @LastEditTime: 2022-10-17 11:27:44
 */
package initialize

import (
	"time"

	"server-fiber/global"
	"server-fiber/middleware"
	"server-fiber/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// 初始化总路由

func Routers() *fiber.App {
	app := fiber.New(global.CONFIG.FiberConfig)
	app.Use(logger.New(global.CONFIG.FiberLogger)) //log 日志配置
	appRouter := router.AppRouter
	systemRouter := router.SystemRouter
	exampleRouter := router.ExampleRouter
	mobile := router.MobileRouter
	app.Static("/api/uploads/", "uploads/", fiber.Static{
		Compress:      true,
		ByteRange:     true,
		Browse:        true,
		CacheDuration: 100 * time.Second,
		MaxAge:        3600,
	}) // 本地的frontend api文件路由转化
	app.Static("/backend/uploads/", "uploads/", fiber.Static{
		Compress:      true,
		ByteRange:     true,
		Browse:        true,
		CacheDuration: 100 * time.Second,
		MaxAge:        3600,
	}) // 本地的backend文件路由转化
	app.Static("/mobile/uploads/", "uploads/", fiber.Static{
		Compress:      true,
		ByteRange:     true,
		Browse:        true,
		CacheDuration: 100 * time.Second,
		MaxAge:        3600,
	}) // 本地的mobile文件路由转化
	app.Static("/backend/form-generator", "resource/page")
	// app.Static("/form-generator", "./resource/page") // 生成form组件的js代码
	// Router.Use(middleware.Cors())        // 直接放行全部跨域请求
	app.Use(middleware.CorsByRules) // 按照配置的规则放行跨域请求
	{
		backendRooterNotLogin := app.Group("/backend")
		systemRouter.InitBaseRouter(backendRooterNotLogin) // 注册基础功能路由 不做鉴权
		systemRouter.InitInitRouter(backendRooterNotLogin) // 自动初始化相关

		backendRooter := backendRooterNotLogin.Use(middleware.JWTAuth, middleware.CasbinHandler) // casbin的拦截规则
		{
			systemRouter.InitApiRouter(backendRooter)                 // 注册功能api路由
			systemRouter.InitJwtRouter(backendRooter)                 // jwt相关路由
			systemRouter.InitUserRouter(backendRooter)                // 注册用户路由
			systemRouter.InitMenuRouter(backendRooter)                // 注册menu路由
			systemRouter.InitSystemRouter(backendRooter)              // system相关路由
			systemRouter.InitCasbinRouter(backendRooter)              // 权限相关路由
			systemRouter.InitAutoCodeRouter(backendRooter)            // 创建自动化代码
			systemRouter.InitAuthorityRouter(backendRooter)           // 注册角色路由
			systemRouter.InitSysDictionaryRouter(backendRooter)       // 字典管理
			systemRouter.InitAutoCodeHistoryRouter(backendRooter)     // 自动化代码历史
			systemRouter.InitSysOperationRecordRouter(backendRooter)  // 操作记录
			systemRouter.InitSysDictionaryDetailRouter(backendRooter) // 字典详情管理
			systemRouter.InitAuthorityBtnRouterRouter(backendRooter)  // 字典详情管理
			systemRouter.InitProblemRouter(backendRooter)             // problem
			systemRouter.InitGithubRouter(backendRooter)              // github commit

			exampleRouter.InitExcelRouter(backendRooter)                 // 表格导入导出
			exampleRouter.InitCustomerRouter(backendRooter)              // 客户路由
			exampleRouter.InitFileUploadAndDownloadRouter(backendRooter) // 文件上传下载功能路由

			// Code generated by server Begin; DO NOT EDIT.

			appRouter.InitTagRouter(backendRooter)         // tab
			appRouter.InitArticleRouter(backendRooter)     //article
			appRouter.InitCommentRouter(backendRooter)     // comment
			appRouter.InitBaseMessageRouter(backendRooter) // baseMessage
			appRouter.InitTaskRouter(backendRooter)        //task 任务
			appRouter.InitUserRouter(backendRooter)        // frontend user
			mobile.InitMobileRouter(backendRooter)
			// Code generated by server End; DO NOT EDIT.
		}
		frontendRouter := router.FrontendRouter
		PublicGroup := app.Group("api")
		{
			// 前台的API
			frontendRouter.InitFrontendRouter(PublicGroup)
		}
		// MobleGroup := app.Group("mobile")
		// {
		// 	mobile.InitMobileRouter(MobleGroup)
		// }
		InstallPlugin(backendRooter, PublicGroup)
	}

	return app
}
