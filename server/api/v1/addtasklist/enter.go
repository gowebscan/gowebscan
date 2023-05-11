package addtasklist

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service"
)

type ApiGroup struct {
	AddTaskListApi
}

var (
	InfoScanService = service.ServiceGroupApp.InfoScanServiceGroup.InfoScanService
)
