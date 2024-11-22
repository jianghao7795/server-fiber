package middleware

import (
	"github.com/gofiber/fiber/v2"
)

// https

// 用https把这个中间件在router里面use一下就好

func LoadTls(c *fiber.Ctx) error {
	//middleware := secure.New(secure.Options{
	//	SSLRedirect: true,
	//	SSLHost:     "localhost:443",
	//})
	//err := middleware.Process(c.Response(), c.Request())
	//if err != nil {
	//	// 如果出现错误，请不要继续
	//	return nil
	//}
	// 继续往下处理
	return c.Next()
}
