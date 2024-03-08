package mobile

import (
	"server-fiber/global"
	"server-fiber/model/common/request"
	"server-fiber/model/common/response"
	"server-fiber/model/mobile"
	mobileReq "server-fiber/model/mobile/request"
	"server-fiber/service"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type MobileUserApi struct {
}

var moblieUserService = service.ServiceGroupApp.MobileServiceGroup.MobileUserService

// CreateMoblieUser 创建MoblieUser
// @Tags MoblieUser
// @Summary 创建MoblieUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body mobile.MoblieUser true "创建MoblieUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /moblieUser/createMoblieUser [post]
func (mobileUserApi *MobileUserApi) CreateMoblieUser(c *fiber.Ctx) error {
	var moblieUser mobile.MobileUser
	_ = c.QueryParser(&moblieUser)
	if err := moblieUserService.CreateMoblieUser(moblieUser); err != nil {
		global.LOG.Error("创建失败!", zap.Error(err))
		return response.FailWithMessage("创建失败", c)
	} else {
		return response.OkWithMessage("创建成功", c)
	}
}

// DeleteMoblieUser 删除MoblieUser
// @Tags MoblieUser
// @Summary 删除MoblieUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body mobile.MoblieUser true "删除MoblieUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /moblieUser/deleteMoblieUser [delete]
func (mobileUserApi *MobileUserApi) DeleteMoblieUser(c *fiber.Ctx) error {
	idString := c.Params("id")
	id, _ := strconv.Atoi(idString)
	if err := moblieUserService.DeleteMoblieUser(uint(id)); err != nil {
		global.LOG.Error("删除失败!", zap.Error(err))
		return response.FailWithMessage("删除失败", c)
	} else {
		return response.OkWithMessage("删除成功", c)
	}
}

// DeleteMoblieUserByIds 批量删除MoblieUser
// @Tags MoblieUser
// @Summary 批量删除MoblieUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除MoblieUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /moblieUser/deleteMoblieUserByIds [delete]
func (mobileUserApi *MobileUserApi) DeleteMoblieUserByIds(c *fiber.Ctx) error {
	var IDS request.IdsReq
	_ = c.QueryParser(&IDS)
	if err := moblieUserService.DeleteMoblieUserByIds(IDS); err != nil {
		global.LOG.Error("批量删除失败!", zap.Error(err))
		return response.FailWithMessage("批量删除失败", c)
	} else {
		return response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateMoblieUser 更新MoblieUser
// @Tags MoblieUser
// @Summary 更新MoblieUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body mobile.MoblieUser true "更新MoblieUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /moblieUser/updateMoblieUser [put]
func (mobileUserApi *MobileUserApi) UpdateMoblieUser(c *fiber.Ctx) error {
	var moblieUser mobile.MobileUser
	_ = c.QueryParser(&moblieUser)
	if err := moblieUserService.UpdateMoblieUser(moblieUser); err != nil {
		global.LOG.Error("更新失败!", zap.Error(err))
		return response.FailWithMessage("更新失败", c)
	} else {
		return response.OkWithMessage("更新成功", c)
	}
}

// FindMoblieUser 用id查询MoblieUser
// @Tags MoblieUser
// @Summary 用id查询MoblieUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query mobile.MoblieUser true "用id查询MoblieUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /moblieUser/findMoblieUser [get]
func (mobileUserApi *MobileUserApi) FindMoblieUser(c *fiber.Ctx) error {
	idString := c.Params("id")
	id, _ := strconv.Atoi(idString)
	if remoblieUser, err := moblieUserService.GetMoblieUser(uint(id)); err != nil {
		global.LOG.Error("查询失败!", zap.Error(err))
		return response.FailWithMessage("查询失败", c)
	} else {
		return response.OkWithData(gin.H{"remoblieUser": remoblieUser}, c)
	}
}

// GetMoblieUserList 分页获取MoblieUser列表
// @Tags MoblieUser
// @Summary 分页获取MoblieUser列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query mobileReq.MoblieUserSearch true "分页获取MoblieUser列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /moblieUser/getMoblieUserList [get]
func (mobileUserApi *MobileUserApi) GetMoblieUserList(c *fiber.Ctx) error {
	var pageInfo mobileReq.MoblieUserSearch
	_ = c.QueryParser(&pageInfo)
	if list, total, err := moblieUserService.GetMoblieUserInfoList(pageInfo); err != nil {
		global.LOG.Error("获取失败!", zap.Error(err))
		return response.FailWithMessage("获取失败", c)
	} else {
		return response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
