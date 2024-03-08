package frontend

import (
	"errors"
	"server-fiber/global"
	"server-fiber/model/common/response"
	"server-fiber/model/frontend/request"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type FrontendArticleApi struct{}

func (s *FrontendArticleApi) GetArticleList(c *fiber.Ctx) error {
	var pageInfo request.ArticleSearch
	_ = c.QueryParser(&pageInfo)
	if pageInfo.Page == 0 {
		pageInfo.Page = 1
	}

	if pageInfo.PageSize == 0 {
		pageInfo.PageSize = 10
	}

	if list, total, err := frontendService.FrontendArticle.GetArticleList(pageInfo, c); err != nil {
		global.LOG.Error("获取失败!", zap.Error(err))
		return response.FailWithMessage("获取失败", c)
	} else {
		// log.Println("total is ", total)
		return response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

func (s *FrontendArticleApi) GetArticleDetail(c *fiber.Ctx) error {
	id := c.Params("id")
	aritcleId, err := strconv.Atoi(id)
	if err != nil {
		global.LOG.Error("获取Id失败!", zap.Error(err))
		return response.FailWithMessage("获取Id失败", c)

	}

	articleDetail, err := frontendService.FrontendArticle.GetAricleDetail(aritcleId, c)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return response.FailWithMessage("文章没有，请重新查询", c)

	}
	if err != nil {
		global.LOG.Error("获取失败!", zap.Error(err))
		return response.FailWithMessage("获取失败", c)

	} else {
		return response.OkWithData(gin.H{"article": articleDetail}, c)
	}
}

func (s *FrontendArticleApi) GetSearchArticle(c *fiber.Ctx) error {
	var searchValue request.ArticleSearch
	searchValue.Name = c.Params("name")
	searchValue.Value = c.Params("value")
	searchValue.Sort = c.Query("sort")
	if list, err := frontendService.GetSearchArticle(searchValue); err != nil {
		global.LOG.Error("获取失败!", zap.Error(err))
		return response.FailWithMessage("获取失败", c)
	} else {
		return response.OkWithDetailed(response.PageResult{
			List: list,
			// Total:    total,
			// Page:     pageInfo.Page,
			// PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
