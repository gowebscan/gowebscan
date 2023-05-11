<template>
  <div class="page">
    <div class="tit">
      <h1>shellcode混淆</h1>
    </div>
    <div class="main">
      <el-form :model="ConFoundData">
        <div class="input">
          <p>INPUT:</p>
          <div class="inshell">
            <el-input v-model="ConFoundData.shellcode" :autosize="{ minRows: 6, maxRows: 10 }" type="textarea" placeholder="请输入shellcode"/>
          </div>
        </div>
        <div class="goshell">
          <el-form-item label="混淆模式：" prop="ConFoundRadio">
          <el-checkbox-group v-model="ConFoundData.checkList">
            <el-checkbox label="ADD" />
            <el-checkbox label="XOR" />
            <el-checkbox label="NOT" />
            <el-checkbox label="SUB" disabled />
            <el-checkbox label="NEG" disabled />
            <el-checkbox label="SHIFT" disabled />
            <el-checkbox label="MORPH" disabled />
          </el-checkbox-group>
        </el-form-item>
        </div>
        <div class="gogoshell">
          <el-button type="primary" round @click="ConFound">开始混淆</el-button>
        </div>
      </el-form>
      <div class="output">
        <p>OUTPUT:</p>
        <el-empty description="未混淆" />
      </div>
    </div>
  </div>
</template>

<style>
.page{
  width: 100%;
}
.tit{
  padding: 15px 0 0 30px;
  border-bottom: 1px solid rgba(0,0,0,.09);
  margin-bottom: 20px;
}
.main{
  width: 90%;
  /* background-color: aqua; */
}
.inshell{
  border: 1px solid rgba(0, 0, 0, 0);
  border-radius: 3px;
  margin-top: 20px;
}
.goshell {
    margin: 0 auto;
    margin-top: 20px;
}
.gogoshell {
    width: 50px;
    margin: 0 auto;
    margin-top: 20px;
}
.output{
  margin: 20px 0 10px 0;
}
</style>
<script setup>
import { reactive, ref } from 'vue'
import { ElNotification } from 'element-plus'
import { CodeConFound } from '@/api/confound'

const ConFoundData = reactive({
  shellcode: "",
  checkList: [],
});

const ConFound = async() => {
  
  if (ConFoundData.shellcode !== "" && ConFoundData.checkList !== ""){
    ElNotification({
        title: '混淆成功!',
        message: '在页面下方查看混淆后的代码!',
        type: 'success',
      })
      //开始任务
      const res = await CodeConFound({
        ConFoundData
      })
  } else {
      ElNotification({
      title: "未找到相关杀软进程！",
      message: "",
      type: "info",
    });
  }
}

</script>
  