package system

import (
	"server-fiber/global"
	"server-fiber/model/common/response"
	"server-fiber/model/system"
	"server-fiber/model/system/request"
	"server-fiber/utils"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type DictionaryApi struct{}

// @Tags SysDictionary
// @Summary 创建SysDictionary
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysDictionary true "SysDictionary模型"
// @Success 200 {object} response.Response{msg=string} "创建SysDictionary"
// @Router /sysDictionary/createSysDictionary [post]
func (s *DictionaryApi) CreateSysDictionary(c *fiber.Ctx) error {
	var dictionary system.SysDictionary
	_ = c.BodyParser(&dictionary)
	if err := dictionaryService.CreateSysDictionary(dictionary); err != nil {
		global.LOG.Error("创建失败!", zap.Error(err))
		return response.FailWithMessage("创建失败", c)
	} else {
		return response.OkWithMessage("创建成功", c)
	}
}

// @Tags SysDictionary
// @Summary 删除SysDictionary
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysDictionary true "SysDictionary模型"
// @Success 200 {object} response.Response{msg=string} "删除SysDictionary"
// @Router /sysDictionary/deleteSysDictionary/:id [delete]
func (s *DictionaryApi) DeleteSysDictionary(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	if id == 0 {
		global.LOG.Error("id获取失败!")
		return response.FailWithMessage("id获取失败", c)
	}
	if err := dictionaryService.DeleteSysDictionary(uint(id)); err != nil {
		global.LOG.Error("删除失败!", zap.Error(err))
		return response.FailWithMessage("删除失败", c)
	} else {
		return response.OkWithMessage("删除成功", c)
	}
}

// @Tags SysDictionary
// @Summary 更新SysDictionary
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysDictionary true "SysDictionary模型"
// @Success 200 {object} response.Response{msg=string} "更新SysDictionary"
// @Router /sysDictionary/updateSysDictionary/:id [put]
func (s *DictionaryApi) UpdateSysDictionary(c *fiber.Ctx) error {
	var dictionary system.SysDictionary
	_ = c.BodyParser(&dictionary)
	id, _ := c.ParamsInt("id")
	if dictionary.ID != uint(id) {
		return response.FailWithMessage("id不一致", c)
	}
	if err := dictionaryService.UpdateSysDictionary(&dictionary); err != nil {
		global.LOG.Error("更新失败!", zap.Error(err))
		return response.FailWithMessage("更新失败", c)
	} else {
		return response.OkWithMessage("更新成功", c)
	}
}

// @Tags SysDictionary
// @Summary 用id查询SysDictionary
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query system.SysDictionary true "ID或字典英名"
// @Success 200 {object} response.Response{data=map[string]interface{},msg=string} "用id查询SysDictionary"
// @Router /sysDictionary/findSysDictionary/:id [get]
func (s *DictionaryApi) FindSysDictionary(c *fiber.Ctx) error {
	fieldType := c.Params("type")
	if fieldType == "" {
		global.LOG.Error("field 获取失败!")
		return response.FailWithMessage("field 获取失败", c)
	}
	if sysDictionary, err := dictionaryService.GetSysDictionary(fieldType); err != nil {
		global.LOG.Error("查询失败!", zap.Error(err))
		return response.FailWithMessage("查询失败", c)
	} else {
		return response.OkWithDetailed(sysDictionary, "查询成功", c)
	}
}

// @Tags SysDictionary
// @Summary 分页获取SysDictionary列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.SysDictionarySearch true "页码, 每页大小, 搜索条件"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "分页获取SysDictionary列表,返回包括列表,总数,页码,每页数量"
// @Router /sysDictionary/getSysDictionaryList [get]
func (s *DictionaryApi) GetSysDictionaryList(c *fiber.Ctx) error {
	var pageInfo request.SysDictionarySearch
	_ = c.QueryParser(&pageInfo)
	if err := utils.Verify(pageInfo.PageInfo, utils.PageInfoVerify); err != nil {
		return response.FailWithMessage(err.Error(), c)
	}
	if list, total, err := dictionaryService.GetSysDictionaryInfoList(pageInfo); err != nil {
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
