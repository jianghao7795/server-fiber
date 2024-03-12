package middleware

import (
	"github.com/gofiber/fiber/v2"
)

// 处理跨域请求,支持options访问

func NeedInit(c *fiber.Ctx) error {
	return c.Next()
	// 处理请求
}
