package system

import (
	"server-fiber/global"
	"server-fiber/model/common/request"
	"server-fiber/model/common/response"
	"server-fiber/model/system"
	systemReq "server-fiber/model/system/request"
	systemRes "server-fiber/model/system/response"
	"server-fiber/utils"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type AuthorityMenuApi struct{}

// @Tags AuthorityMenu
// @Summary 获取用户动态路由
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body request.Empty true "空"
// @Success 200 {object} response.Response{data=systemRes.SysMenusResponse,msg=string} "获取用户动态路由,返回包括系统菜单详情列表"
// @Router /menu/getMenu [post]
func (a *AuthorityMenuApi) GetMenu(c *fiber.Ctx) error {
	if menus, err := menuService.GetMenuTree(utils.GetUserAuthorityId(c)); err != nil {
		global.LOG.Error("获取失败!", zap.Error(err))
		return response.FailWithMessage("获取失败", c)
	} else {
		if menus == nil {
			menus = []system.SysMenu{}
		}
		return response.OkWithDetailed(systemRes.SysMenusResponse{Menus: menus}, "获取成功", c)
	}
}

// @Tags AuthorityMenu
// @Summary 获取用户动态路由
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body request.Empty true "空"
// @Success 200 {object} response.Response{data=systemRes.SysBaseMenusResponse,msg=string} "获取用户动态路由,返回包括系统菜单列表"
// @Router /menu/getBaseMenuTree [get]
func (a *AuthorityMenuApi) GetBaseMenuTree(c *fiber.Ctx) error {
	if menus, err := menuService.GetBaseMenuTree(); err != nil {
		global.LOG.Error("获取失败!", zap.Error(err))
		return response.FailWithMessage("获取失败", c)
	} else {
		return response.OkWithDetailed(systemRes.SysBaseMenusResponse{Menus: menus}, "获取成功", c)
	}
}

// @Tags AuthorityMenu
// @Summary 增加menu和角色关联关系
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body systemReq.AddMenuAuthorityInfo true "角色ID"
// @Success 200 {object} response.Response{msg=string} "增加menu和角色关联关系"
// @Router /menu/addMenuAuthority [post]
func (a *AuthorityMenuApi) AddMenuAuthority(c *fiber.Ctx) error {
	var authorityMenu systemReq.AddMenuAuthorityInfo
	_ = c.BodyParser(&authorityMenu)
	if err := utils.Verify(authorityMenu, utils.AuthorityIdVerify); err != nil {
		return response.FailWithMessage(err.Error(), c)
	}
	if err := menuService.AddMenuAuthority(authorityMenu.Menus, authorityMenu.AuthorityId); err != nil {
		global.LOG.Error("添加失败!", zap.Error(err))
		return response.FailWithMessage("添加失败", c)
	} else {
		return response.OkWithMessage("添加成功", c)
	}
}

// @Tags AuthorityMenu
// @Summary 获取指定角色menu
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetAuthorityId true "角色ID"
// @Success 200 {object} response.Response{data=map[string]interface{},msg=string} "获取指定角色menu"
// @Router /menu/getMenuAuthority [get]
func (a *AuthorityMenuApi) GetMenuAuthority(c *fiber.Ctx) error {
	var param request.GetAuthorityId
	_ = c.QueryParser(&param)
	if err := utils.Verify(param, utils.AuthorityIdVerify); err != nil {
		return response.FailWithMessage(err.Error(), c)
	}
	if menus, err := menuService.GetMenuAuthority(&param); err != nil {
		global.LOG.Error("获取失败!", zap.Error(err))
		return response.FailWithDetailed(systemRes.SysMenusResponse{Menus: menus}, "获取失败", c)
	} else {
		return response.OkWithDetailed(fiber.Map{"menus": menus}, "获取成功", c)
	}
}

// @Tags Menu
// @Summary 新增菜单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysBaseMenu true "路由path, 父菜单ID, 路由name, 对应前端文件路径, 排序标记"
// @Success 200 {object} response.Response{msg=string} "新增菜单"
// @Router /menu/addBaseMenu [post]
func (a *AuthorityMenuApi) AddBaseMenu(c *fiber.Ctx) error {
	var menu system.SysBaseMenu
	_ = c.BodyParser(&menu)
	if err := utils.Verify(menu, utils.MenuVerify); err != nil {
		return response.FailWithMessage(err.Error(), c)
	}
	if err := utils.Verify(menu.Meta, utils.MenuMetaVerify); err != nil {
		return response.FailWithMessage(err.Error(), c)
	}
	if err := menuService.AddBaseMenu(menu); err != nil {
		global.LOG.Error("添加失败!", zap.Error(err))

		return response.FailWithMessage("添加失败", c)
	} else {
		return response.OkWithMessage("添加成功", c)
	}
}

// @Tags Menu
// @Summary 删除菜单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "菜单id"
// @Success 200 {object} response.Response{msg=string} "删除菜单"
// @Router /menu/deleteBaseMenu [post]
func (a *AuthorityMenuApi) DeleteBaseMenu(c *fiber.Ctx) error {
	var menu request.GetById
	menu.ID, _ = c.ParamsInt("id")
	if err := utils.Verify(menu, utils.IdVerify); err != nil {
		return response.FailWithMessage(err.Error(), c)
	}
	if err := baseMenuService.DeleteBaseMenu(menu.ID); err != nil {
		global.LOG.Error("删除失败!", zap.Error(err))
		return response.FailWithMessage("删除失败", c)
	} else {
		return response.OkWithMessage("删除成功", c)
	}
}

// @Tags Menu
// @Summary 更新菜单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysBaseMenu true "路由path, 父菜单ID, 路由name, 对应前端文件路径, 排序标记"
// @Success 200 {object} response.Response{msg=string} "更新菜单"
// @Router /menu/updateBaseMenu [post]
func (a *AuthorityMenuApi) UpdateBaseMenu(c *fiber.Ctx) error {
	var menu system.SysBaseMenu
	_ = c.BodyParser(&menu)
	if err := utils.Verify(menu, utils.MenuVerify); err != nil {
		return response.FailWithMessage(err.Error(), c)
	}
	if err := utils.Verify(menu.Meta, utils.MenuMetaVerify); err != nil {
		return response.FailWithMessage(err.Error(), c)
	}
	if err := baseMenuService.UpdateBaseMenu(menu); err != nil {
		global.LOG.Error("更新失败!", zap.Error(err))
		return response.FailWithMessage("更新失败", c)
	} else {
		return response.OkWithMessage("更新成功", c)
	}
}

// @Tags Menu
// @Summary 根据id获取菜单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "菜单id"
// @Success 200 {object} response.Response{data=systemRes.SysBaseMenuResponse,msg=string} "根据id获取菜单,返回包括系统菜单列表"
// @Router /menu/getBaseMenuById [post]
func (a *AuthorityMenuApi) GetBaseMenuById(c *fiber.Ctx) error {
	var idInfo request.GetById
	idInfo.ID, _ = c.ParamsInt("id")
	if err := utils.Verify(idInfo, utils.IdVerify); err != nil {
		return response.FailWithMessage(err.Error(), c)
	}
	if menu, err := baseMenuService.GetBaseMenuById(idInfo.ID); err != nil {
		global.LOG.Error("获取失败!", zap.Error(err))
		return response.FailWithMessage("获取失败", c)
	} else {
		return response.OkWithDetailed(systemRes.SysBaseMenuResponse{Menu: menu}, "获取成功", c)
	}
	// return response.OkWithDetailed("", "获取成功", c)
}

// @Tags Menu
// @Summary 分页获取基础menu列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "页码, 每页大小"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "分页获取基础menu列表,返回包括列表,总数,页码,每页数量"
// @Router /menu/getMenuList [post]
func (a *AuthorityMenuApi) GetMenuList(c *fiber.Ctx) error {
	var pageInfo request.PageInfo
	_ = c.QueryParser(&pageInfo)
	if err := utils.Verify(pageInfo, utils.PageInfoVerify); err != nil {
		return response.FailWithMessage(err.Error(), c)
	}
	if menuList, total, err := menuService.GetInfoList(); err != nil {
		global.LOG.Error("获取失败!", zap.Error(err))
		return response.FailWithMessage("获取失败", c)
	} else {
		return response.OkWithDetailed(response.PageResult{
			List:     menuList,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
