package middleware

import (
	"server-fiber/model/common/response"
	"server-fiber/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func JWTAuthMobileMiddleware() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return response.FailWithMessage401("token 失效", c)
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			return response.FailWithMessage401("token 不正确", c)
		}
		j := utils.NewJWT()
		user, err := j.ParseTokenMobile(parts[1])
		if err != nil {
			return response.FailWithMessage401("token 失效， 请重新登录", c)
		}
		c.Locals("user_id", uint(user.ID))
		return c.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
	}
}
