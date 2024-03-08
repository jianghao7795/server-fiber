package system

import (
	v1 "server-fiber/api/v1"

	"github.com/gofiber/fiber/v2"
)

type AuthorityRouter struct{}

func (s *AuthorityRouter) InitAuthorityRouter(Router fiber.Router) {
	authorityRouter := Router.Group("authority") //.Use(middleware.OperationRecord)
	authorityRouterWithoutRecord := Router.Group("authority")
	authorityApi := v1.ApiGroupApp.SystemApiGroup.AuthorityApi
	{
		authorityRouter.Post("createAuthority", authorityApi.CreateAuthority)   // 创建角色
		authorityRouter.Delete("DeleteAuthority", authorityApi.DeleteAuthority) // 删除角色
		authorityRouter.Put("updateAuthority", authorityApi.UpdateAuthority)    // 更新角色
		authorityRouter.Post("copyAuthority", authorityApi.CopyAuthority)       // 拷贝角色
		authorityRouter.Post("setDataAuthority", authorityApi.SetDataAuthority) // 设置角色资源权限
	}
	{
		authorityRouterWithoutRecord.Get("getAuthorityList", authorityApi.GetAuthorityList) // 获取角色列表
	}
}
