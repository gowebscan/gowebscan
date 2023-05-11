package addtasklist

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type AddTaskListRouter struct {
}

// 初始化  路由信息
func (s *AddTaskListRouter) InitAddtasklistRouter(Router *gin.RouterGroup) {
	AddTaskListRouter := Router.Group("InfoScan").Use(middleware.OperationRecord())

	var AddTaskListApi = v1.ApiGroupApp.AddTaskListApiGroup.AddTaskListApi
	{
		AddTaskListRouter.POST("addtasklist", AddTaskListApi.AddtaskApi)
		//AddTaskListRouter.POST("creattask", AddTaskListApi.CreateTaskApi)
	}
}
