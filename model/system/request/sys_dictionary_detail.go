package request

import (
	"github.com/ykstudy/stock/model/common/request"
	"github.com/ykstudy/stock/model/system"
)

type SysDictionaryDetailSearch struct {
	system.SysDictionaryDetail
	request.PageInfo
}
