package app

import (
	"server-fiber/global"
	"server-fiber/model/app"
)

type BaseMessageService struct{}

/**
 * @description: 创建baseMessage
 * @param {app.BaseMessage} baseMessage
 * @return {*}
 */
func (*BaseMessageService) CreateBaseMessage(baseMessage *app.BaseMessage) (err error) {
	err = global.DB.Create(baseMessage).Error
	return
}

/**
 * @description: 更新baseMessage
 * @param {app.BaseMessage} baseMessage
 * @return {*}
 */
func (*BaseMessageService) UpdateBaseMessage(baseMessage *app.BaseMessage) (err error) {
	err = global.DB.Save(baseMessage).Error
	return
}

/**
 * @description: 获取baseMessage
 * @return {*}
 */
func (*BaseMessageService) FindBaseMessage(id uint) (app.BaseMessage, error) {
	var base app.BaseMessage
	err := global.DB.Where("user_id = ?", id).First(&base).Error
	return base, err
}
