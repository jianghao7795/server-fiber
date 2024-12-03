package system

import (
	global "server-fiber/model"
	"server-fiber/model/common/response"
	"server-fiber/model/system/request"
	"server-fiber/utils"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type CasbinApi struct{}

// @Tags Casbin
// @Summary 更新角色api权限
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.CasbinInReceive true "权限id, 权限模型列表"
// @Success 200 {object} response.Response{msg=string} "更新角色api权限"
// @Router /casbin/UpdateCasbin [post]
func (cas *CasbinApi) UpdateCasbin(c *fiber.Ctx) error {
	var cmr request.CasbinInReceive
	if err := c.BodyParser(&cmr); err != nil {
		global.LOG.Error("获取数据失败!", zap.Error(err))
		return response.FailWithMessage("获取数据失败", c)
	}
	if err := utils.Verify(cmr, utils.AuthorityIdVerify); err != nil {
		return response.FailWithMessage(err.Error(), c)
	}
	if err := casbinService.UpdateCasbin(cmr.AuthorityId, cmr.CasbinInfos); err != nil {
		global.LOG.Error("更新失败!", zap.Error(err))
		return response.FailWithMessage("更新失败", c)
	} else {
		return response.OkWithMessage("更新成功", c)
	}
}

// @Tags Casbin
// @Summary 获取权限列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.CasbinInReceive true "权限id, 权限模型列表"
// @Success 200 {object} response.Response{data=string,msg=string} "获取权限列表,返回包括casbin详情列表"
// @Router /casbin/getPolicyPathByAuthorityId [get]
func (cas *CasbinApi) GetPolicyPathByAuthorityId(c *fiber.Ctx) error {
	var casbin request.CasbinInReceive
	_ = c.QueryParser(&casbin)
	casbin.AuthorityId = c.Params("id")
	if err := utils.Verify(casbin, utils.AuthorityIdVerify); err != nil {
		return response.FailWithMessage(err.Error(), c)
	}
	paths := casbinService.GetPolicyPathByAuthorityId(casbin.AuthorityId)
	return response.OkWithDetailed(paths, "获取成功", c)
}
