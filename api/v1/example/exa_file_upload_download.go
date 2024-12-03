package example

import (
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"server-fiber/model/common/request"
	"server-fiber/model/common/response"
	"server-fiber/model/example"
	"strconv"

	global "server-fiber/model"

	fileDimensionReq "server-fiber/model/example/request"
	exampleRes "server-fiber/model/example/response"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

// @Tags ExaFileUploadAndDownload
// @Summary 上传文件示例
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param file formData file true "上传文件示例"
// @Success 200 {object} response.Response{data=exampleRes.ExaFileResponse,msg=string} "上传文件示例,返回包括文件详情"
// @Router /fileUploadAndDownload/upload [post]
func (u *FileUploadAndDownloadApi) UploadFile(c *fiber.Ctx) error {
	var file example.ExaFileUploadAndDownload
	noSave := c.Query("noSave", "0")
	isCropper, err := strconv.Atoi(c.Query("is_cropper", "1"))
	if err != nil {
		global.LOG.Error("获取是否为裁剪图片失败", zap.Error(err))
	}
	fileImages, err := c.FormFile("file")
	if err != nil {
		global.LOG.Error("接收文件失败!", zap.Error(err))
		return response.FailWithMessage("接收文件失败", c)
	}
	if fileImages.Size > global.CONFIG.Local.Size*1024*1024 {
		global.LOG.Error("文件大小超过10M!")
		return response.FailWithMessage("文件大小超过10M", c)
	} // 文件大小限制10M
	header := c.Get("content-type")
	if string(header) == "image/svg+xml" {
		var fileDimension fileDimensionReq.FileDimension
		fileDimension.Height = 2
		fileDimension.Width = 1
		fileDimension.Proportion = 2.00
		file, err = fileUploadAndDownloadService.UploadFile(fileImages, noSave, &fileDimension, isCropper) // 文件上传后拿到文件路径
		if err != nil {
			global.LOG.Error("修改数据库链接失败!", zap.Error(err))
			return response.FailWithMessage("修改数据库链接失败", c)
		}
	} else {
		reader, _ := fileImages.Open()
		ct, _, err := image.Decode(reader)
		if err != nil {
			global.LOG.Error("获取文件失败!", zap.Error(err))
			return response.FailWithMessage("获取文件失败", c)
		}
		fileCtx := ct.Bounds()
		var fileDimension fileDimensionReq.FileDimension
		fileDimension.Height = fileCtx.Dy()
		fileDimension.Width = fileCtx.Dx()
		fileDimension.Proportion = float64(fileCtx.Dx()) / float64(fileCtx.Dy())

		file, err = fileUploadAndDownloadService.UploadFile(fileImages, noSave, &fileDimension, isCropper) // 文件上传后拿到文件路径
		if err != nil {
			global.LOG.Error("修改数据库链接失败!", zap.Error(err))
			return response.FailWithMessage("修改数据库链接失败"+err.Error(), c)
		}

		defer reader.Close()
	}

	return response.OkWithDetailed(exampleRes.ExaFileResponse{File: file}, "上传成功", c)
}

// EditFileName 编辑文件名或者备注
func (u *FileUploadAndDownloadApi) EditFileName(c *fiber.Ctx) error {
	var data example.ExaFileUploadAndDownload
	if err := c.BodyParser(&data); err != nil {
		global.LOG.Error("获取文件失败!", zap.Error(err))
		return response.FailWithMessage("获取文件失败"+err.Error(), c)
	}
	if err := fileUploadAndDownloadService.EditFileName(&data); err != nil {
		global.LOG.Error("编辑失败!", zap.Error(err))
		return response.FailWithMessage("编辑失败"+err.Error(), c)
	}
	return response.OkWithMessage("编辑成功", c)
}

// @Tags ExaFileUploadAndDownload
// @Summary 删除文件
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body example.ExaFileUploadAndDownload true "传入文件里面id即可"
// @Success 200 {object} response.Response{msg=string} "删除文件"
// @Router /fileUploadAndDownload/deleteFile [delete]
func (u *FileUploadAndDownloadApi) DeleteFile(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		global.LOG.Error("获取id失败!", zap.Error(err))
		return response.FailWithMessage("获取id失败"+err.Error(), c)
	}
	if err := fileUploadAndDownloadService.DeleteFile(uint(id)); err != nil {
		global.LOG.Error("删除失败!", zap.Error(err))
		return response.FailWithDetailed(err.Error(), "删除失败", c)
	}
	return response.OkWithMessage("删除成功", c)
}

// @Tags ExaFileUploadAndDownload
// @Summary 分页文件列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "页码, 每页大小"
// @Success 200 {object} response.Response{data=response.PageResult{list=example.ExaFileUploadAndDownload[]},msg=string} "分页文件列表,返回包括列表,总数,页码,每页数量"
// @Router /fileUploadAndDownload/getFileList [get]
func (u *FileUploadAndDownloadApi) GetFileList(c *fiber.Ctx) error {
	var pageInfo request.PageInfo
	_ = c.QueryParser(&pageInfo)
	list, total, err := fileUploadAndDownloadService.GetFileRecordInfoList(&pageInfo)
	if err != nil {
		global.LOG.Error("获取失败!", zap.Error(err))
		return response.FailWithMessage("获取失败"+err.Error(), c)
	} else {
		return response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
