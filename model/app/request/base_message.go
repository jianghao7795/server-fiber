package request

import (
	"server-fiber/model/app"
	"server-fiber/model/common/request"
)

type BaseMessageBody struct {
	app.BaseMessage
	request.PageInfo
}
