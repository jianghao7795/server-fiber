package mobile

import (
	"errors"
	"strconv"

	"server-fiber/global"
	"server-fiber/model/common/response"
	"server-fiber/model/mobile"
	"server-fiber/model/mobile/request"
	mobileServer "server-fiber/service/mobile"
	"server-fiber/utils"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type MobileLoginApi struct{}

var mobileService = new(mobileServer.MobileLoginService)
var mobileRegisterService = new(mobileServer.MobileRegisterService)

func (*MobileLoginApi) Login(c *fiber.Ctx) error {
	var l mobile.Login
	_ = c.BodyParser(&l)
	if err := utils.Verify(l, utils.MobileLoginVerify); err != nil {
		return response.FailWithMessage(err.Error(), c)
	}
	token, err := mobileService.Login(l)
	if err != nil {
		global.LOG.Error("登陆失败! 用户名不存在或者密码错误!", zap.Error(err))
		return response.FailWithMessage400("用户名不存在或者密码错误", c)
	} else {
		return response.OkWithDetailed(token, "登录成功", c)
	}

}

func (*MobileLoginApi) GetUserInfo(c *fiber.Ctx) error {
	authorization := c.Get("user_id")
	if authorization == "" {
		global.LOG.Error("获取user_id失败!", zap.Error(errors.New("失败")))
		return response.FailWithMessage400("获取失败", c)
	}
	authorityId, _ := strconv.Atoi(authorization)
	if user, err := mobileService.GetUserInfo(uint(authorityId)); err != nil {
		global.LOG.Error("获取失败!", zap.Error(err))
		return response.FailWithMessage400("获取失败", c)
	} else {
		return response.OkWithDetailed(user, "获取成功", c)
	}

}

func (*MobileLoginApi) UpdateMobileUser(c *fiber.Ctx) error {
	var data request.MobileUpdate
	_ = c.BodyParser(&data)
	authorization := c.Get("user_id")

	if authorization == "" {
		return response.FailWithDetailed400(map[string]string{"id": "0"}, "更新失败", c)
	} else {
		authId, _ := strconv.Atoi(authorization)
		if err := mobileService.UpdateUser(data, (uint(authId))); err != nil {
			global.LOG.Error("更新用户失败!", zap.Error(err))
			return response.FailWithMessage("更新用户失败", c)
		} else {
			return response.OkWithDetailed(data, "获取成功", c)
		}
	}

}

func (*MobileLoginApi) UpdatePassword(c *fiber.Ctx) error {
	var data request.MobileUpdatePassword
	_ = c.BodyParser(&data)

	if err := utils.Verify(data, utils.MobileUpdatePasswordVerify); err != nil {
		return response.FailWithMessage(err.Error(), c)
	}

	if err := mobileService.UpdatePassword(data); err != nil {
		global.LOG.Error("更新密码失败!", zap.Error(err))
		return response.FailWithMessage("更新用户密码失败", c)
	} else {
		return response.OkWithDetailed(data.NewPassword, "更新成功", c)
	}
}
