package frontend

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"server-fiber/global"
	"server-fiber/model/common/response"
	"server-fiber/model/frontend"
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
	_ = c.BodyParser(&comment)
	if err := commentServiceApp.CreatedComment(&comment); err != nil {
		global.LOG.Error("评论失败!", zap.Error(err))
		return response.FailWithMessage("评论失败", c)
	} else {
		return response.OkWithId("评论成功", comment.ID, c)
	}

}
