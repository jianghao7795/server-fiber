package request

import (
	"server-fiber/model/common/request"
	"server-fiber/model/system"
)

type SysDictionaryDetailSearch struct {
	system.SysDictionaryDetail
	request.PageInfo
}
