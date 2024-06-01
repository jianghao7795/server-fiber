package response

import (
	"github.com/gofiber/fiber/v2"
)

type ResponseMobile struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

const (
	ERROR401 = fiber.StatusUnauthorized
	ERROR403 = fiber.StatusForbidden
)

// 返回400 错误信息
func FailWithDetailed401(data interface{}, message string, c *fiber.Ctx) error {
	return Result400(ERROR401, data, message, c)
}

func FailWithMessage401(message string, c *fiber.Ctx) error {
	return Result400(ERROR401, map[string]interface{}{}, message, c)
}

func FailWithMessage403(message string, c *fiber.Ctx) error {
	return Result400(ERROR403, map[string]interface{}{}, message, c)
}
