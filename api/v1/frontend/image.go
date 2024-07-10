package frontend

import (
	"server-fiber/global"
	"server-fiber/model/common/response"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func (u *User) GetImages(c *fiber.Ctx) error {
	imageList, err := imagesServiceApp.GetImagesList()
	if err != nil {
		global.LOG.Error("获取失败!", zap.Error(err))
		return response.FailWithMessage("获取失败", c)
	} else {
		return response.OkWithDetailed(imageList, "获取成功", c)
	}
}
