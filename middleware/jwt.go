package middleware

import (
	"strings"

	"server-fiber/utils"

	"server-fiber/model/common/response"
	systemService "server-fiber/service/system"

	"github.com/gofiber/fiber/v2"
)

var jwtService = new(systemService.JwtService)

func JWTAuth(c *fiber.Ctx) error {
	// 解决访问文件的401问题
	if strings.Contains(c.Path(), "uploads/excel/") || strings.Contains(c.Path(), "uploads/file/") {
		code := c.Response().StatusCode()
		return c.Status(code).SendFile(strings.Join(strings.Split(c.Path(), "/")[2:], "/"))
	}
	tokenString := c.Get("Authorization")

	if tokenString == "" {
		return response.FailWithDetailed401(fiber.Map{"reload": true}, "未登录或非法访问", c)
	}
	token := strings.Replace(tokenString, "Bearer ", "", 1)
	if token == "" {
		return response.FailWithDetailed401(fiber.Map{"reload": true}, "未登录或非法访问", c)
	}
	if jwtService.IsBlacklist(token) {
		return response.FailWithDetailed401(fiber.Map{"reload": true}, "您的帐户异地登陆或令牌失效", c)
	}
	j := utils.NewJWT()
	// parseToken 解析token包含的信息
	claims, err := j.ParseToken(token)
	if err != nil {
		if err == utils.ErrTokenExpired {
			return response.FailWithDetailed401(fiber.Map{"reload": true}, "授权已过期", c)
		}
		return response.FailWithDetailed401(fiber.Map{"reload": true}, err.Error(), c)
	}
	// 继续交由下一个路由处理,并将解析出的信息传递下去
	c.Locals("claims", claims)
	return c.Next()
}
