package frontend

import (
	"errors"
	"server-fiber/global"
	"server-fiber/model/common/response"
	"server-fiber/model/frontend/request"
	frontend "server-fiber/service/frontend"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type FrontendArticleApi struct{}

var frontendArtileService = new(frontend.FrontendArticle)

// GetArticleList 分页获取article列表
// FindArticle Get Article
// @Tags Article
// @Summary Get Article
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body frontend.Article true "Get Article"
// @Success 200 {string} string "{"success":true, "msg":"获得成功"}"
// @Router /getArticleList [get]
func (s *FrontendArticleApi) GetArticleList(c *fiber.Ctx) error {
	var pageInfo request.ArticleSearch
	_ = c.QueryParser(&pageInfo)
	if pageInfo.Page == 0 {
		pageInfo.Page = 1
	}

	if pageInfo.PageSize == 0 {
		pageInfo.PageSize = 10
	}

	if list, total, err := frontendArtileService.GetArticleList(&pageInfo, c); err != nil {
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

// FindArticle get单个Article
// @Tags Article
// @Summary get单个Article
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body frontend.Article true "get单个Article"
// @Success 200 {string} string "{"success":true, "msg":"获得成功"}"
// @Router /getArticle/:id [get]
func (s *FrontendArticleApi) GetArticleDetail(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		global.LOG.Error("获取Id失败!", zap.Error(err))
		return response.FailWithMessage("获取Id失败", c)

	}
	articleDetail, err := frontendArtileService.GetAricleDetail(id, c)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return response.FailWithMessage("文章没有，请重新查询", c)

	}
	if err != nil {
		global.LOG.Error("获取失败!", zap.Error(err))
		return response.FailWithMessage("获取失败", c)

	} else {
		return response.OkWithData(fiber.Map{"article": articleDetail}, c)
	}
}

// FindArticle get单个Article
// @Tags Article
// @Summary get单个Article
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body frontend.Article true "Serach Artcle"
// @Success 200 {string} string "{"success":true, "msg":"获得成功"}"
// @Router /getSearchArticle/:name/:value [get]
func (s *FrontendArticleApi) GetSearchArticle(c *fiber.Ctx) error {
	var searchValue request.ArticleSearch
	err := c.ParamsParser(&searchValue)
	if err != nil {
		global.LOG.Error("获取数据失败!", zap.Error(err))
		return response.FailWithMessage("获取数据失败", c)
	}
	searchValue.Sort = c.Query("sort")
	if searchValue.Name != "tags" && searchValue.Name != "articles" {
		return response.FailWithMessage("查询的不是tag 或 article", c)
	}
	if list, err := frontendArtileService.GetSearchArticle(&searchValue); err != nil {
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
