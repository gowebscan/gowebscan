package smalltools

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type SmallToolsRouter struct {
}

// 初始化  路由信息
func (s *SmallToolsRouter) InitSmallToolsRouter(Router *gin.RouterGroup) {
	SmallToolsRouter := Router.Group("smalltools").Use(middleware.OperationRecord())

	var KillApi = v1.ApiGroupApp.SmallToolsApiGroup.KillApi
	{
		SmallToolsRouter.POST("killPorcesses", KillApi.GankApi)
	}
}
