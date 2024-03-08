package request

import (
	"server-fiber/model/common/request"
	"server-fiber/model/system"
)

type SysDictionarySearch struct {
	system.SysDictionary
	request.PageInfo
}
