package request

import (
	"github.com/wubainian/godaemon/model/common/request"
	"github.com/wubainian/godaemon/model/system"
)

type SysOperationRecordSearch struct {
	system.SysOperationRecord
	request.PageInfo
}
