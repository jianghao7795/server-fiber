package middleware

import (
	"strings"

	global "server/model"
	"server/model/common/response"

	"github.com/gofiber/fiber/v3"
)

// NeedInit 检查数据库是否已初始化，未初始化则返回错误
func NeedInit(c fiber.Ctx) error {
	var tables []string
	if err := global.DB.Raw("show tables").Scan(&tables).Error; err != nil {
		return response.FailWithMessage("数据库连接失败: "+err.Error(), 3, err, c)
	}
	if strings.Contains(strings.Join(tables, ""), "sys_users") {
		return c.Next()
	}
	return response.FailWithMessage("数据库未初始化，请先初始化数据库", 3, nil, c)
}
