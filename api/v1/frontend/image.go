package frontend

import (
	"server-fiber/global"
	"server-fiber/model/common/response"
	frontend "server-fiber/service/frontend"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

var frontendUploadService = new(frontend.FrontendImages)

func (u *FrontendUser) GetImages(c *fiber.Ctx) error {
	imageList, err := frontendUploadService.GetImagesList()
	if err != nil {
		global.LOG.Error("获取失败!", zap.Error(err))
		return response.FailWithMessage("获取失败", c)
	} else {
		return response.OkWithDetailed(imageList, "获取成功", c)
	}
}
