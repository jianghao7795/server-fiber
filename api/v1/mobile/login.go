package mobile

import (
	"errors"
	"strconv"

	"server-fiber/global"
	"server-fiber/model/common/response"
	"server-fiber/model/mobile"
	"server-fiber/model/mobile/request"
	"server-fiber/utils"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type LoginApi struct{}

func (*LoginApi) Login(c *fiber.Ctx) error {
	var l mobile.Login
	_ = c.BodyParser(&l)
	if err := utils.Verify(l, utils.MobileLoginVerify); err != nil {
		return response.FailWithMessage(err.Error(), c)
	}
	token, err := loginService.Login(l)
	if err != nil {
		global.LOG.Error("登陆失败! 用户名不存在或者密码错误!", zap.Error(err))
		return response.FailWithMessage400("用户名不存在或者密码错误", c)
	} else {
		return response.OkWithDetailed(token, "登录成功", c)
	}

}

func (*LoginApi) GetUserInfo(c *fiber.Ctx) error {
	authorization := c.Get("user_id")
	if authorization == "" {
		global.LOG.Error("获取user_id失败!", zap.Error(errors.New("失败")))
		return response.FailWithMessage400("获取失败", c)
	}
	authorityId, _ := strconv.Atoi(authorization)
	if user, err := loginService.GetUserInfo(uint(authorityId)); err != nil {
		global.LOG.Error("获取失败!", zap.Error(err))
		return response.FailWithMessage400("获取失败", c)
	} else {
		return response.OkWithDetailed(user, "获取成功", c)
	}

}

func (*LoginApi) UpdateMobileUser(c *fiber.Ctx) error {
	var data request.MobileUpdate
	if err := c.BodyParser(&data); err != nil {
		global.LOG.Error("获取数据失败", zap.Error(err))
		return response.FailWithMessage("获取数据失败", c)
	}
	authorization := c.Get("user_id") // user_id 在请求头信息中

	if authorization == "" {
		return response.FailWithDetailed400(map[string]string{"id": "0"}, "更新失败", c)
	} else {
		authId, _ := strconv.Atoi(authorization)
		if err := loginService.UpdateUser(&data, uint(authId)); err != nil {
			global.LOG.Error("更新用户失败!", zap.Error(err))
			return response.FailWithMessage("更新用户失败", c)
		} else {
			return response.OkWithDetailed(data, "更新成功", c)
		}
	}

}

func (*LoginApi) UpdatePassword(c *fiber.Ctx) error {
	var data request.MobileUpdatePassword
	if err := c.BodyParser(&data); err != nil {
		global.LOG.Error("获取数据失败", zap.Error(err))
		return response.FailWithMessage("获取数据失败", c)
	}

	if err := utils.Verify(data, utils.MobileUpdatePasswordVerify); err != nil {
		return response.FailWithMessage(err.Error(), c)
	}

	if err := loginService.UpdatePassword(data); err != nil {
		global.LOG.Error("更新密码失败!", zap.Error(err))
		return response.FailWithMessage("更新用户密码失败", c)
	} else {
		return response.OkWithDetailed(data.NewPassword, "更新成功", c)
	}
}
