<template>
    <div>
      <warning-bar title="当前正在执行的任务列表" />
      <div class="gva-table-box">
        <el-table
          ref="multipleTable"
          :data="tableData"
          style="width: 100%"
          tooltip-effect="dark"
          row-key="ID"
        >
          <!-- <el-table-column type="selection" width="55" /> -->
          <el-table-column align="left" label="收集时间" width="180">
            <template #default="scope">
              <span>{{ formatDate(scope.row.CreatedAt) }}</span>
            </template>
          </el-table-column>
          <el-table-column align="left" label="任务名称" prop="customerName" width="250" />
          <el-table-column align="left" label="目标" prop="customerPhoneData" width="250" />
          <el-table-column align="left" label="收集者" prop="sysUserId" width="120" />
          <el-table-column align="left" label="任务状态" min-width="160">
            <template #default="scope">
              <el-button type="primary" link icon="edit" @click="updateCustomer(scope.row)">进行中</el-button>
            </template>
          </el-table-column>
        </el-table>
        <div class="gva-pagination">
          <el-pagination
            :current-page="page"
            :page-size="pageSize"
            :page-sizes="[10, 30, 50, 100]"
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
  import {
    createExaCustomer,
    updateExaCustomer,
    deleteExaCustomer,
    getExaCustomer,
    getExaCustomerList
  } from '@/api/customer'
  import WarningBar from '@/components/warningBar/warningBar.vue'
  import { ref } from 'vue'
  import { ElMessage } from 'element-plus'
  import { formatDate } from '@/utils/format'
  
  const form = ref({
    customerName: '',
    customerPhoneData: ''
  })
  
  const page = ref(1)
  const total = ref(0)
  const pageSize = ref(10)
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
    const table = await getExaCustomerList({ page: page.value, pageSize: pageSize.value })
    if (table.code === 0) {
      tableData.value = table.data.list
      total.value = table.data.total
      page.value = table.data.page
      pageSize.value = table.data.pageSize
    }
  }
  
  getTableData()
  
  const dialogFormVisible = ref(false)
  const type = ref('')
  const updateCustomer = async(row) => {
    const res = await getExaCustomer({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
      form.value = res.data.customer
      dialogFormVisible.value = true
    }
  }

  </script>
  
  <script>
  
  export default {
    name: 'Customer'
  }
  </script>
  
  <style></style>
  