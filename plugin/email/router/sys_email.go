package router

import (
	"server-fiber/plugin/email/api"

	"github.com/gofiber/fiber/v2"
)

type EmailRouter struct{}

func (s *EmailRouter) InitEmailRouter(Router fiber.Router) {
	emailRouter := Router
	EmailApi := api.ApiGroupApp.EmailApi.EmailTest
	SendEmail := api.ApiGroupApp.EmailApi.SendEmail
	{
		emailRouter.Post("emailTest", EmailApi)  // 发送测试邮件
		emailRouter.Post("sendEmail", SendEmail) // 发送邮件
	}
}
