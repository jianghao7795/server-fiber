package mobile

import (
	"server-fiber/global"
	"server-fiber/model/mobile"
)

type MobileRegisterService struct{}

func (*MobileRegisterService) Register(data mobile.Register) error {
	return global.DB.Create(&data).Error
}
