package addtasklist

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
)

type RequestBody struct {
	RuleFormRef map[string]interface{} `json:"ruleFormRef"`
}

// 创建 任务的结构体
type TaskInfo struct {
	global.GVA_MODEL
	TaskName           string         `json:"TaskName" form:"TaskName" gorm:"comment:任务名"`                        // 任务名
	TaskNameData       string         `json:"TaskNameData" form:"customerPhoneData" gorm:"comment:任务目标"`          // 任务目标
	SysUserID          uint           `json:"sysUserId" form:"sysUserId" gorm:"comment:管理ID"`                     // 管理ID
	SysUserAuthorityID uint           `json:"sysUserAuthorityID" form:"sysUserAuthorityID" gorm:"comment:管理角色ID"` // 管理角色ID
	SysUser            system.SysUser `json:"sysUser" form:"sysUser" gorm:"comment:管理详情"`                         // 管理详情
}

// TableName Test 表名
//func (TaskForm) TableName() string {
//	return "test"
//}
