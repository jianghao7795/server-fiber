package mobile

import (
	global "server-fiber/model"
	"server-fiber/model/mobile"
)

type MobileRegisterService struct{}

func (*MobileRegisterService) Register(data mobile.Register) error {
	return global.DB.Create(&data).Error
}
