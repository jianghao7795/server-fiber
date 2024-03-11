package system

import (
	"server-fiber/global"
	"server-fiber/model/common/response"
	problemReq "server-fiber/model/system"
	"server-fiber/service"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type UserProblem struct{}

var userProblem = service.ServiceGroupApp.SystemServiceGroup.Problem

func (*UserProblem) GetProblemSetting(c *fiber.Ctx) error {
	var search problemReq.SysUserProblem
	SysUserId := c.Params("id")
	SysUserProblemId, err := strconv.Atoi(SysUserId)
	search.SysUserId = SysUserProblemId
	if err != nil {
		global.LOG.Error("传参错误!", zap.Error(err))
		return response.FailWithMessage("传错参数,请传user id", c)
	}
	list, err := userProblem.GetUserProblemSettingList(&search)
	if err != nil {
		global.LOG.Error("获取失败!", zap.Error(err))
		return response.FailWithMessage("获取失败", c)
	}
	return response.OkWithDetailed(response.PageResult{
		List: list,
	}, "获取成功", c)
}

type UpdateProblemSettingData struct {
	Data []problemReq.SysUserProblem `json:"data"`
}

func (*UserProblem) UpdateProblemSetting(c *fiber.Ctx) error {
	var dataProblem UpdateProblemSettingData
	err := c.QueryParser(&dataProblem)
	if err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			global.LOG.Error("传参错误!", zap.Error(err))
			return response.FailWithDetailed(err.Error(), "传错参数", c)
		}
		global.LOG.Error("传参错误!", zap.Error(err))

		return response.FailWithDetailed(errs.Translate(global.Validate), "传错参数", c)
	}
	message, err := userProblem.SetUserProblemSetting(dataProblem.Data)
	if err != nil {
		global.LOG.Error("传参错误!", zap.Error(err))
		return response.FailWithDetailed(err, "传错参数", c)
	}
	return response.OkWithMessage(message, c)
}

func (*UserProblem) HasSetting(c *fiber.Ctx) error {
	SysUserId := c.Params("uid")
	SysUserProblemId, _ := strconv.Atoi(SysUserId)
	isSetting, err := userProblem.HasSetting(SysUserProblemId)
	if err != nil {
		global.LOG.Error("传参错误!", zap.Error(err))
		return response.FailWithDetailed(err, "传错参数", c)
	}
	return response.OkWithDetailed(isSetting, "获取成功", c)
}

type VerifyProblemSettingData struct {
	Data problemReq.SysUserProblem `json:"data"`
}

func (*UserProblem) VerifyAnswer(c *fiber.Ctx) error {
	var dataProblem VerifyProblemSettingData
	err := c.QueryParser(&dataProblem)
	if err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			global.LOG.Error("传参错误!", zap.Error(err))
			return response.FailWithDetailed(err.Error(), "传错参数", c)
		}
		global.LOG.Error("传参错误!", zap.Error(err))
		return response.FailWithDetailed(errs.Translate(global.Validate), "传错参数", c)
	}
	ispassed, err := userProblem.VerifyAnswer(&dataProblem.Data)
	if err != nil {
		global.LOG.Error("未查到此问题!", zap.Error(err))
		return response.FailWithDetailed(err, "未查到此问题", c)
	}
	return response.OkWithDetailed(ispassed, "已验证", c)
}
