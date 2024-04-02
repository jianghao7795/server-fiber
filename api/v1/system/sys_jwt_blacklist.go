package system

import (
	"server-fiber/global"
	"server-fiber/model/common/response"
	"server-fiber/model/system"
	"strings"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type JwtApi struct{}

// @Tags Jwt
// @Summary jwt加入黑名单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} return response.Response{msg=string} "jwt加入黑名单"
// @Router /jwt/jsonInBlacklist [post]
func (j *JwtApi) JsonInBlacklist(c *fiber.Ctx) error {
	tokenString := c.Get("Authorization")
	token := strings.Replace(tokenString, "Bearer ", "", 1)
	if token == "" {
		return response.FailWithMessage401("token 失效， 请重新登录", c)
	}
	jwt := system.JwtBlacklist{Jwt: token}
	if err := jwtService.JsonInBlacklist(jwt); err != nil {
		global.LOG.Error("jwt作废失败!", zap.Error(err))
		return response.FailWithMessage("jwt作废失败", c)
	} else {
		return response.OkWithMessage("jwt作废成功", c)
	}
}
