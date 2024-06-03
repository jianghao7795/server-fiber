package example

import (
	ioutil "io"
	"mime/multipart"
	"strconv"

	"server-fiber/global"
	"server-fiber/model/common/response"
	"server-fiber/model/example"
	"server-fiber/model/example/request"
	exampleRes "server-fiber/model/example/response"

	"server-fiber/utils"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

// BreakpointContinue @Tags ExaFileUploadAndDownload
// @Summary 断点续传到服务器
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param file formData file true "an example for breakpoint resume, 断点续传示例"
// @Success 200 {object} response.Response{msg=string} "断点续传到服务器"
// @Router /fileUploadAndDownload/breakpointContinue [post]
func (u *FileUploadAndDownloadApi) BreakpointContinue(c *fiber.Ctx) error {
	fileMd5 := c.FormValue("fileMd5")
	fileName := c.FormValue("fileName")
	chunkMd5 := c.FormValue("chunkMd5")
	// log.Println("fileMd5:", fileMd5, "fileName:", fileName, "chunkMd5:", chunkMd5)
	chunkNumber, _ := strconv.Atoi(c.FormValue("chunkNumber"))
	chunkTotal, _ := strconv.Atoi(c.FormValue("chunkTotal"))
	FileHeader, err := c.FormFile("file")
	if err != nil {
		global.LOG.Error("接收文件失败!", zap.Error(err))
		return response.FailWithMessage("接收文件失败: "+err.Error(), c)
	}
	f, err := FileHeader.Open()
	if err != nil {
		global.LOG.Error("文件读取失败!", zap.Error(err))
		return response.FailWithMessage("文件读取失败"+err.Error(), c)
	}
	defer func(f multipart.File) {
		err := f.Close()
		if err != nil {
			global.LOG.Error("文件关闭失败!", zap.Error(err))
		}
	}(f)
	cen, err := ioutil.ReadAll(f)
	if err != nil {
		global.LOG.Error("文件分段读取失败!", zap.Error(err))
		return response.FailWithMessage("文件分段读取失败", c)
	}
	if !utils.CheckMd5(cen, chunkMd5) {
		global.LOG.Error("检查md5失败!", zap.Error(err))
		return response.FailWithMessage("检查md5失败", c)
	}
	file, err := fileUploadAndDownloadService.FindOrCreateFile(fileMd5, fileName, chunkTotal)
	if err != nil {
		global.LOG.Error("查找或创建记录失败!", zap.Error(err))
		return response.FailWithMessage("查找或创建记录失败", c)
	}
	paths, err := utils.BreakPointContinue(cen, fileName, chunkNumber, chunkTotal, fileMd5)
	if err != nil {
		global.LOG.Error("断点续传失败!", zap.Error(err))
		return response.FailWithMessage("断点续传失败: "+err.Error(), c)
	}

	if err = fileUploadAndDownloadService.CreateFileChunk(file.ID, paths, chunkNumber); err != nil {
		global.LOG.Error("创建文件记录失败!", zap.Error(err))
		return response.FailWithMessage("创建文件记录失败: "+err.Error(), c)
	}
	return response.OkWithMessage("切片创建成功", c)
}

// @Tags ExaFileUploadAndDownload
// @Summary 查找文件
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param file formData file true "Find the file, 查找文件"
// @Success 200 {object} response.Response{data=exampleRes.FileResponse,msg=string} "查找文件,返回包括文件详情"
// @Router /fileUploadAndDownload/findFile [get]
func (u *FileUploadAndDownloadApi) FindFile(c *fiber.Ctx) error {
	fileMd5 := c.Query("fileMd5")
	fileName := c.Query("fileName")
	chunkTotal, _ := strconv.Atoi(c.Query("chunkTotal"))
	file, err := fileUploadAndDownloadService.FindOrCreateFile(fileMd5, fileName, chunkTotal)
	if err != nil {
		global.LOG.Error("查找失败!", zap.Error(err))
		return response.FailWithMessage("查找失败: "+err.Error(), c)
	} else {
		return response.OkWithDetailed(exampleRes.FileResponse{File: file}, "查找成功", c)
	}
}

// @Tags ExaFileUploadAndDownload
// @Summary 创建文件
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param file formData file true "上传文件完成"
// @Success 200 {object} response.Response{data=exampleRes.FilePathResponse,msg=string} "创建文件,返回包括文件路径"
// @Router /fileUploadAndDownload/findFile [post]
func (b *FileUploadAndDownloadApi) BreakpointContinueFinish(c *fiber.Ctx) error {
	var file request.BreakPoint
	err := c.BodyParser(&file)
	if err != nil {
		global.LOG.Error("获取文件信息错误", zap.Error(err))
		return response.FailWithMessage("获取文件信息错误: "+err.Error(), c)
	}
	filePath, err := utils.MakeFile(file.FileName, file.FileMd5)
	if err != nil {
		global.LOG.Error("文件创建失败!", zap.Error(err))
		return response.FailWithDetailed(exampleRes.FilePathResponse{FilePath: filePath}, "文件创建失败", c)
	} else {
		return response.OkWithDetailed(exampleRes.FilePathResponse{FilePath: filePath}, "文件创建成功", c)
	}
}

// @Tags ExaFileUploadAndDownload
// @Summary 删除切片
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param file formData file true "删除缓存切片"
// @Success 200 {object} response.Response{msg=string} "删除切片"
// @Router /fileUploadAndDownload/removeChunk [delete]
func (u *FileUploadAndDownloadApi) RemoveChunk(c *fiber.Ctx) error {
	var file example.ExaFile
	c.BodyParser(&file)
	err := utils.RemoveChunk(file.FileMd5)
	if err != nil {
		global.LOG.Error("缓存切片文件删除失败!", zap.Error(err))
		return response.FailWithDetailed(err, "缓存切片文件删除失败", c)
	}
	err = fileUploadAndDownloadService.DeleteFileChunk(file.FileMd5, file.FileName, file.FilePath)
	if err != nil {
		global.LOG.Error(err.Error(), zap.Error(err))
		return response.FailWithMessage(err.Error(), c)
	} else {
		return response.OkWithMessage("表缓存切片删除成功", c)
	}
}

func (u *FileUploadAndDownloadApi) FindFileBreakpoint(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	file, total, err := fileUploadAndDownloadService.FindFileBreakpoint(page, pageSize)
	if err != nil {
		global.LOG.Error("查找失败!", zap.Error(err))
		return response.FailWithMessage("获取失败: "+err.Error(), c)
	} else {
		return response.OkWithDetailed(response.PageResult{
			List:     file,
			Total:    total,
			Page:     page,
			PageSize: pageSize,
		}, "查找成功", c)
	}
}

func (u *FileUploadAndDownloadApi) DeleteFileBreakpoint(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		global.LOG.Error("获取id失败!", zap.Error(err))
		return response.FailWithMessage("获取id失败: "+err.Error(), c)
	}
	err = fileUploadAndDownloadService.DeleteFileBreakpoint(id)
	if err != nil {
		global.LOG.Error("删除失败!", zap.Error(err))
		return response.FailWithMessage("删除失败: "+err.Error(), c)
	} else {
		return response.OkWithMessage("删除失败", c)
	}
}
