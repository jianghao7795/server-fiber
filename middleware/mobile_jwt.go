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
		token := strings.Replace(authHeader, "Bearer ", "", 1)
		if token == "" {
			return response.FailWithMessage401("token 失效， 请重新登录", c)
		}
		j := utils.NewJWT()
		user, err := j.ParseTokenMobile(token)
		if err != nil {
			return response.FailWithMessage401("token 失效， 请重新登录", c)
		}
		c.Locals("user_id", uint(user.ID))
		return c.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
	}
}
