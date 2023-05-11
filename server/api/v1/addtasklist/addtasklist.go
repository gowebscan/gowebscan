package addtasklist

import (
	"encoding/json"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/model/addtasklist"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

type AddTaskListApi struct{}

// AddTaskListApi
// @Tags      AddTaskListApi
// @Summary   获取前端参数
// @Produce   application/json
// @Success   {"ruleFormRef":{"TaskName":"Hello","TaskRadio":"单一目标","TaskTarget":"efwf","TaskCont":["子域名","POC探测"]}}

func (m *AddTaskListApi) AddtaskApi(c *gin.Context) {
	fmt.Println("打开了API操作")
	// 读取POST请求的Body数据
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		// 错误处理
	}
	// 解析POST请求的JSON数据到结构体中
	var requestBody addtasklist.RequestBody
	err = json.Unmarshal(body, &requestBody)
	if err != nil {
		// 错误处理
	}

	go func() {
		response.Ok(c)
	}()
	go InfoScanService.FofaInfoScan(requestBody)

}
