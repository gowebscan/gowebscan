package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/addtasklist"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/example"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/smalltools"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/vuldetection"
)

type ApiGroup struct {
	SystemApiGroup      system.ApiGroup
	ExampleApiGroup     example.ApiGroup
	AddTaskListApiGroup addtasklist.ApiGroup
	AddVulApiGroup      vuldetection.ApiGroup
	SmallToolsApiGroup  smalltools.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
