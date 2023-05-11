package smalltools

import (
	"encoding/json"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/smalltools"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

type KillApi struct{}

func (m *KillApi) GankApi(c *gin.Context) {
	fmt.Println("打开了杀软识别API操作")
	// 读取POST请求的Body数据
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		// 错误处理
	}
	// 解析POST请求的JSON数据到结构体中
	var killRequestBody smalltools.KillRequestBody
	err = json.Unmarshal(body, &killRequestBody)
	if err != nil {
		// 错误处理
	}
	resp := KillService.PorcessesScan(killRequestBody)
	fmt.Println(resp)
	// 返回resp信息
	response.OkWithData(resp, c)
}
