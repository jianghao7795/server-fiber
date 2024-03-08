package frontend

import (
	v1 "server-fiber/api/v1"
	"server-fiber/middleware"

	"github.com/gofiber/fiber/v2"
)

type FrontendRouter struct{}

func (s *FrontendRouter) InitFrontendRouter(Router fiber.Router) {
	frontend := Router.Group("")
	var frontendTagApi = v1.ApiGroupApp.FrontendApiGroup.FrontendTagApi
	{
		frontend.Get("getTagList", frontendTagApi.GetTagList)
		frontend.Get("getTagArticleList/:id", frontendTagApi.GetTag)
	}
	var frontendArticleApi = v1.ApiGroupApp.FrontendApiGroup.FrontendArticleApi
	{
		frontend.Get("getArticleList", frontendArticleApi.GetArticleList)
		frontend.Get("getArticle/:id", frontendArticleApi.GetArticleDetail)
		frontend.Get("getSearchArticle/:name/:value", frontendArticleApi.GetSearchArticle)
	}
	var frontendCommentApi = v1.ApiGroupApp.FrontendApiGroup.CommentApi
	{
		frontend.Get("getArticleComment/:articleId", frontendCommentApi.GetCommentByArticleId)
		frontend.Post("createdComment", middleware.OperationRecord, frontendCommentApi.CreatedComment)
	}
	var frontendUserApi = v1.ApiGroupApp.FrontendApiGroup.FrontendUser
	{
		frontend.Get("getImages", middleware.JWTAuth, frontendUserApi.GetImages)
		frontend.Post("login", frontendUserApi.Login)
		frontend.Get("getCurrentUser", middleware.JWTAuth, frontendUserApi.GetCurrent)
		frontend.Put("updateBackgroundImage", middleware.JWTAuth, middleware.OperationRecord, frontendUserApi.UpdateUserBackgroudImage)
		frontend.Put("resetPassword", middleware.JWTAuth, middleware.OperationRecord, frontendUserApi.UpdatePassword)
		frontend.Post("register", middleware.JWTAuth, frontendUserApi.RegisterUser)
		frontend.Put("updateUser", middleware.JWTAuth, middleware.OperationRecord, frontendUserApi.UpdateUser)
	}
	var frontendUploadApi = v1.ApiGroupApp.AppApiGroup.FileUploadAndDownloadApi
	{
		frontend.Post("upload", frontendUploadApi.UploadFile)
	}
}
