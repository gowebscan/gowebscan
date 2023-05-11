<template>
    <div class="page">
      <div class="title">资产信息收集</div>
      <div class="main">
        <el-form ref="VulFormRef" :model="form" :rules="rules" label-width="120px" class="demo-ruleForm" :size="formSize" status-icon>
          <el-form-item label="任务名称：" prop="TaskName">
            <el-input v-model="form.TaskName" />
          </el-form-item>
          
          <el-form-item label="选择资产：" prop="TaskRadio">
            <el-radio-group v-model="form.TaskRadio">
              <el-radio label="已有资产" />
              <el-radio label="自定义添加" />
            </el-radio-group>
          </el-form-item>
  
          <el-form-item label=" " prop="TaskTarget">
            <el-input v-model="form.TaskTarget" type="textarea" />
          </el-form-item>

          <el-form-item label="扫描内容 : " prop="TaskRadioCon">
            <el-radio-group v-model="form.TaskRadioCon">
              <el-radio label="只识别指纹" />
              <el-radio label="识别出指纹在进行poc探测" />
              <el-radio label="不需要指纹盲打poc" />
            </el-radio-group>
          </el-form-item>
  
          <el-form-item class="btn">
            <el-button type="primary" @click="submitForm(VulFormRef)">
              开始扫描
            </el-button>
            <el-button @click="resetForm(VulFormRef)">清空</el-button>
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
  import { AddVul } from '@/api/addvul'

  const router = useRouter()
  const formSize = ref('default')
  const VulFormRef = ref<FormInstance>()
  const form = reactive({
    TaskName: '',
    TaskRadio: '',
    TaskTarget: '',
  })

  const rules = reactive<FormRules>({
    TaskName: [
      { required: true, message: '请输入任务名', trigger: 'blur' },
      { min: 4, max: 15, message: '任务名长度在4-15个字符之间', trigger: 'blur' },
    ],
    TaskCont: [
      {
        type: 'array',
        required: true,
        message: '请选择扫描的内容',
        trigger: 'change',
      },
    ],
    TaskRadio: [
      {
        required: true,
        message: '请选择扫描资产',
        trigger: 'change',
      },
    ],
    TaskRadioCon: [
      {
        required: true,
        message: '请选择扫描目标类型',
        trigger: 'change',
      },
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
        const res = await AddVul({
          VulFormRef: form
        })
      } else {
        console.log('扫描错误', fields)
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
