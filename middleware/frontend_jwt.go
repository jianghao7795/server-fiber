package middleware

import (
	"server-fiber/model/common/response"
	"server-fiber/service/frontend"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func JWTAuthMiddleware(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	token := strings.Replace(authHeader, "Bearer", "", 1)
	if token == "" {
		return response.FailWithMessage401("token 失效， 请重新登录", c)
	}
	_, err := frontend.ParseToken(token)
	if err != nil {
		return response.FailWithMessage("token 失效， 请重新登录", c)
	}

	return c.Next() //
}
