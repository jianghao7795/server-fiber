package app

import (
	v1 "server-fiber/api/v1"
	"server-fiber/middleware"

	"github.com/gofiber/fiber/v2"
)

type UserRouter struct{}

// InitUserRouter 初始化 User 路由信息
func (s *UserRouter) InitUserRouter(Router fiber.Router) {
	userRouter := Router.Group("frontend-user").Use(middleware.OperationRecord)
	userRouterWithoutRecord := Router.Group("frontend-user")
	var userApi = v1.ApiGroupApp.AppApiGroup.UserApi
	{
		userRouter.Post("createUser", userApi.CreateUser)             // 新建User
		userRouter.Delete("deleteUser/:id", userApi.DeleteUser)       // 删除User
		userRouter.Delete("deleteUserByIds", userApi.DeleteUserByIds) // 批量删除User
		userRouter.Put("updateUser/:id", userApi.UpdateUser)          // 更新User
	}
	{
		userRouterWithoutRecord.Get("findUser/:id", userApi.FindUser)   // 根据ID获取User
		userRouterWithoutRecord.Get("getUserList", userApi.GetUserList) // 获取User列表
	}
}
