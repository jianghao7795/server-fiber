package request

import (
	"server-fiber/model/common/request"
	"server-fiber/model/system"
)

type SysOperationRecordSearch struct {
	system.SysOperationRecord
	request.PageInfo
}
