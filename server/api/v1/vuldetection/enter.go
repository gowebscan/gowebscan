package vuldetection

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service"
)

type ApiGroup struct {
	AddVulApi
}

var (
	AddVulService = service.ServiceGroupApp.AddVulServiceGroup.AddVulService
)
