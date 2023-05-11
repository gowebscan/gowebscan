import service from '@/utils/request'

// @Tags api
// @Summary 代码混淆
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
export const CodeConFound = (data) => {
  return service({
    url: '/smalltools/CodeConFound',
    method: 'post',
    data
  })
}