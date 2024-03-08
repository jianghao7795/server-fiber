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
	tokenValue := strings.Split(tokenString, " ")
	token := tokenValue[1]
	jwt := system.JwtBlacklist{Jwt: token}
	if err := jwtService.JsonInBlacklist(jwt); err != nil {
		global.LOG.Error("jwt作废失败!", zap.Error(err))
		return response.FailWithMessage("jwt作废失败", c)
	} else {
		return response.OkWithMessage("jwt作废成功", c)
	}
}
