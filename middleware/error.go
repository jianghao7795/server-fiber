package middleware

import (
	"github.com/gofiber/fiber/v2"
)

// GinRecovery recover掉项目可能出现的panic，并使用zap记录相关日志
func GinRecovery(c *fiber.Ctx) error {
	return c.Next()
}
