package app

import (
	"strconv"

	"server-fiber/global"
	"server-fiber/model/app"
	appReq "server-fiber/model/app/request"
	"server-fiber/model/common/request"
	"server-fiber/model/common/response"
	"server-fiber/service"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type ArticleApi struct{}

var articleService = service.ServiceGroupApp.AppServiceGroup.ArticleService

// CreateArticle 创建Article
// @Tags Article
// @Summary 创建Article
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body app.Article true "创建Article"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Article/createArticle [post]
func (a *ArticleApi) CreateArticle(c *fiber.Ctx) error {
	var article app.Article
	_ = c.BodyParser(&article)
	if err := articleService.CreateArticle(article); err != nil {
		global.LOG.Error("创建失败!", zap.Error(err))
		return response.FailWithDetailed(map[string]string{
			"msg": err.Error(),
		}, "创建失败", c)

	}
	return response.OkWithMessage("创建成功", c)
}

// DeleteArticle 删除Article
// @Tags Article
// @Summary 删除Article
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body app.Article true "删除Article"
// @Success 200 {string} string "{}"
// @Router /article/deleteArticle/:id [delete]
func (*ArticleApi) DeleteArticle(c *fiber.Ctx) error {
	ids := c.Params("id")
	id, _ := strconv.Atoi(ids)
	if err := articleService.DeleteArticle(uint(id)); err != nil {
		global.LOG.Error("删除失败", zap.Error(err))
		return response.FailWithDetailed(map[string]string{
			"msg": err.Error(),
		}, "删除失败", c)

	}
	return response.OkWithMessage("删除成功", c)

}

// DeleteArticleByIds 批量删除Article
// @Tags Article
// @Summary 批量删除Article
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body app.Article true "批量删除Article"
// @Success 200 {string} string "{}"
// @Router /article/deleteArticleByIds [delete]
func (a *ArticleApi) DeleteArticleByIds(c *fiber.Ctx) error {
	var IDS request.IdsReq
	_ = c.QueryParser(&IDS)
	if err := articleService.DeleteArticleByIds(IDS); err != nil {
		global.LOG.Error("批量删除失败!", zap.Error(err))
		return response.FailWithDetailed(map[string]string{
			"msg": err.Error(),
		}, "批量删除失败", c)

	}
	return response.OkWithMessage("批量删除成功", c)

}

// UpdateArticle 更新Article
// @Tags Article
// @Summary 更新Article
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body app.Article true "更新Article"
// @Success 200 {string} string "{"success":true, "msg":"更新成功"}"
// @Router /article/updateArticle/:id [put]
func (*ArticleApi) UpdateArticle(c *fiber.Ctx) error {
	var article app.Article
	_ = c.BodyParser(&article)
	ids := c.Params("id")
	id, _ := strconv.Atoi(ids)
	article.ID = uint(id)
	if err := articleService.UpdateArticle(article); err != nil {
		global.LOG.Error("更新失败!", zap.Error(err))
		return response.FailWithDetailed(map[string]string{
			"msg": err.Error(),
		}, "更新失败", c)

	}
	return response.OkWithMessage("更新成功", c)

}

// FindArticle get单个Article
// @Tags Article
// @Summary get单个Article
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body app.Article true "get单个Article"
// @Success 200 {string} string "{"success":true, "msg":"获得成功"}"
// @Router /article/findArticle/:id [get]
func (*ArticleApi) FindArticle(c *fiber.Ctx) error {
	ids := c.Params("id")
	id, _ := strconv.Atoi(ids)
	if rearticle, err := articleService.GetArticle(uint(id)); err != nil {
		global.LOG.Error("查询失败!", zap.Error(err))
		return response.FailWithDetailed(map[string]string{
			"msg": err.Error(),
		}, "查询失败", c)
	} else {
		return response.OkWithData(fiber.Map{"rearticle": rearticle}, c)
	}

}

// GetArticleList 分页获取article列表
// FindArticle Get Article
// @Tags Article
// @Summary Get Article
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body app.Article true "Get Article"
// @Success 200 {string} string "{"success":true, "msg":"获得成功"}"
// @Router /article/getArticleList [get]
func (*ArticleApi) GetArticleList(c *fiber.Ctx) error {
	var pageInfo appReq.ArticleSearch
	_ = c.QueryParser(&pageInfo)
	if list, total, err := articleService.GetArticleInfoList(pageInfo); err != nil {
		global.LOG.Error("获取失败!", zap.Error(err))
		return response.FailWithDetailed(map[string]string{
			"msg": err.Error(),
		}, "获取失败", c)

	} else {
		return response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// PutArticleByIds 批量更新Article
// @Tags Article
// @Summary 批量更新Article
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body app.Article true "批量更新Article"
// @Success 200 {string} string "{}"
// @Router /article/putArticleByIds [put]
func (*ArticleApi) PutArticleByIds(c *fiber.Ctx) error {
	var IDS request.IdsReq
	_ = c.BodyParser(&IDS)
	if err := articleService.PutArticleByIds(IDS); err != nil {
		global.LOG.Error("批量更新失败!", zap.Error(err))
		return response.FailWithDetailed(map[string]string{
			"msg": err.Error(),
		}, "批量更新失败", c)

	} else {
		return response.OkWithMessage("批量更新成功", c)
	}
}

// GetArticleReading 获取文章阅读量
// @Tags Article
// @Summary 获取文章阅读量
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body app.Article true "获取文章阅读量"
// @Success 200 {string} string "{"code":0,"data":{},"msg":"获取成功"}"
// @Router /article/getArticleReading [get]
func (*ArticleApi) GetArticleReading(c *fiber.Ctx) error {
	count, err := articleService.GetArticleReading()
	if err != nil {
		global.LOG.Error("获取阅读量失败!", zap.Error(err))
		return response.FailWithDetailed(map[string]string{
			"msg": err.Error(),
		}, "获取阅读量失败", c)
	} else {
		return response.OkWithDetailed(fiber.Map{"reading_quantity": count}, "获取成功", c)
	}
}
