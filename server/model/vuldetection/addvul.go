package vuldetection

// 读取响应 结构体
type VulRequestBody struct {
	VulFormRef map[string]interface{} `json:"VulFormRef"`
}
