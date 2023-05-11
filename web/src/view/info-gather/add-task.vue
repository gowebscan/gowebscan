<template>
  <div class="page">
    <div class="title">资产信息收集</div>
    <div class="main">
      <el-form ref="ruleFormRef" :model="form" :rules="rules" label-width="120px" class="demo-ruleForm" :size="formSize" status-icon>
        <el-form-item label="任务名称：" prop="TaskName">
          <el-input v-model="form.TaskName" />
        </el-form-item>
        
        <el-form-item label="输入目标：" prop="TaskRadio">
          <el-radio-group v-model="form.TaskRadio">
            <el-radio label="单一目标" />
            <el-radio label="多个目标" />
          </el-radio-group>
        </el-form-item>

        <el-form-item label=" " prop="TaskTarget">
          <el-input v-model="form.TaskTarget" type="textarea" />
        </el-form-item>

        <el-form-item label="选择搜集内容 : " prop="TaskCont">
          <el-checkbox-group v-model="form.TaskCont">
            <el-checkbox label="子域名" name="TaskCont" />
            <el-checkbox label="IP/IP段" name="TaskCont" />
            <el-checkbox label="指纹识别" name="TaskCont" />
            <el-checkbox label="POC探测" name="TaskCont" />
          </el-checkbox-group>
        </el-form-item>

        <el-form-item class="btn">
          <el-button type="primary" @click="submitForm(ruleFormRef)">
            开始收集
          </el-button>
          <el-button @click="resetForm(ruleFormRef)">清空</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>
<style>
.page{
  background-color: #fff;
  width: 100%;
}
.main{
  width: 50%;  
  margin: 0 auto;
  padding: 30px 0 100px 0;
}
.title{
    padding: 16px 0 15px 32px;
    margin: 10px;
    border-bottom: 1px solid rgba(0,0,0,.09);
    font-style: normal;
    font-weight: 500;
    font-size: 16px;
    line-height: 24px;
}
.btn{
  float: right;
}
</style>

<script lang="ts" setup>
import { ElNotification } from 'element-plus'
import { reactive, ref } from 'vue'
import type { FormInstance, FormRules } from 'element-plus'
import { useRouter } from 'vue-router'
import { AddTaskInfo } from '@/api/addtask'
import { createExaCustomer } from '@/api/customer'

const router = useRouter()
const formSize = ref('default')
const ruleFormRef = ref<FormInstance>()
const form = reactive({
  TaskName: '',
  TaskRadio: '',
  TaskTarget: '',
  TaskCont: [],
})
var customerName = ''
var customerPhoneData = ''
var addtask = ''

const rules = reactive<FormRules>({
  TaskName: [
    { required: true, message: '请输入任务名', trigger: 'blur' },
    { min: 4, max: 15, message: '任务名长度在4-15个字符之间', trigger: 'blur' },
  ],
  TaskCont: [
    {
      type: 'array',
      required: true,
      message: '请选择收集的内容',
      trigger: 'change',
    },
  ],
  TaskRadio: [
    {
      required: true,
      message: '请选择收集目标类型',
      trigger: 'change',
    },
  ],
  TaskTarget: [
    { required: true, message: '请输入目标', trigger: 'blur' },
  ],
})

const submitForm = async (formEl: FormInstance | undefined) => {
  if (!formEl) return
  await formEl.validate(async (valid, fields) => {
    if (valid) {
      console.log('开始收集!')
      console.log(form)
      ElNotification({
        title: '开始收集!',
        message: '成功添加信息收集任务，开始收集!',
        type: 'success',
      })
      //开始任务
      const res = await AddTaskInfo({
        ruleFormRef: form
      })
      //添加任务
      customerName = form.TaskName
      customerPhoneData = form.TaskTarget
      // {"customerName":"safasasf","customerPhoneData":"asdfasf"}
      addtask = "{\"customerName\":\"" + customerName + "\",\"" + "customerPhoneData\":\"" + customerPhoneData + "\"}"
      console.log(addtask)
      const res1 = await createExaCustomer(addtask)
      if (res.code = "0" ){
        console.log("跳转url")
        router.push({ name: 'tasklist' })
      }else{
        ElNotification({
        title: '出现状态故障!',
        message: '故障请联系管理员!',
        type: 'error',
      })
      }
    } else {
      console.log('收集错误', fields)
    }
  })
}

const resetForm = (formEl: FormInstance | undefined) => {
  if (!formEl) return
  formEl.resetFields()
}

const options = Array.from({ length: 10000 }).map((_, idx) => ({
  value: `${idx + 1}`,
  label: `${idx + 1}`,
}))
</script>
