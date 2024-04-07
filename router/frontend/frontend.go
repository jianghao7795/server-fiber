package frontend

import (
	fileUpload "server-fiber/api/v1/app"
	v1 "server-fiber/api/v1/frontend"
	"server-fiber/middleware"

	"github.com/gofiber/fiber/v2"
)

type FrontendRouter struct{}

func (s *FrontendRouter) InitFrontendRouter(Router fiber.Router) {
	frontend := Router.Group("")
	var frontendTagApi = new(v1.FrontendTagApi)
	{
		frontend.Get("getTagList", frontendTagApi.GetTagList)
		frontend.Get("getTagArticleList/:id", frontendTagApi.GetTag)
	}
	var frontendArticleApi = new(v1.FrontendArticleApi)
	{
		frontend.Get("getArticleList", frontendArticleApi.GetArticleList)
		frontend.Get("getArticle/:id", frontendArticleApi.GetArticleDetail)
		frontend.Get("getSearchArticle/:name/:value", frontendArticleApi.GetSearchArticle)
	}
	var frontendCommentApi = new(v1.CommentApi)
	{
		frontend.Get("getArticleComment/:articleId", frontendCommentApi.GetCommentByArticleId)
		frontend.Post("createdComment", middleware.OperationRecord, frontendCommentApi.CreatedComment)
	}
	var frontendUserApi = new(v1.FrontendUser)
	{
		frontend.Get("getImages", middleware.JWTAuth, frontendUserApi.GetImages)
		frontend.Post("login", frontendUserApi.Login)
		frontend.Get("getCurrentUser", middleware.JWTAuth, frontendUserApi.GetCurrent)
		frontend.Put("updateBackgroundImage", middleware.JWTAuth, middleware.OperationRecord, frontendUserApi.UpdateUserBackgroudImage)
		frontend.Put("resetPassword", middleware.JWTAuth, middleware.OperationRecord, frontendUserApi.UpdatePassword)
		frontend.Post("register", middleware.JWTAuth, frontendUserApi.RegisterUser)
		frontend.Put("updateUser", middleware.JWTAuth, middleware.OperationRecord, frontendUserApi.UpdateUser)
	}
	var frontendUploadApi = new(fileUpload.FileUploadAndDownloadApi)
	{
		frontend.Post("upload", frontendUploadApi.UploadFile)
	}
}
