package middleware

import (

	// "io/ioutil"

	"server-fiber/global"
	"server-fiber/model/system"
	"server-fiber/plugin/email/utils"
	"server-fiber/service"
	utils2 "server-fiber/utils"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

var userService = service.ServiceGroupApp.SystemServiceGroup.UserService

func ErrorToEmail(c *fiber.Ctx) error {
	var username string
	claims, _ := utils2.GetClaims(c)
	if claims.Username != "" {
		username = claims.Username
	} else {
		id, _ := strconv.Atoi(c.Get("x-user-id"))
		user, err := userService.FindUserById(id)
		if err != nil {
			username = "Unknown"
		} else {
			username = user.Username
		}
	}
	body := c.Request().Body()
	// 再重新写回请求体body中，ioutil.ReadAll会清空c.Request.Body中的数据
	// c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	record := system.SysOperationRecord{
		Ip:       c.IP(),
		Method:   c.Method(),
		Path:     c.Path(),
		Agent:    string(c.Request().Header.UserAgent()),
		Body:     string(body),
		TypePort: system.Backend,
	}
	now := time.Now()
	latency := time.Since(now)
	status := c.Response().StatusCode()
	record.ErrorMessage = string(c.Response().Body())
	str := "接收到的请求为" + record.Body + "\n" + "请求方式为" + record.Method + "\n" + "报错信息如下" + record.ErrorMessage + "\n" + "耗时" + latency.String() + "\n"
	if status != 200 {
		subject := username + "" + record.Ip + "调用了" + record.Path + "报错了"
		if err := utils.ErrorToEmail(subject, str); err != nil {
			global.LOG.Error("ErrorToEmail Failed, err:", zap.Error(err))
		}
	}
	return c.Next()
}
