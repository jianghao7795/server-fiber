/*
 * @Author: jianghao
 * @Date: 2022-10-17 11:13:10
 * @LastEditors: jianghao
 * @LastEditTime: 2022-10-17 17:15:16
 */

package app

import (
	v1 "server-fiber/api/v1"
	"server-fiber/middleware"

	"github.com/gofiber/fiber/v2"
)

type BaseMessageRouter struct{}

// InitArticleRouter 初始化 base_message 路由信息
func (r *BaseMessageRouter) InitBaseMessageRouter(c fiber.Router) {
	baseMessageRouter := c.Group("base_message")
	var baseMessageApi = v1.ApiGroupApp.AppApiGroup.BaseMessageApi
	var uploadFileApi = v1.ApiGroupApp.AppApiGroup.FileUploadAndDownloadApi

	baseMessageRouter.Post("createBaseMessage", middleware.OperationRecord, baseMessageApi.CreateBaseMessage)
	baseMessageRouter.Put("updateBaseMessage", middleware.OperationRecord, baseMessageApi.UpdateBaseMessage)
	baseMessageRouter.Post("upload_file", middleware.OperationRecord, uploadFileApi.UploadFile)

	baseMessageRouter.Get("getBaseMessage/:id", baseMessageApi.FindBaseMessage)

}
