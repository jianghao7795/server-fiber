package middleware

import (
	"server-fiber/global"
	"server-fiber/model/common/response"

	"github.com/gofiber/fiber/v2"
)

// 处理跨域请求,支持options访问

func NeedInit(c *fiber.Ctx) error {
	if global.DB == nil {
		return response.OkWithDetailed(fiber.Map{
			"needInit": true,
		}, "前往初始化数据库", c)

	}
	return c.Next()

	// 处理请求
}
