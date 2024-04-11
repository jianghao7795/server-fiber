package app

import (
	"server-fiber/global"
	comment "server-fiber/model/app"
	commentReq "server-fiber/model/app/request"
	"server-fiber/model/common/request"
	"server-fiber/model/common/response"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

// CreateComment 创建Comment
// @Tags Comment
// @Summary 创建Comment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body comment.Comment true "创建Comment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /comment/createComment [post]
func (commentApi *CommentApi) CreateComment(c *fiber.Ctx) error {
	var commentData comment.Comment
	err := c.BodyParser(&commentData)
	if err != nil {
		global.LOG.Error("获取数据失败!", zap.Error(err))
		return response.FailWithMessage("获取数据失败", c)
	}
	if err := commentService.CreateComment(&commentData); err != nil {
		global.LOG.Error("创建失败!", zap.Error(err))
		return response.FailWithMessage("创建失败", c)
	} else {
		return response.OkWithMessage("创建成功", c)
	}
}

// DeleteComment 删除Comment
// @Tags Comment
// @Summary 删除Comment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body comment.Comment true "删除Comment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /comment/deleteComment [delete]
func (commentApi *CommentApi) DeleteComment(c *fiber.Ctx) error {
	var comment2 comment.Comment
	_ = c.BodyParser(&comment2)
	if err := commentService.DeleteComment(comment2.ID); err != nil {
		global.LOG.Error("删除失败!", zap.Error(err))
		return response.FailWithDetailed(fiber.Map{"msg": err.Error()}, "删除失败", c)
	} else {
		return response.OkWithMessage("删除成功", c)
	}
}

// DeleteCommentByIds 批量删除Comment
// @Tags Comment
// @Summary 批量删除Comment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Comment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /comment/deleteCommentByIds [delete]
func (commentApi *CommentApi) DeleteCommentByIds(c *fiber.Ctx) error {
	var IDS request.IdsReq
	_ = c.BodyParser(&IDS)
	if err := commentService.DeleteCommentByIds(IDS); err != nil {
		global.LOG.Error("批量删除失败!", zap.Error(err))
		return response.FailWithMessage("批量删除失败", c)
	} else {
		return response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateComment 更新Comment
// @Tags Comment
// @Summary 更新Comment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body comment.Comment true "更新Comment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /comment/updateComment [put]
func (commentApi *CommentApi) UpdateComment(c *fiber.Ctx) error {
	var comment2 comment.Comment
	err := c.BodyParser(&comment2)
	if err != nil {
		global.LOG.Error("获取数据失败!", zap.Error(err))
		return response.FailWithMessage("获取数据失败", c)
	}
	if err = commentService.UpdateComment(&comment2); err != nil {
		global.LOG.Error("更新失败!", zap.Error(err))
		return response.FailWithMessage("更新失败", c)
	} else {
		return response.OkWithMessage("更新成功", c)
	}
}

// FindComment 用id查询Comment
// @Tags Comment
// @Summary 用id查询Comment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query comment.Comment true "用id查询Comment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /comment/findComment [get]
func (commentApi *CommentApi) FindComment(c *fiber.Ctx) error {
	var comment2 comment.Comment
	_ = c.QueryParser(&comment2)
	if recomment, err := commentService.GetComment(comment2.ID); err != nil {
		global.LOG.Error("查询失败!", zap.Error(err))
		return response.FailWithMessage("查询失败", c)
	} else {
		return response.OkWithData(fiber.Map{"recomment": recomment}, c)
	}
}

// GetCommentList 分页获取Comment列表
// @Tags Comment
// @Summary 分页获取Comment列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query commentReq.CommentSearch true "分页获取Comment列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /comment/getCommentList [get]
func (commentApi *CommentApi) GetCommentList(c *fiber.Ctx) error {
	var pageInfo commentReq.CommentSearch
	_ = c.QueryParser(&pageInfo)
	if list, total, err := commentService.GetCommentInfoList(&pageInfo); err != nil {
		global.LOG.Error("获取失败!", zap.Error(err))
		return response.FailWithMessage("获取失败", c)
	} else {
		return response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

func (*CommentApi) GetCommentTreeList(c *fiber.Ctx) error {
	var pageInfo commentReq.CommentSearch
	_ = c.QueryParser(&pageInfo)

	if list, total, err := commentService.GetCommentTreeList(&pageInfo); err != nil {
		global.LOG.Error("获取失败!", zap.Error(err))
		return response.FailWithMessage("获取失败", c)
	} else {
		return response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// PutLikeItOrDislike 点赞
func (*CommentApi) PutLikeItOrDislike(c *fiber.Ctx) error {
	var likeIt comment.Praise
	_ = c.QueryParser(&likeIt)

	if err := commentService.PutLikeItOrDislike(&likeIt); err != nil {
		global.LOG.Error("点赞失败!", zap.Error(err))
		return response.FailWithDetailed(err, "点赞失败", c)
	} else {
		return response.OkWithDetailed(likeIt, "点赞成功", c)
	}
}
