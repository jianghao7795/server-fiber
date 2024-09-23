package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Update_updator(c *fiber.Ctx) error {
	fmt.Println("after write user")
	return nil
}
