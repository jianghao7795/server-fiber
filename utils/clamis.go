package utils

import (
	"errors"
	"server-fiber/global"
	systemReq "server-fiber/model/system/request"
	"strings"

	"github.com/gofiber/fiber/v2"
	uuid "github.com/satori/go.uuid"
)

func GetClaims(c *fiber.Ctx) (*systemReq.CustomClaims, error) {
	tokenString := c.Get("Authorization")
	tokenValue := strings.Split(tokenString, " ")
	if tokenValue[0] != "Bearer" {
		return nil, errors.New("token 错误")
	}
	token := tokenValue[1]
	j := NewJWT()
	claims, err := j.ParseToken(token)
	if err != nil {
		global.LOG.Error("从fiber的Context中获取从jwt解析信息失败, 请检查请求头是否存在token且claims是否为规定结构")
	}
	return claims, err
}

// 从fiber的Context中获取从jwt解析出来的用户ID
func GetUserID(c *fiber.Ctx) uint {
	var claims = c.Locals("claims")
	waitUse := claims.(*systemReq.CustomClaims)
	if waitUse.BaseClaims.ID == 0 {
		if cl, err := GetClaims(c); err != nil {
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
