package app

import (
	"server-fiber/global"
	"server-fiber/model/app"
	appReq "server-fiber/model/app/request"
	"server-fiber/model/common/request"
	"server-fiber/model/common/response"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

// CreateTag 创建Tag
// @Tags Tag
// @Summary 创建Tag
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body app.Tag true "创建Tag"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appTab/createTag [post]
func (TagApi *TagApi) CreateTag(c *fiber.Ctx) error {
	var appTab app.Tag
	err := c.BodyParser(&appTab)
	if err != nil {
		global.LOG.Error("获取数据失败!", zap.Error(err))
		return response.FailWithMessage("获取数据失败", c)
	}
	if err := appTabService.CreateTag(&appTab); err != nil {
		global.LOG.Error("创建失败!", zap.Error(err))
		return response.FailWithMessage("创建失败", c)
	} else {
		return response.OkWithId("创建成功", appTab.ID, c)
	}
}

// DeleteTag 删除Tag
// @Tags Tag
// @Summary 删除Tag
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body app.Tag true "删除Tag"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /appTab/deleteTag [delete]
func (TagApi *TagApi) DeleteTag(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	if err := appTabService.DeleteTag(uint(id)); err != nil {
		global.LOG.Error("删除失败!", zap.Error(err))
		return response.FailWithMessage("删除失败", c)
	} else {
		return response.OkWithMessage("删除成功", c)
	}
}

// DeleteTagByIds 批量删除Tag
// @Tags Tag
// @Summary 批量删除Tag
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Tag"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /appTab/deleteTagByIds [delete]
func (TagApi *TagApi) DeleteTagByIds(c *fiber.Ctx) error {
	var IDS request.IdsReq
	_ = c.BodyParser(&IDS)
	if err := appTabService.DeleteTagByIds(IDS); err != nil {
		global.LOG.Error("批量删除失败!", zap.Error(err))
		return response.FailWithMessage("批量删除失败", c)
	} else {
		return response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateTag 更新Tag
// @Tags Tag
// @Summary 更新Tag
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body app.Tag true "更新Tag"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /appTab/updateTag [put]
func (TagApi *TagApi) UpdateTag(c *fiber.Ctx) error {
	var appTab app.Tag
	err := c.BodyParser(&appTab)
	if err != nil {
		global.LOG.Error("获取数据失败!", zap.Error(err))
		return response.FailWithMessage("获取数据失败", c)
	}
	if err = appTabService.UpdateTag(&appTab); err != nil {
		global.LOG.Error("更新失败!", zap.Error(err))
		return response.FailWithMessage("更新失败", c)
	} else {
		return response.OkWithMessage("更新成功", c)
	}
}

// FindTag 用id查询Tag
// @Tags Tag
// @Summary 用id查询Tag
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query app.Tag true "用id查询Tag"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /appTab/findTag [get]
func (TagApi *TagApi) FindTag(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	if tag, err := appTabService.GetTag(uint(id)); err != nil {
		global.LOG.Error("查询失败!", zap.Error(err))
		return response.FailWithMessage("查询失败", c)
	} else {
		return response.OkWithData(tag, c)
	}
}

// GetTagList 分页获取Tag列表
// @Tags Tag
// @Summary 分页获取Tag列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query appReq.TagSearch true "分页获取Tag列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appTab/getTagList [get]
func (TagApi *TagApi) GetTagList(c *fiber.Ctx) error {
	var pageInfo appReq.TagSearch
	_ = c.QueryParser(&pageInfo)
	if list, total, err := appTabService.GetTagInfoList(&pageInfo); err != nil {
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
