/*
 * @Author: jianghao
 * @Date: 2022-10-17 11:13:10
 * @LastEditors: jianghao
 * @LastEditTime: 2022-10-17 17:15:16
 */

package app

import (
	v1 "server-fiber/api/v1"

	"github.com/gofiber/fiber/v2"
)

type BaseMessageRouter struct{}

// InitArticleRouter 初始化 base_message 路由信息
func (r *BaseMessageRouter) InitBaseMessageRouter(c fiber.Router) {
	baseMessageRouter := c.Group("base_message") //.Use(middleware.OperationRecord)
	baseMessageRouterWithoutRecord := c.Group("base_message")
	var baseMessageApi = v1.ApiGroupApp.AppApiGroup.BaseMessageApi
	var uploadFileApi = v1.ApiGroupApp.AppApiGroup.FileUploadAndDownloadApi
	{
		baseMessageRouter.Post("createBaseMessage", baseMessageApi.CreateBaseMessage)
		baseMessageRouter.Put("updateBaseMessage", baseMessageApi.UpdateBaseMessage)
		baseMessageRouter.Post("upload_file", uploadFileApi.UploadFile)
	}
	{
		baseMessageRouterWithoutRecord.Get("getBaseMessage/:id", baseMessageApi.FindBaseMessage)
	}
}
