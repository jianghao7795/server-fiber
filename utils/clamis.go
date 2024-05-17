package utils

import (
	"server-fiber/global"
	"server-fiber/model/common/response"
	systemReq "server-fiber/model/system/request"
	"strings"

	"github.com/gofiber/fiber/v2"
	uuid "github.com/satori/go.uuid"
)

// GetClaims 从fiber的Context中获取JWT token并解析出CustomClaims结构体
func GetClaims(c *fiber.Ctx) (*systemReq.CustomClaims, error) {
	tokenString := c.Get("Authorization")                   // 从请求头中获取token字符串
	token := strings.Replace(tokenString, "Bearer ", "", 1) // 去除token字符串中的"Bearer "前缀
	if token == "" {
		return nil, response.FailWithMessage401("token 失效， 请重新登录", c) // 若token为空则返回401错误响应
	}
	j := NewJWT()                      // 初始化JWT对象
	claims, err := j.ParseToken(token) // 解析token获取CustomClaims结构体和可能的错误
	if err != nil {
		global.LOG.Error("从fiber的Context中获取从jwt解析信息失败, 请检查请求头是否存在token且claims是否为规定结构") // 若解析出错则记录错误日志
	}
	return claims, err // 返回解析出的CustomClaims结构体和错误
}

// 从fiber的Context中获取从jwt解析出来的用户ID
func GetUserID(c *fiber.Ctx) uint {
	var claims = c.Locals("claims")
	if claims == nil {
		// 处理claims为nil的情况，例如返回错误信息或默认值
		return 0
	}

	waitUse, ok := claims.(*systemReq.CustomClaims)
	if !ok {
		// 处理类型断言失败的情况，例如返回错误信息或默认值
		return 0
	}

	if waitUse.BaseClaims.ID == 0 {
		// 处理ID为0的情况，可能需要调用GetClaims函数重新获取claims
		if cl, err := GetClaims(c); err != nil {
			// 处理获取claims失败的情况，例如返回错误信息或默认值
			return 0
		} else {
			return uint(cl.BaseClaims.ID)
		}
	} else {
		return waitUse.BaseClaims.ID
	}
}

// 从fiber的Context中获取从jwt解析出来的用户UUID
func GetUserUuid(c *fiber.Ctx) uuid.UUID {
	var claims = c.Locals("claims")
	waitUse := claims.(*systemReq.CustomClaims)
	if waitUse.BaseClaims.ID == 0 {
		if cl, err := GetClaims(c); err != nil {
			return uuid.UUID{}
		} else {
			return cl.UUID
		}
	} else {
		return waitUse.UUID
	}

}

// 从fiber的Context中获取从jwt解析出来的用户角色id
func GetUserAuthorityId(c *fiber.Ctx) string {
	var claims = c.Locals("claims")
	waitUse := claims.(*systemReq.CustomClaims)
	if waitUse.BaseClaims.ID == 0 {
		if cl, err := GetClaims(c); err != nil {
			return ""
		} else {
			return cl.AuthorityId
		}
	} else {
		return waitUse.AuthorityId
	}

}

// 从fiber的Context中获取从jwt解析出来的用户角色id
func GetUserInfo(c *fiber.Ctx) *systemReq.CustomClaims {
	var claims = c.Locals("claims")
	waitUse := claims.(*systemReq.CustomClaims)
	if waitUse.BaseClaims.ID == 0 {
		if cl, err := GetClaims(c); err != nil {
			return nil
		} else {
			return cl
		}
	}
	return waitUse
}
