package vuldetection

import (
	"encoding/json"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/vuldetection"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

type AddVulApi struct{}

func (m *AddVulApi) AddtaskApi(c *gin.Context) {
	fmt.Println("打开了AddVulApi操作")
	// 读取POST请求的Body数据
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		// 错误处理
	}
	// 解析POST请求的JSON数据到结构体中
	var requestBody vuldetection.VulRequestBody
	err = json.Unmarshal(body, &requestBody)
	if err != nil {
		// 错误处理
	}
	fmt.Println(requestBody)
	//后台扫描前台跳转
	go func() {
		response.Ok(c)
	}()
	go AddVulService.AddVulScan(requestBody)

}
