package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/addtasklist"
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/smalltools"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
	"github.com/flipped-aurora/gin-vue-admin/server/router/vuldetection"
)

type RouterGroup struct {
	System           system.RouterGroup
	Example          example.RouterGroup
	AddTaskListGroup addtasklist.RouterGroup
	AddVulGroup      vuldetection.RouterGroup
	SmallToolsGroup  smalltools.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
