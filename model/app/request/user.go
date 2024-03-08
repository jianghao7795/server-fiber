package request

import (
	"server-fiber/model/app"
	"server-fiber/model/common/request"
)

type UserSearch struct {
	app.User
	request.PageInfo
}
