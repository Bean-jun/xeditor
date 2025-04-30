<template>
  <el-dialog v-model="dialogVisible" width="800">
    <el-table :data="gridData">
      <el-table-column prop="name" width="300" label="Python解释器">
        <template #default="scope">
          <div style="display: flex; align-items: center">
            <el-button type="success" style="width: 100%;">Python {{ scope.row.name }}</el-button>
          </div>
        </template>
      </el-table-column>
      <el-table-column prop="install" width="150" label="是否安装">
        <template #default="scope">
          <div style="display: flex; align-items: center">
            <el-tag type="success" v-if="scope.row.install">已安装</el-tag>
            <el-tag type="danger" v-else>未安装</el-tag>
          </div>
        </template>
      </el-table-column>
      <el-table-column prop="current" width="150" label="默认解释器">
        <template #default="scope">
          <div style="display: flex; align-items: center">
            <el-tag type="success" v-if="scope.row.current">默认解释器</el-tag>
          </div>
        </template>
      </el-table-column>
      <el-table-column fixed="right" label="操作" min-width="150">
        <template #default="scope">
          <el-button size="small" type="danger" v-if="scope.row.install">卸载</el-button size="samil">
          <el-button size="small" type="success" v-else>安装</el-button>
          <el-button type="success" size="small">设置默认</el-button>
        </template>
      </el-table-column>
    </el-table>
  </el-dialog>

  <div class="nav">
    <div class="editor-operation">
      <div class="cell" @click="dialogVisible = !dialogVisible">
        <el-button link>
          解释器设置
        </el-button>
      </div>
    </div>
    <div class="common-operation">
      <div class="cell" v-for="item in commonOperation" :key="item.key" @click="item.callback()">
        <!-- {{ item.title }} -->
        <Operate :color="item.color" width="15px" height="15px"></Operate>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive, computed, watch, onMounted } from 'vue';
import { ElMessageBox } from 'element-plus';
import { Timer } from '@element-plus/icons-vue';
import Operate from "../button/Operate.vue";
import {
  WindowMinimise,
  WindowToggleMaximise,
  Quit,
} from "../../../wailsjs/runtime";
import { GetInterpreterList } from "../../../wailsjs/go/controller/PythonController";


const gridData = ref([]);
function getInterpreter() {
  GetInterpreterList().then((e) => {
    const list = []
    for (let idx = 0; idx < e.list.length; idx++) {
      const element = e.list[idx];
      list.push({
        "name": element,
        "install": e.available.includes(element),
        "current": e.current === element,
      })
      gridData.value = list;
    }
  }
  )
}

const commonOperation = [
  {
    key: 1,
    title: "最小化",
    color: "green",
    callback: WindowMinimise,
  },
  {
    key: 2,
    title: "最大化",
    color: "yellow",
    callback: WindowToggleMaximise,
  },
  {
    key: 3,
    title: "关闭",
    color: "red",
    callback: Quit,
  },
];

const dialogVisible = ref(false);

onMounted(() => {
  getInterpreter()
})
</script>

<style lang="less" scoped>
.nav {
  color: white;
  background-color: #000000;
  height: 30px;
  --wails-draggable: drag;

  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: space-between;

  .editor-operation {
    display: flex;
    flex-direction: row;
    justify-content: flex-start;
    align-items: center;

    line-height: 100%;

    .cell {
      margin: 0 5px;
      padding: 1px;
      border-radius: 5px;
      cursor: pointer;
      color: white;

    }

    .cell:hover {
      background-color: red;
    }
  }

  .common-operation {
    display: flex;
    flex-direction: row;
    justify-content: flex-end;
    align-items: center;
    height: 100%;

    .cell {
      margin: 0 5px;
      cursor: pointer;
    }

    // .cell:hover {
    //     background-color: red;
    // }
  }
}
</style>
