package request

import (
	"github.com/ykstudy/stock/model/common/request"
	"github.com/ykstudy/stock/model/system"
)

type SysDictionarySearch struct {
	system.SysDictionary
	request.PageInfo
}
