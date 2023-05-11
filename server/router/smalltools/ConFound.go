package smalltools

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ConFoundRouter struct {
}

// 初始化  路由信息
func (s *ConFoundRouter) InitConFoundRouter(Router *gin.RouterGroup) {
	ConFoundRouter := Router.Group("smalltools").Use(middleware.OperationRecord())

	var ConFoundApi = v1.ApiGroupApp.SmallToolsApiGroup.ConFoundApi
	{
		ConFoundRouter.POST("CodeConFound", ConFoundApi.CodeFoundApi)
	}
}
