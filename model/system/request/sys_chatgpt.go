package request

import (
	"github.com/ykstudy/stock/model/common/request"
	"github.com/ykstudy/stock/model/system"
)

type ChatGptRequest struct {
	system.ChatGpt
	request.PageInfo
}
