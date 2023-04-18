package service

import (
	"github.com/ykstudy/stock/service/example"
	"github.com/ykstudy/stock/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup  system.ServiceGroup
	ExampleServiceGroup example.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
