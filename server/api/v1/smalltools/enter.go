package smalltools

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	KillApi
	ConFoundApi
}

var (
	KillService     = service.ServiceGroupApp.SmallToolsServiceGroup.KillService
	ConFoundService = service.ServiceGroupApp.SmallToolsServiceGroup.ConFoundService
)
