import service from '@/utils/request'

// @Tags api
// @Summary 添加信息收集任务
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelInterface.PageInfo true "添加信息收集任务"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/getApiList [post]
// {
//  page     int
//	pageSize int
// }
export const AddTaskInfo = (data) => {
  return service({
    url: '/InfoScan/addtasklist',
    method: 'post',
    data
  })
}
// @Tags Api
// @Summary 删除客户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dbModel.ExaCustomer true "删除客户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /customer/customer [post]
export const CreateTask = (data) => {
  return service({
    url: '/InfoScan/creattask',
    method: 'post',
    data
  })
}