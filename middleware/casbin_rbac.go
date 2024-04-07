package middleware

import (
	"server-fiber/model/common/response"
	service "server-fiber/service/system"
	"server-fiber/utils"

	"github.com/gofiber/fiber/v2"
)

var casbinService = new(service.CasbinService)

// 拦截器
func CasbinHandler(c *fiber.Ctx) error {
	waitUse, _ := utils.GetClaims(c)
	// 获取请求的PATH
	obj := c.Path()
	// 获取请求方法
	act := c.Method()
	// 获取用户的角色
	sub := waitUse.AuthorityId
	e := casbinService.Casbin()
	// 判断策略中是否存在
	success, _ := e.Enforce(sub, obj, act)
	// if global.CONFIG.System.Env == "develop" || success {
	if success {
		return c.Next()
	} else {
		// 上传文件 由于是ajxs 必须返回400 错误 才能展示错误信息
		if obj == "/base_message/upload_file" {
			return response.FailWithDetailed400(nil, "权限不足", c)

		}
		return response.FailWithDetailed(nil, "权限不足", c)
	}
}
