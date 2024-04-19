package system

import (
	"server-fiber/global"
	"server-fiber/model/common/response"
	"server-fiber/model/system/request"

	"go.uber.org/zap"

	"github.com/gofiber/fiber/v2"
)

type DBApi struct{}

// InitDB
// @Tags InitDB
// @Summary 初始化用户数据库
// @Produce  application/json
// @Param data body request.InitDB true "初始化数据库参数"
// @Success 200 {object} response.Response{data=string} "初始化用户数据库"
// @Router /init/initdb [post]
func (i *DBApi) InitDB(c *fiber.Ctx) error {
	if global.DB != nil {
		global.LOG.Error("已存在数据库配置!")
		return response.FailWithMessage("已存在数据库配置", c)
	}
	var dbInfo request.InitDB
	if err := c.QueryParser(&dbInfo); err != nil {
		global.LOG.Error("参数校验不通过!", zap.Error(err))
		return response.FailWithMessage("参数校验不通过", c)
	}
	if err := initDBService.InitDB(dbInfo); err != nil {
		global.LOG.Error("自动创建数据库失败!", zap.Error(err))
		return response.FailWithMessage("自动创建数据库失败，请查看后台日志，检查后在进行初始化", c)
	}
	return response.OkWithData("自动创建数据库成功", c)
}

// CheckDB
// @Tags CheckDB
// @Summary 初始化用户数据库
// @Produce  application/json
// @Success 200 {object} response.Response{data=map[string]interface{},msg=string} "初始化用户数据库"
// @Router /init/checkdb [post]
func (i *DBApi) CheckDB(c *fiber.Ctx) error {
	var (
		message  = "前往初始化数据库"
		needInit = true
	)

	if global.DB != nil {
		message = "数据库无需初始化"
		needInit = false
	}
	global.LOG.Info(message)
	return response.OkWithDetailed(fiber.Map{"needInit": needInit}, message, c)
}
