/*
 * @Author: jianghao
 * @Date: 2022-10-17 15:48:45
 * @LastEditors: jianghao
 * @LastEditTime: 2022-10-17 15:51:35
 */

package app

import (
	"mime/multipart"
	"server-fiber/global"
	"server-fiber/model/app"
	"server-fiber/utils/upload"
	"strings"
)

type FileUploadService struct{}

//@author: wuhao
//@function: UploadFile
//@description: 根据配置文件判断是文件上传到本地或者七牛云
//@param: header *multipart.FileHeader, noSave string
//@return: err error, file model.ExaFileUploadAndDownload

func (e *FileUploadService) UploadFile(header *multipart.FileHeader, noSave string) (file app.FileUploadAndDownload, err error) {
	oss := upload.NewOss()
	filePath, key, uploadErr := oss.UploadFile(header)
	if uploadErr != nil {
		return app.FileUploadAndDownload{}, uploadErr
	}
	if noSave == "0" {
		s := strings.Split(header.Filename, ".")
		f := app.FileUploadAndDownload{
			Url:  filePath,
			Name: header.Filename,
			Tag:  s[len(s)-1],
			Key:  key,
		}
		return f, e.Upload(f)
	}
	return
}

//@author: wuhao
//@function: Upload
//@description: 创建文件上传记录
//@param: file model.ExaFileUploadAndDownload
//@return: error

func (e *FileUploadService) Upload(file app.FileUploadAndDownload) error {
	return global.DB.Create(&file).Error
}
