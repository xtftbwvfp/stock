package initialize

import (
	_ "github.com/ykstudy/stock/source/example"
	_ "github.com/ykstudy/stock/source/system"
)

func init() {
	// do nothing,only import source package so that inits can be registered
}
