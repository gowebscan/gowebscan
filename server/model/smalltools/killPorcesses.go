package smalltools

// 读取响应 结构体
type KillRequestBody struct {
	KillFormData map[string]interface{} `json:"killFormData"`
}

// TableName Test 表名
//func (Test) TableName() string {
//	return "test"
//}
