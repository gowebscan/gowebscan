<template>
  <div>
    <el-row :gutter="10">
      <el-col :span="12">
        <el-card>
          <template #header>
            <el-divider><h3>GoWabScan</h3></el-divider>
          </template>
          <div>
            <el-row>
              <el-col :span="8" :offset="8">
                <a href="#">
                  <img
                    class="org-img dom-center"
                    src="@/assets/logo.png"
                    alt="gin-vue-admin"
                  >
                </a>
              </el-col>
            </el-row>
            <el-row :gutter="10">
              <el-col :span="8">
                <a href="#">
                  <img
                    class="dom-center"
                    src="https://img.shields.io/github/watchers/flipped-aurora/gin-vue-admin.svg?label=Watch"
                    alt=""
                  >
                </a>
              </el-col>
              <el-col :span="8">
                <a href="#">
                  <img
                    class="dom-center"
                    src="https://img.shields.io/github/stars/flipped-aurora/gin-vue-admin.svg?style=social"
                    alt=""
                  >
                </a>
              </el-col>
              <el-col :span="8">
                <a href="#">
                  <img
                    class="dom-center"
                    src="https://img.shields.io/github/forks/flipped-aurora/gin-vue-admin.svg?label=Fork"
                    alt=""
                  >
                </a>
              </el-col>
            </el-row>
          </div>
        </el-card>
        <el-card style="margin-top: 20px">
          <template #header>
            <div><h3>GoWabScan平台网站的使用免责声明</h3></div>
          </template>
          <div>
            <el-row>
              <el-col :span="20" :offset="2">
    <p>&nbsp;&nbsp;&nbsp;&nbsp;GoWabScan渗透攻防平台 提醒您：在使用GoWebScan渗透攻防平台前，请您务必仔细阅读并透彻理解本声明。您可以选择不使用GoWebScan平台 ，但如果您使用GoWebScan平台 ，您的使用行为将被视为对本声明全部内容的认可。</p>
    <ol>
      <li>&nbsp;&nbsp;&nbsp;&nbsp;GoWabScan渗透攻防平台目前进行信息搜集的数据，目前都是在网上公开的信息数据，从未对他人隐私信息数据进行搜集及侵犯，系统会以非人工检索方式自动检索到第三方网页的链接，以便用户能够找到和使用第三方网页上各种文档、资料等内容；</li>
      <li>&nbsp;&nbsp;&nbsp;&nbsp;GoWabScan渗透攻防平台是为了更好地服务我们的用户，并不代表支持用户将利用本平台搜索到的信息用于任何用途。在使用GoWabScan渗透攻防平台时应当自行保证其使用目的及方式合法合规，不得从事危害国家安全、破坏社会稳定、扰乱社会秩序、侵犯他人合法权益的活动；</li>
      <li>&nbsp;&nbsp;&nbsp;&nbsp;在法律允许的范围内，GoWabScan渗透攻防平台在此声明，不承担用户或任何人士就使用GoWabScan渗透攻防平台所引致的任何直接、间接、附带、从属、特殊、惩罚性或惩戒性的损害赔偿（包括但不限于收益、预期利润的损失或失去的业务、未实现预期的节省）；</li>
      <li>&nbsp;&nbsp;&nbsp;&nbsp;凡以任何方式登陆GoWabScan渗透攻防平台，视为自愿接受本声明的约束: 不得利用本网站以任何方式直接或者间接的从事违反中国法律、国际公约以及社会公德的行为，且访问者应当恪守下述承诺（如若违反由用户自行承担相关责任）：</li>
      <p>(1) 传输和利用信息符合中国法律、国际公约的规定、符合公序良俗;</p>
      <p>(2) 不将本网站以及与之相关的网络服务用作非法用途以及非正当用途;</p>
      <p>(3) 不干扰和扰乱本网站以及与之相关的网络服务;</p>
      <p>(4) 遵守与本网站以及与之相关的网络服务的协议、规定、程序和惯例等。</p>
    </ol>
              </el-col>
            </el-row>
          </div>
        </el-card>
      </el-col>
      <el-col :span="12">
        <el-card>
          <template #header>
            <div><h3>更新记录</h3></div>
          </template>
          <div>
            2023-05-01 Version 1.0 : 初步完成各部分功能设计
          </div>
          <div>
            2023-05-01 Version 1.0 : 初步完成各部分功能设计
          </div>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script>
export default {
  name: 'About',
}
</script>

<script setup>


</script>

<style scoped>
.load-more {
  margin-left: 120px;
}

.avatar-img {
  float: left;
  height: 40px;
  width: 40px;
  border-radius: 50%;
  -webkit-border-radius: 50%;
  -moz-border-radius: 50%;
  margin-top: 15px;
}

.org-img {
  height: 150px;
  width: 150px;
}

.author-name {
  float: left;
  line-height: 65px !important;
  margin-left: 10px;
  color: darkblue;
  line-height: 100px;
  font-family: "Lucida Sans", "Lucida Sans Regular", "Lucida Grande",
    "Lucida Sans Unicode", Geneva, Verdana, sans-serif;
}

.dom-center {
  margin-left: 50%;
  transform: translateX(-50%);
}
</style>
