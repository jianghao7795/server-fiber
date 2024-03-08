package mobile

import (
	v1 "server-fiber/api/v1"

	"github.com/gofiber/fiber/v2"
)

type MobileLoginRouter struct{}

func (m *MobileUserRouter) InitMobileLoginRouter(Router fiber.Router) {
	moblieLoginRouter := Router.Group("") //.Use(middleware.JWTAuthMobileMiddleware())
	var mobileLoginApi = v1.ApiGroupApp.MobileApiGroup.MobileLoginApi
	var registerApi = v1.ApiGroupApp.MobileApiGroup.RegisterMobile
	{
		moblieLoginRouter.Post("login", mobileLoginApi.Login)
		moblieLoginRouter.Post("register", registerApi.Register)
	}
	mobileGetUserApi := Router.Group("") //.Use(middleware.JWTAuthMobileMiddleware())
	{
		mobileGetUserApi.Get("getUserInfo", mobileLoginApi.GetUserInfo)
		mobileGetUserApi.Put("updateUser", mobileLoginApi.UpdateMobileUser)
		mobileGetUserApi.Put("updatePassword", mobileLoginApi.UpdatePassword)
	}
	exaFileUploadAndDownloadApi := v1.ApiGroupApp.ExampleApiGroup.FileUploadAndDownloadApi
	{
		mobileGetUserApi.Post("uploadImage", exaFileUploadAndDownloadApi.UploadFile)
	}

}
