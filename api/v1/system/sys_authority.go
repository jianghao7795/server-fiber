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

type AuthorityApi struct{}

// @Tags Authority
// @Summary 创建角色
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysAuthority true "权限id, 权限名, 父角色id"
// @Success 200 {object} response.Response{data=systemRes.SysAuthorityResponse,msg=string} "创建角色,返回包括系统角色详情"
// @Router /authority/createAuthority [post]
func (a *AuthorityApi) CreateAuthority(c *fiber.Ctx) error {
	var authority system.SysAuthority
	_ = c.BodyParser(&authority)
	if err := utils.Verify(authority, utils.AuthorityVerify); err != nil {
		return response.FailWithMessage(err.Error(), c)
	}
	if authBack, err := authorityService.CreateAuthority(authority); err != nil {
		global.LOG.Error("创建失败!", zap.Error(err))
		return response.FailWithMessage("创建失败"+err.Error(), c)
	} else {
		_ = menuService.AddMenuAuthority(systemReq.DefaultMenu(), authority.AuthorityId)
		_ = casbinService.UpdateCasbin(authority.AuthorityId, systemReq.DefaultCasbin())
		return response.OkWithDetailed(systemRes.SysAuthorityResponse{Authority: authBack}, "创建成功", c)
	}
}

// @Tags Authority
// @Summary 拷贝角色
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body systemRes.SysAuthorityCopyResponse true "旧角色id, 新权限id, 新权限名, 新父角色id"
// @Success 200 {object} response.Response{data=systemRes.SysAuthorityResponse,msg=string} "拷贝角色,返回包括系统角色详情"
// @Router /authority/copyAuthority [post]
func (a *AuthorityApi) CopyAuthority(c *fiber.Ctx) error {
	var copyInfo systemRes.SysAuthorityCopyResponse
	err := c.BodyParser(&copyInfo)
	if err != nil {
		global.LOG.Error("获取数据错误", zap.Error(err))
		return response.FailWithMessage("获取数据错误", c)
	}
	if err := utils.Verify(copyInfo, utils.OldAuthorityVerify); err != nil {
		return response.FailWithMessage(err.Error(), c)
	}
	if err := utils.Verify(copyInfo.Authority, utils.AuthorityVerify); err != nil {
		return response.FailWithMessage(err.Error(), c)
	}
	if authBack, err := authorityService.CopyAuthority(copyInfo); err != nil {
		global.LOG.Error("拷贝失败!", zap.Error(err))
		return response.FailWithMessage("拷贝失败"+err.Error(), c)
	} else {
		return response.OkWithDetailed(systemRes.SysAuthorityResponse{Authority: authBack}, "拷贝成功", c)
	}
}

// @Tags Authority
// @Summary 删除角色
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysAuthority true "删除角色"
// @Success 200 {object} response.Response{msg=string} "删除角色"
// @Router /authority/deleteAuthority [delete]
func (a *AuthorityApi) DeleteAuthority(c *fiber.Ctx) error {
	var authority system.SysAuthority
	_ = c.QueryParser(&authority)
	if err := utils.Verify(authority, utils.AuthorityIdVerify); err != nil {
		return response.FailWithMessage(err.Error(), c)
	}
	if err := authorityService.DeleteAuthority(&authority); err != nil { // 删除角色之前需要判断是否有用户正在使用此角色
		global.LOG.Error("删除失败!", zap.Error(err))
		return response.FailWithMessage("删除失败"+err.Error(), c)
	} else {
		return response.OkWithMessage("删除成功", c)
	}
}

// @Tags Authority
// @Summary 更新角色信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysAuthority true "权限id, 权限名, 父角色id"
// @Success 200 {object} response.Response{data=systemRes.SysAuthorityResponse,msg=string} "更新角色信息,返回包括系统角色详情"
// @Router /authority/updateAuthority [post]
func (a *AuthorityApi) UpdateAuthority(c *fiber.Ctx) error {
	var auth system.SysAuthority
	if err := c.BodyParser(&auth); err != nil {
		global.LOG.Error("获取数据失败", zap.Error(err))
		return response.FailWithMessage("获取数据失败", c)
	}

	if err := utils.Verify(auth, utils.AuthorityVerify); err != nil {
		return response.FailWithMessage(err.Error(), c)
	}
	if authority, err := authorityService.UpdateAuthority(auth); err != nil {
		global.LOG.Error("更新失败!", zap.Error(err))
		return response.FailWithMessage("更新失败"+err.Error(), c)
	} else {
		return response.OkWithDetailed(systemRes.SysAuthorityResponse{Authority: authority}, "更新成功", c)
	}
}

// @Tags Authority
// @Summary 分页获取角色列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "页码, 每页大小"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "分页获取角色列表,返回包括列表,总数,页码,每页数量"
// @Router /authority/getAuthorityList [post]
func (a *AuthorityApi) GetAuthorityList(c *fiber.Ctx) error {
	var pageInfo request.PageInfo
	_ = c.QueryParser(&pageInfo)
	if pageInfo.Page == 0 {
		pageInfo.Page = 1
	}
	if pageInfo.PageSize == 0 {
		pageInfo.PageSize = 10
	}
	if err := utils.Verify(pageInfo, utils.PageInfoVerify); err != nil {
		return response.FailWithMessage(err.Error(), c)
	}
	if list, total, err := authorityService.GetAuthorityInfoList(pageInfo); err != nil {
		global.LOG.Error("获取失败!", zap.Error(err))
		return response.FailWithMessage("获取失败"+err.Error(), c)
	} else {
		return response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// @Tags Authority
// @Summary 设置角色资源权限
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysAuthority true "设置角色资源权限"
// @Success 200 {object} response.Response{msg=string} "设置角色资源权限"
// @Router /authority/setDataAuthority [post]
func (a *AuthorityApi) SetDataAuthority(c *fiber.Ctx) error {
	var auth system.SysAuthority
	if err := c.BodyParser(&auth); err != nil {
		global.LOG.Error("获取数据失败!", zap.Error(err))
		return response.FailWithMessage("获取数据失败", c)
	}
	if err := utils.Verify(auth, utils.AuthorityIdVerify); err != nil {
		global.LOG.Error("验证角色权限错误", zap.Error(err))
		return response.FailWithMessage(err.Error(), c)
	}
	if err := authorityService.SetDataAuthority(&auth); err != nil {
		global.LOG.Error("设置失败!", zap.Error(err))
		return response.FailWithMessage("设置失败"+err.Error(), c)
	} else {
		return response.OkWithMessage("设置成功", c)
	}
}
