package system

import (
	"errors"
	"fmt"
	"net/url"
	"os"

	"server-fiber/global"
	"server-fiber/model/common/response"
	"server-fiber/model/system"
	"server-fiber/utils"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type AutoCodeApi struct{}

// PreviewTemp
// @Tags AutoCode
// @Summary 预览创建后的代码
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.AutoCodeStruct true "预览创建代码"
// @Success 200 {object} return response.Response{data=map[string]interface{},msg=string} "预览创建后的代码"
// @Router /autoCode/preview [post]
func (autoApi *AutoCodeApi) PreviewTemp(c *fiber.Ctx) error {
	var a system.AutoCodeStruct
	_ = c.BodyParser(&a)
	if err := utils.Verify(a, utils.AutoCodeVerify); err != nil {
		return response.FailWithMessage(err.Error(), c)
	}
	a.PackageT = cases.Title(language.Dutch).String(a.Package)
	autoCode, err := autoCodeService.PreviewTemp(a)
	if err != nil {
		global.LOG.Error("预览失败!", zap.Error(err))
		return response.FailWithMessage("预览失败", c)
	} else {
		return response.OkWithDetailed(fiber.Map{"autoCode": autoCode}, "预览成功", c)
	}
}

// CreateTemp
// @Tags AutoCode
// @Summary 自动代码模板
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.AutoCodeStruct true "创建自动代码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /autoCode/createTemp [post]
func (autoApi *AutoCodeApi) CreateTemp(c *fiber.Ctx) error {
	var a system.AutoCodeStruct
	_ = c.BodyParser(&a)
	if err := utils.Verify(a, utils.AutoCodeVerify); err != nil {
		return response.FailWithMessage(err.Error(), c)
	}
	var apiIds []uint
	if a.AutoCreateApiToSql {
		if ids, err := autoCodeService.AutoCreateApi(&a); err != nil {
			global.LOG.Error("自动化创建失败!请自行清空垃圾数据!", zap.Error(err))
			c.Set("success", "false")
			c.Set("msg", url.QueryEscape("自动化创建失败!请自行清空垃圾数据!"))
		} else {
			apiIds = ids
		}
	}
	a.PackageT = cases.Title(language.Dutch).String(a.Package)
	err := autoCodeService.CreateTemp(a, apiIds...)
	if err != nil {
		if errors.Is(err, system.ErrAutoMove) {
			c.Locals("success", "true")
			c.Locals("msg", url.QueryEscape(err.Error()))
		} else {
			c.Locals("success", "false")
			c.Locals("msg", url.QueryEscape(err.Error()))
			_ = os.Remove("./ginvueadmin.zip")
		}
	} else {
		c.Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", "ginvueadmin.zip")) // fmt.Sprintf("attachment; filename=%s", filename)对下载的文件重命名
		c.Set("Content-Type", "application/json")
		c.Locals("success", "true")
		c.Download("./ginvueadmin.zip")
		_ = os.Remove("./ginvueadmin.zip")
	}
	return nil
}

// GetDB
// @Tags AutoCode
// @Summary 获取当前所有数据库
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} return response.Response{data=map[string]interface{},msg=string} "获取当前所有数据库"
// @Router /autoCode/getDatabase [get]
func (autoApi *AutoCodeApi) GetDB(c *fiber.Ctx) error {
	dbs, err := autoCodeService.Database().GetDB()
	if err != nil {
		global.LOG.Error("获取失败!", zap.Error(err))
		return response.FailWithMessage("获取失败", c)
	} else {
		return response.OkWithDetailed(fiber.Map{"dbs": dbs}, "获取成功", c)
	}
}

// GetTables
// @Tags AutoCode
// @Summary 获取当前数据库所有表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} return response.Response{data=map[string]interface{},msg=string} "获取当前数据库所有表"
// @Router /autoCode/getTables [get]
func (autoApi *AutoCodeApi) GetTables(c *fiber.Ctx) error {
	dbName := c.Query("dbName", global.CONFIG.Mysql.Dbname)
	tables, err := autoCodeService.Database().GetTables(dbName)
	if err != nil {
		global.LOG.Error("查询table失败!", zap.Error(err))
		return response.FailWithMessage("查询table失败", c)
	} else {
		return response.OkWithDetailed(fiber.Map{"tables": tables}, "获取成功", c)
	}
}

// GetColumn
// @Tags AutoCode
// @Summary 获取当前表所有字段
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} return response.Response{data=map[string]interface{},msg=string} "获取当前表所有字段"
// @Router /autoCode/getColumn [get]
func (autoApi *AutoCodeApi) GetColumn(c *fiber.Ctx) error {
	dbName := c.Query("dbName", global.CONFIG.Mysql.Dbname)
	tableName := c.Query("tableName")
	columns, err := autoCodeService.Database().GetColumn(tableName, dbName)
	if err != nil {
		global.LOG.Error("获取失败!", zap.Error(err))
		return response.FailWithMessage("获取失败", c)
	} else {
		return response.OkWithDetailed(fiber.Map{"columns": columns}, "获取成功", c)
	}
}

// CreatePackage
// @Tags AutoCode
// @Summary 创建package
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysAutoCode true "创建package"
// @Success 200 {object} return response.Response{data=map[string]interface{},msg=string} "创建package成功"
// @Router /autoCode/createPackage [post]
func (autoApi *AutoCodeApi) CreatePackage(c *fiber.Ctx) error {
	var a system.SysAutoCode
	_ = c.QueryParser(&a)
	if err := utils.Verify(a, utils.AutoPackageVerify); err != nil {
		return response.FailWithMessage(err.Error(), c)
	}
	err := autoCodeService.CreateAutoCode(&a)
	if err != nil {
		global.LOG.Error("创建成功!", zap.Error(err))
		return response.FailWithMessage("创建失败", c)
	} else {
		return response.OkWithMessage("创建成功", c)
	}
}

// GetPackage
// @Tags AutoCode
// @Summary 获取package
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} return response.Response{data=map[string]interface{},msg=string} "创建package成功"
// @Router /autoCode/getPackage [post]
func (autoApi *AutoCodeApi) GetPackage(c *fiber.Ctx) error {
	pkgs, err := autoCodeService.GetPackage()
	if err != nil {
		global.LOG.Error("获取失败!", zap.Error(err))
		return response.FailWithMessage("获取失败", c)
	} else {
		return response.OkWithDetailed(fiber.Map{"pkgs": pkgs}, "获取成功", c)
	}
}

// DelPackage
// @Tags AutoCode
// @Summary 删除package
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysAutoCode true "创建package"
// @Success 200 {object} return response.Response{data=map[string]interface{},msg=string} "删除package成功"
// @Router /autoCode/delPackage [post]
func (autoApi *AutoCodeApi) DelPackage(c *fiber.Ctx) error {
	var a system.SysAutoCode
	_ = c.QueryParser(&a)
	err := autoCodeService.DelPackage(a)
	if err != nil {
		global.LOG.Error("删除失败!", zap.Error(err))
		return response.FailWithMessage("删除失败", c)
	} else {
		return response.OkWithMessage("删除成功", c)
	}
}
