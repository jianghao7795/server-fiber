package example

import (
	v1 "server-fiber/api/v1"

	"github.com/gofiber/fiber/v2"
)

type FileUploadAndDownloadRouter struct{}

func (e *FileUploadAndDownloadRouter) InitFileUploadAndDownloadRouter(Router fiber.Router) {
	fileUploadAndDownloadRouter := Router.Group("fileUploadAndDownload")
	exaFileUploadAndDownloadApi := v1.ApiGroupApp.ExampleApiGroup.FileUploadAndDownloadApi
	{
		fileUploadAndDownloadRouter.Post("upload", exaFileUploadAndDownloadApi.UploadFile)                                 // 上传文件
		fileUploadAndDownloadRouter.Get("getFileList", exaFileUploadAndDownloadApi.GetFileList)                            // 获取上传文件列表
		fileUploadAndDownloadRouter.Delete("deleteFile/:id", exaFileUploadAndDownloadApi.DeleteFile)                       // 删除指定文件
		fileUploadAndDownloadRouter.Put("editFileName", exaFileUploadAndDownloadApi.EditFileName)                          // 编辑文件名或者备注
		fileUploadAndDownloadRouter.Post("breakpointContinue", exaFileUploadAndDownloadApi.BreakpointContinue)             // 断点续传
		fileUploadAndDownloadRouter.Get("findFile", exaFileUploadAndDownloadApi.FindFile)                                  // 查询当前文件成功的切片
		fileUploadAndDownloadRouter.Post("breakpointContinueFinish", exaFileUploadAndDownloadApi.BreakpointContinueFinish) // 切片传输完成
		fileUploadAndDownloadRouter.Post("removeChunk", exaFileUploadAndDownloadApi.RemoveChunk)                           // 删除切片
	}
}
