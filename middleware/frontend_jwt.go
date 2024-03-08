package middleware

import (
	"server-fiber/model/common/response"
	"server-fiber/service/frontend"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func JWTAuthMiddleware(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return response.FailWithMessage("token 失效", c)

	}

	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		response.FailWithMessage("token 不正确", c)

	}

	_, err := frontend.ParseToken(parts[1])
	if err != nil {
		return response.FailWithMessage("token 失效， 请重新登录", c)
	}

	return c.Next() //
}
