import service from '@/utils/request'

// @Tags api
// @Summary 添加扫描
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
export const AddVul = (data) => {
  return service({
    url: '/vuldetection/addvul',
    method: 'post',
    data
  })
}