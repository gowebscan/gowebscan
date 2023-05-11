package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/infoscan"
	"github.com/flipped-aurora/gin-vue-admin/server/service/smalltools"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/flipped-aurora/gin-vue-admin/server/service/vuldetection"
)

type ServiceGroup struct {
	SystemServiceGroup     system.ServiceGroup
	ExampleServiceGroup    example.ServiceGroup
	InfoScanServiceGroup   infoscan.ServiceGroup
	AddVulServiceGroup     vuldetection.ServiceGroup
	SmallToolsServiceGroup smalltools.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
