<template>
    <div v-loading.fullscreen.lock="fullscreenLoading">
      <div class="gva-table-box">
        <warning-bar
          title="信息收集的内容在此处下载。"
        />
        <el-table :data="tableData">
          <!-- <el-table-column align="left" label="预览" width="100">
            <template #default="scope">
              <CustomPic pic-type="file" :pic-src="scope.row.url" />
            </template>
          </el-table-column> -->
          <el-table-column align="left" label="日期" prop="UpdatedAt" width="180">
            <template #default="scope">
              <div>{{ formatDate(scope.row.UpdatedAt) }}</div>
            </template>
          </el-table-column>
          <el-table-column align="left" label="任务名称" prop="name" width="180">
            
          </el-table-column>
          <el-table-column align="left" label="链接" prop="url" min-width="300" />
          <el-table-column align="left" label="标签" prop="tag" width="100">
            <template #default="scope">
              <el-tag
                :type="scope.row.tag === 'xlsx' ? 'primary' : 'success'"
                disable-transitions
              >{{ scope.row.tag }}
              </el-tag>
            </template>
          </el-table-column>
          <el-table-column align="left" label="操作" width="160">
            <template #default="scope">
              <el-button icon="download" type="primary" link @click="downloadFile(scope.row)">下载</el-button>
            </template>
          </el-table-column>
        </el-table>
        <div class="gva-pagination">
          <el-pagination
            :current-page="page"
            :page-size="pageSize"
            :page-sizes="[10, 30, 50, 100]"
            :style="{ float: 'right', padding: '20px' }"
            :total="total"
            layout="total, sizes, prev, pager, next, jumper"
            @current-change="handleCurrentChange"
            @size-change="handleSizeChange"
          />
        </div>
      </div>
    </div>
  </template>
  
  <script setup>
  import { getFileList } from '@/api/fileUploadAndDownload'
  import { 
    downloadImage,
    downloadXLSX
} from '@/utils/downloadImg'
  import CustomPic from '@/components/customPic/index.vue'
  import { formatDate } from '@/utils/format'
  import WarningBar from '@/components/warningBar/warningBar.vue'
  
  import { ref } from 'vue'
  import { ElMessage, ElMessageBox } from 'element-plus'
  
  const path = ref(import.meta.env.VITE_BASE_API)
  
  const imageUrl = ref('')
  const imageCommon = ref('')
  
  const page = ref(1)
  const total = ref(0)
  const pageSize = ref(10)
  const search = ref({})
  const tableData = ref([])
  
  // 分页
  const handleSizeChange = (val) => {
    pageSize.value = val
    getTableData()
  }
  
  const handleCurrentChange = (val) => {
    page.value = val
    getTableData()
  }
  
  // 查询
  const getTableData = async() => {
    const table = await getFileList({ page: page.value, pageSize: pageSize.value, ...search.value })
    if (table.code === 0) {
      tableData.value = table.data.list
      total.value = table.data.total
      page.value = table.data.page
      pageSize.value = table.data.pageSize
    }
  }
  getTableData()
  
  
  const downloadFile = (row) => {
    if (row.url.indexOf('http://') > -1 || row.url.indexOf('https://') > -1) {
    //   http://localhost:8080/api/uploads/file/23176efb31ac3c0dc045da412916901a_20230424154702.xlsx
    // uploads/file/23176efb31ac3c0dc045da412916901a_20230424154702.xlsx 202304231811.xlsx
      // downloadImage(row.url, row.name)
      console.log("11")
      console.log(row.url, row.name)

    } else {
    //   debugger1
      console.log("22")
      console.log(row.url, row.name)
      // downloadImage(path.value + '/' + row.url, row.name) 
    downloadXLSX(row.url, row.name)
    }
  }
  
  /**
   * 编辑文件名或者备注
   * @param row
   * @returns {Promise<void>}
   */
  const editFileNameFunc = async(row) => {
    ElMessageBox.prompt('请输入文件名或者备注', '编辑', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      inputPattern: /\S/,
      inputErrorMessage: '不能为空',
      inputValue: row.name
    }).then(async({ value }) => {
      row.name = value
      // console.log(row)
      const res = await editFileName(row)
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '编辑成功!',
        })
        getTableData()
      }
    }).catch(() => {
      ElMessage({
        type: 'info',
        message: '取消修改'
      })
    })
  }
  </script>
  
  <script>
  
  export default {
    name: 'Upload',
  }
  </script>
  <style scoped>
  .name {
    cursor: pointer;
  }
  
  .upload-btn + .upload-btn {
    margin-left: 12px;
  }
  </style>
  