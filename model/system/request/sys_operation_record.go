package request

import (
	"github.com/ykstudy/stock/model/common/request"
	"github.com/ykstudy/stock/model/system"
)

type SysOperationRecordSearch struct {
	system.SysOperationRecord
	request.PageInfo
}
