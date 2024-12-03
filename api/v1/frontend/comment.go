package frontend

import (
	global "server-fiber/model"
	"server-fiber/model/common/response"
	"server-fiber/model/frontend"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type CommentApi struct{}

func (s *CommentApi) GetCommentByArticleId(c *fiber.Ctx) error {
	id, err := c.ParamsInt("articleId")
	if err != nil {
		global.LOG.Error("获取articleId 失败!", zap.Error(err))
		return response.FailWithMessage("获取articleId 失败", c)
	}
	if articleComment, err := commentServiceApp.GetCommentByArticleId(id); err != nil {
		global.LOG.Error("获取失败!", zap.Error(err))
		return response.FailWithMessage("获取失败", c)
	} else {
		return response.OkWithDetailed(articleComment, "获取成功", c)
	}
}

func (s *CommentApi) CreatedComment(c *fiber.Ctx) error {
	var comment frontend.Comment
	if err := c.BodyParser(&comment); err != nil {
		global.LOG.Error("获取数据失败", zap.Error(err))
		return response.FailWithMessage(err.Error(), c)
	}
	if err := commentServiceApp.CreatedComment(&comment); err != nil {
		global.LOG.Error("评论失败!", zap.Error(err))
		return response.FailWithMessage("评论失败", c)
	} else {
		return response.OkWithId("评论成功", comment.ID, c)
	}
}
