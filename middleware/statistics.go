package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func Statistics(ctx *fiber.Ctx) error {
	return ctx.Next()
}
