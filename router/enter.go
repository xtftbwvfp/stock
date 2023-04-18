package router

import (
	"github.com/ykstudy/stock/router/example"
	"github.com/ykstudy/stock/router/system"
)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
