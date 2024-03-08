package system

import (
	v1 "server-fiber/api/v1"

	"github.com/gofiber/fiber/v2"
)

type UserRouter struct{}

func (s *UserRouter) InitUserRouter(Router fiber.Router) {
	userRouter := Router.Group("user") ////.Use(middleware.OperationRecord)
	userRouterWithoutRecord := Router.Group("user")
	baseApi := v1.ApiGroupApp.SystemApiGroup.BaseApi
	{
		userRouter.Post("admin_register", baseApi.Register)               // 管理员注册账号
		userRouter.Post("changePassword", baseApi.ChangePassword)         // 用户修改密码
		userRouter.Post("setUserAuthority", baseApi.SetUserAuthority)     // 设置用户权限
		userRouter.Delete("deleteUser", baseApi.DeleteUser)               // 删除用户
		userRouter.Put("setUserInfo", baseApi.SetUserInfo)                // 设置用户信息
		userRouter.Put("setSelfInfo", baseApi.SetSelfInfo)                // 设置自身信息
		userRouter.Post("setUserAuthorities", baseApi.SetUserAuthorities) // 设置用户权限组
		userRouter.Post("resetPassword", baseApi.ResetPassword)           // 重置密码
	}
	{
		userRouterWithoutRecord.Get("getUserList", baseApi.GetUserList)   // 分页获取用户列表
		userRouterWithoutRecord.Get("getUserInfo", baseApi.GetUserInfo)   // 获取自身信息
		userRouterWithoutRecord.Get("getUserCount", baseApi.GetUserCount) // 获取用户数
		userRouterWithoutRecord.Get("getFlowmeter", baseApi.GetFlowmeter) // 获取摸个ip流量
	}
}
