package response

import (
	"github.com/ykstudy/stock/model/system/request"
)

type PolicyPathResponse struct {
	Paths []request.CasbinInfo `json:"paths"`
}
