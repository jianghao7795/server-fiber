package example

import (
	v1 "server-fiber/api/v1"
	"server-fiber/middleware"

	"github.com/gofiber/fiber/v2"
)

type ExcelRouter struct{}

func (e *ExcelRouter) InitExcelRouter(Router fiber.Router) {
	excelRouter := Router.Group("excel")
	exaExcelApi := v1.ApiGroupApp.ExampleApiGroup.ExcelApi

	excelRouter.Post("importExcel", middleware.OperationRecord, exaExcelApi.ImportExcel)          // 导入文件
	excelRouter.Get("getFileInfoList", exaExcelApi.GetFileList)                                   // 获取上传文件成功列表
	excelRouter.Delete("deleteFile/:id", middleware.OperationRecord, exaExcelApi.DeleteFile)      // 删除文件
	excelRouter.Get("loadExcel", middleware.OperationRecord, exaExcelApi.LoadExcel)               // 加载Excel数据
	excelRouter.Post("exportExcel", middleware.OperationRecord, exaExcelApi.ExportExcel)          // 导出Excel
	excelRouter.Get("downloadTemplate", middleware.OperationRecord, exaExcelApi.DownloadTemplate) // 下载模板文件

}
