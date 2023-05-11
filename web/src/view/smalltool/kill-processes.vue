<template>
    <div class="killporc">
      <div class="tit">
        <h1>杀软进程高亮</h1>
      </div>
        <div class="main">
            <div class="el-table el-table--fit el-table--striped el-table--border el-table--enable-row-hover el-table--mini list"
                element-loading-text="加载中...">
                <div>
                <el-table ref="filterTable" :data="tableData" style="width: 100%">
                    <el-table-column prop="pidname" label="杀软名称" sortable width="180"></el-table-column>
                    <el-table-column prop="pid" label="PID" width="180"></el-table-column>
                    <el-table-column prop="killcontents" label="描述" :formatter="formatter"></el-table-column>
                    <el-table-column prop="tag" label="标签" width="100" :filters="[{ text: '杀软', value: '杀软' }]" :filter-method="filterTag" filter-placement="bottom-end">
                    <template #default="{ row }">
                        <el-tag :type="row.tag === '家' ? 'primary' : 'success'" disable-transitions>{{ row.tag }}</el-tag>
                    </template>
                    </el-table-column>
                </el-table>
                </div>
                <div class="el-table__column-resize-proxy" style="display: none;">
                </div>
            </div>
            <!-- 进程输入 -->
            <div class="proc">
                <el-form :model="KillFormData">
                    <el-input type="textarea" :autosize="{ minRows: 14, maxRows: 16 }"
                        placeholder="执行 tasklist /SVC > tasklist.txt" v-model="KillFormData.killtextarea"></el-input>
                </el-form>
            </div>
            <div class="shibie">
                <el-button type="primary" round @click="OnDiscriminate">开始识别</el-button>
            </div>
        </div>
    </div>
</template>    
<style>
.killporc {
    background-color: #fff;
    width: 100%;
}
.tit{
  border-bottom: 1px solid rgba(0,0,0,.09);
}
.main {
    width: 95%;
    margin: 0 auto;
    padding: 20px 0;
}

.list {
    box-shadow: 1px 1px 1px 1px #efefef;
}

.proc {
    box-shadow: 1px 1px 1px 1px #efefef;
}

.shibie {
    width: 50px;
    margin: 0 auto;
    margin-top: 20px;
}
</style>

<script setup>
import { ElNotification } from "element-plus";
import { reactive, ref } from "vue";
import { killPorcesses } from "@/api/killPorcesses";

const KillFormData = reactive({
  killtextarea: "",
});

const tableData = ref([]);

const OnDiscriminate = async () => {
  if (KillFormData.killtextarea !== "") {
    // 清空表格数据
    tableData.value = [];

    ElNotification({
      title: "开始识别!",
      message: "正在识别杀软进程。。。",
      type: "success",
    });

    try {
      const res = await killPorcesses({ KillFormData });

      if (res.data && res.data.length > 0) {
        res.data.forEach((item) => {
          const json = JSON.parse(item);
          const { killname, killdescribe, pidname, pid } = json;
          tableData.value.push({
            pidname,
            pid,
            killcontents: killdescribe,
            tag: "杀软",
          });
        });
      } else {
        ElNotification({
          title: "未找到相关杀软进程！",
          message: "",
          type: "info",
        });
      }
    } catch (error) {
      console.error(error);

      ElNotification({
        title: "网络错误！",
        message: "请检查网络连接后重试",
        type: "error",
      });
    }
  } else {
    ElNotification({
      title: "识别失败",
      message: "请在cmd命令执行 “tasklist” 复制杀软进程 粘贴到文本框进行识别!",
      type: "error",
    });
  }
};

const formatter = (row, column) => {
  return row.killcontents;
};

const filterTag = (value, row) => {
  return row.tag === value;
};
</script>