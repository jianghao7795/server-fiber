package system

import (
	v1 "server-fiber/api/v1"

	"github.com/gofiber/fiber/v2"
)

type MenuRouter struct{}

func (s *MenuRouter) InitMenuRouter(Router fiber.Router) {
	menuRouter := Router.Group("menu") //.Use(middleware.OperationRecord)
	menuRouterWithoutRecord := Router.Group("menu")
	authorityMenuApi := v1.ApiGroupApp.SystemApiGroup.AuthorityMenuApi
	{
		menuRouter.Post("addBaseMenu", authorityMenuApi.AddBaseMenu)             // 新增菜单
		menuRouter.Post("addMenuAuthority", authorityMenuApi.AddMenuAuthority)   //	增加menu和角色关联关系
		menuRouter.Delete("deleteBaseMenu/:id", authorityMenuApi.DeleteBaseMenu) // 删除菜单
		menuRouter.Put("updateBaseMenu", authorityMenuApi.UpdateBaseMenu)        // 更新菜单
	}
	{
		menuRouterWithoutRecord.Get("getMenu", authorityMenuApi.GetMenu)                     // 获取菜单树
		menuRouterWithoutRecord.Get("getMenuList", authorityMenuApi.GetMenuList)             // 分页获取基础menu列表
		menuRouterWithoutRecord.Get("getBaseMenuTree", authorityMenuApi.GetBaseMenuTree)     // 获取用户动态路由
		menuRouterWithoutRecord.Get("getMenuAuthority", authorityMenuApi.GetMenuAuthority)   // 获取指定角色menu
		menuRouterWithoutRecord.Get("getBaseMenuById/:id", authorityMenuApi.GetBaseMenuById) // 根据id获取菜单
	}
}
