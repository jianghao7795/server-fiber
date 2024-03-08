package frontend

import (
	"server-fiber/global"
	appReq "server-fiber/model/app/request"
	"server-fiber/model/common/response"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type FrontendTagApi struct{}

func (appTabApi *FrontendTagApi) GetTagList(c *fiber.Ctx) error {
	var pageInfo appReq.TagSearch
	_ = c.QueryParser(&pageInfo)
	if list, err := frontendService.FrontendTag.GetTagList(c); err != nil {
		global.LOG.Error("获取失败!", zap.Error(err))
		return response.FailWithMessage("获取失败", c)
	} else {
		return response.OkWithDetailed(response.PageResult{
			List: list,
		}, "获取成功", c)
	}
}

func (appTabApi *FrontendTagApi) GetTag(c *fiber.Ctx) error {
	id := c.Params("id")
	tagId, err := strconv.Atoi(id)
	if err != nil {
		return response.FailWithMessage("获取Ids失败", c)
	}
	if tagArticles, err := frontendService.FrontendTag.GetTagArticle(tagId, c); err != nil {
		global.LOG.Error("获取失败!", zap.Error(err))
		return response.FailWithMessage("获取失败", c)
	} else {
		return response.OkWithData(gin.H{"tag": tagArticles}, c)
	}
}
