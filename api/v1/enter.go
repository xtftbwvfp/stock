package v1

import (
	"github.com/ykstudy/stock/api/v1/example"
	"github.com/ykstudy/stock/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup  system.ApiGroup
	ExampleApiGroup example.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
