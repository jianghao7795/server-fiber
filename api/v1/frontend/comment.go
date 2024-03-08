package frontend

import (
	"server-fiber/global"
	"server-fiber/model/common/response"
	"server-fiber/model/frontend"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type CommentApi struct{}

var commentService = frontendService.Comment

func (s *CommentApi) GetCommentByArticleId(c *fiber.Ctx) error {
	articleId := c.Params("articleId")
	id, err := strconv.Atoi(articleId)
	if err != nil {
		global.LOG.Error("获取articleId 失败!", zap.Error(err))
		return response.FailWithMessage("获取articleId 失败", c)
	}
	if articleComment, err := commentService.GetCommentByArticleId(id); err != nil {
		global.LOG.Error("获取失败!", zap.Error(err))
		return response.FailWithMessage("获取失败", c)
	} else {
		return response.OkWithDetailed(articleComment, "获取成功", c)
	}
}

func (s *CommentApi) CreatedComment(c *fiber.Ctx) error {
	var comment frontend.Comment
	_ = c.QueryParser(&comment)
	if id, err := commentService.CreatedComment(comment); err != nil {
		global.LOG.Error("创建失败!", zap.Error(err))
		return response.FailWithMessage("创建失败", c)
	} else {
		return response.OkWithId("创建成功", id, c)
	}

}
