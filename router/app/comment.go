package app

import (
	v1 "server-fiber/api/v1"
	"server-fiber/middleware"

	"github.com/gofiber/fiber/v2"
)

type CommentRouter struct {
}

// InitCommentRouter 初始化 Comment 路由信息
func (s *CommentRouter) InitCommentRouter(Router fiber.Router) {
	commentRouter := Router.Group("comment").Use(middleware.OperationRecord)
	commentRouterWithoutRecord := Router.Group("comment")
	var commentApi = v1.ApiGroupApp.AppApiGroup.CommentApi
	{
		commentRouter.Post("createComment", commentApi.CreateComment)             // 新建Comment
		commentRouter.Delete("DeleteComment", commentApi.DeleteComment)           // 删除Comment
		commentRouter.Delete("DeleteCommentByIds", commentApi.DeleteCommentByIds) // 批量删除Comment
		commentRouter.Put("updateComment", commentApi.UpdateComment)              // 更新Comment
		commentRouter.Put("pariseComment", commentApi.PutLikeItOrDislike)         //点赞
	}
	{
		commentRouterWithoutRecord.Get("findComment", commentApi.FindComment)               // 根据ID获取Comment
		commentRouterWithoutRecord.Get("getCommentList", commentApi.GetCommentList)         // 获取Comment列表
		commentRouterWithoutRecord.Get("getCommentTreeList", commentApi.GetCommentTreeList) //  获取Comment Tree列表
	}
}
