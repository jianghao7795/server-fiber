/*
 * @Author: jianghao
 * @Date: 2022-10-17 15:45:55
 * @LastEditors: jianghao
 * @LastEditTime: 2022-10-17 17:04:47
 */

package app

import (
	"server-fiber/global"
	"server-fiber/model/app"
	"server-fiber/model/common/response"
	appService "server-fiber/service/app"

	responseUplodFile "server-fiber/model/app/response"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type FileUploadAndDownloadApi struct{}

var fileUploadService = new(appService.FileUploadService)

// UploadFile @Tags UploadFile ExaFileUploadAndDownload
// @Summary 上传文件示例
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param file formData file true "上传文件示例"
// @Success 200 {object} return response.Response{data=string,msg=string} "上传文件示例,返回包括文件详情"
// @Router /fileUploadAndDownload/upload [post]
func (u *FileUploadAndDownloadApi) UploadFile(c *fiber.Ctx) error {
	var file app.FileUploadAndDownload
	noSave := c.Query("noSave", "0")
	header, err := c.FormFile("file")
	if err != nil {
		global.LOG.Error("接收文件失败!", zap.Error(err))
		return response.FailWithMessage400("接收文件失败", c)
	}
	file, err = fileUploadService.UploadFile(header, noSave) // 文件上传后拿到文件路径
	if err != nil {
		global.LOG.Error("修改数据库链接失败!", zap.Error(err))
		return response.FailWithMessage("修改数据库链接失败", c)
	}
	return response.OkWithDetailed(responseUplodFile.ResponseUploadFile{File: file}, "上传成功", c)
}
