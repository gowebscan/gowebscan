package vuldetection

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type AddVulRouter struct {
}

// 初始化  路由信息
func (s *AddVulRouter) InitAddVulRouter(Router *gin.RouterGroup) {
	AddVulRouter := Router.Group("vuldetection").Use(middleware.OperationRecord())

	var AddVulApi = v1.ApiGroupApp.AddVulApiGroup.AddVulApi
	{
		AddVulRouter.POST("addvul", AddVulApi.AddtaskApi)
	}
}
