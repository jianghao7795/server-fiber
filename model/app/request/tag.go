package request

import (
	"server-fiber/model/app"
	"server-fiber/model/common/request"
)

type TagSearch struct {
	app.Tag
	request.PageInfo
}
