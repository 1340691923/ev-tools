<template>
  <div class="app-container">

    <div class="search-container">
      <el-form :inline="true">
        <el-form-item label="行为">
          <el-select
              style="width:200px"
              v-model="input.actions"
              reserve-keyword
              collapse-tags
              :placeholder="$t('行为')"
              clearable
              multiple
              filterable
              @change="searchByFilter()"
          >
            <el-option v-for="(action, index) in actions" :key="index" :label="action" :value="action" />
          </el-select>
        </el-form-item>
        <el-form-item label="任务ID">
          <el-select
              style="width:200px"
              v-model="input.task_id"
              reserve-keyword
              collapse-tags
              :placeholder="$t('任务ID')"
              clearable
              multiple
              filterable
              @change="searchByFilter()"
          >
            <el-option v-for="(taskID, index) in taskIdList" :key="index" :label="taskID" :value="taskID" />
          </el-select>
        </el-form-item>
        <el-form-item label="节点ID">
          <el-select
              style="width:200px"
              v-model="input.node_id"
              reserve-keyword
              collapse-tags
              :placeholder="$t('节点ID')"
              clearable
              multiple
              filterable
              @change="searchByFilter()"
          >
            <el-option v-for="(node, index) in nodeIdList" :key="index" :label="node" :value="node" />
          </el-select>
        </el-form-item>
        <el-form-item label="父任务ID">
          <el-select
              style="width:200px"
              v-model="input.parent_task_id"
              :placeholder="$t('父任务ID')"
              clearable
              filterable
              @change="searchByFilter()"
          >
            <el-option v-for="(taskID, index) in parentTaskIdList" :key="index" :label="taskID" :value="taskID" />
          </el-select>
        </el-form-item>
        <el-form-item>
        </el-form-item>
        <el-form-item>
          <el-button

              type="primary"
              @click="search(true)"
          >{{ $t('搜索') }}
          </el-button>
        </el-form-item>
        <el-form-item>
          <el-button

              type="success"
              @click="refresh"
          >{{ $t('刷新') }}
          </el-button>
        </el-form-item>
      </el-form>

    </div>
    <el-table
        :data="tableData"
    >
      <el-table-column
          :label="$t('序号')"
          align="center"
          fixed
          width="50"
      >
        <template #default="scope">
          {{ scope.$index + 1 }}
        </template>
      </el-table-column>
      <el-table-column align="center" :label="$t('行为')" prop="action" width="200" />
      <el-table-column align="center" :label="$t('任务详细ID')" prop="taskID" width="300" />
      <el-table-column align="center" :label="$t('任务ID')" prop="id" width="100" />

      <el-table-column align="center" :label="$t('节点')" prop="node" width="200" />

      <el-table-column align="center" :label="$t('父任务ID')" prop="parent_task_id" width="300">
        <template #default="scope">
          <div>{{ scope.row.parent_task_id }}</div>
        </template>
      </el-table-column>
      <el-table-column align="center" :label="$t('开始详细时间')" width="180">
        <template #default="scope">
          <div>{{ timestampToTime(scope.row.start_time_in_millis) }}</div>
        </template>
      </el-table-column>
      <el-table-column align="center" :label="$t('状态')" width="300">
        <template #default="scope">

          <el-popover
              placement="top-start"
              :title="$t('描述')"
              width="600"
              trigger="hover"
          >

            <div>{{ JSON.stringify(scope.row.status) }}</div>
            <span slot="reference">{{ JSON.stringify(scope.row.status).substr(0, 50) + "..." }}</span>

          </el-popover>
        </template>
      </el-table-column>
      <el-table-column align="center" :label="$t('运行时间（秒）')" width="180">
        <template #default="scope">
          <div>{{ Number(scope.row.running_time_in_nanos / 1000000000) }}</div>
        </template>
      </el-table-column>

      <el-table-column align="center" :label="$t('是否可撤销')" prop="cancellable" width="100">
        <template #default="scope">
          <div>{{ scope.row.cancellable }}</div>
        </template>
      </el-table-column>

      <el-table-column align="center" :label="$t('描述')" prop="description" width="300">
        <template #default="scope">

          <el-popover
              placement="top-start"
              :title="$t('描述')"
              width="600"
              trigger="hover"
          >

            <div>{{ scope.row.description }}</div>
            <span slot="reference">{{ scope.row.description.substr(0, 50) + "..." }}</span>
          </el-popover>
        </template>
      </el-table-column>
      <el-table-column align="center" :label="$t('开始时间')" prop="start_time" width="100" />
      <el-table-column align="center" :label="$t('开始时间（毫秒）')" prop="start_time_in_millis" width="150" />

      <el-table-column align="center" :label="$t('类型')" prop="type" width="100" />

      <el-table-column align="center" :label="$t('操作')" fixed="right" width="350">
        <template #default="scope">
          <el-button-group>
            <el-button

                type="success"

                @click="look(scope.row)"
            >{{ $t('查看') }}
            </el-button>

            <el-button

                type="danger"

                @click="cancelActionHandler(scope.row.taskID)"
            >{{ $t('取消') }}
            </el-button>

          </el-button-group>

        </template>
      </el-table-column>
    </el-table>

    <el-drawer
        ref="drawer"
        :title="$t('任务详细信息')"
        :before-close="drawerHandleClose"
        v-model="drawerShow"
        direction="rtl"
        close-on-press-escape
        destroy-on-close
        size="80%"
    >

      <json-editor
          v-model:value="taskDetail"
          styles="width: 100%"
          :read="true"
          :title="$t('任务详细信息')"
      />
    </el-drawer>

  </div>
</template>

<script setup>
import {ref, onMounted, getCurrentInstance} from 'vue'
import Moment from 'moment'
import { CancelAction, ListAction } from '@/api/es-task'
import JsonEditor from '@/components/JsonEditor/index.vue'

import {ElMessage} from "element-plus";

const ctx = getCurrentInstance().appContext.config.globalProperties
import {sdk} from "@/plugin_sdk/sdk";

const taskDetail = ref('')
const loading = ref(false)
const drawerShow = ref(false)
const tableData = ref([])
const taskIdList = ref([])
const parentTaskIdList = ref([])
const nodeIdList = ref([])
const actions = ref([])
const input = ref({
  task_id: [],
  node_id: [],
  parent_task_id: '',
  actions: []
})

const inArray = (search, array) => {
  return array.includes(search)
}

const searchByFilter = () => {
  const filteredData = []
  let flag = false

  if (input.value.parent_task_id !== '') {
    flag = true
    for (const data of tableData.value) {
      if (data.parent_task_id === input.value.parent_task_id) {
        filteredData.push(data)
      }
    }
  }

  if (input.value.actions.length > 0) {
    flag = true
    for (const data of tableData.value) {
      if (inArray(data.action, input.value.actions)) {
        filteredData.push(data)
      }
    }
  }

  if (input.value.node_id.length > 0) {
    flag = true
    for (const data of tableData.value) {
      if (inArray(data.node, input.value.node_id)) {
        filteredData.push(data)
      }
    }
  }

  if (input.value.task_id.length > 0) {
    flag = true
    for (const data of tableData.value) {
      if (inArray(data.taskID, input.value.task_id)) {
        filteredData.push(data)
      }
    }
  }

  if (flag) {
    tableData.value = filteredData
  } else {
    search()
  }
}

const cancelActionHandler = async (taskid) => {
  const inputParams = {
    es_connect: sdk.GetSelectEsConnID(),
    task_id: taskid
  }

  const { data, code, msg } = await CancelAction(inputParams)
  if (code !== 0) {
    ElMessage({
      type: 'error',
      message: msg
    })
    return
  } else {
    setTimeout(() => {
      search()
    }, 3000)
    ElMessage({
      type: 'success',
      message: msg
    })
  }
}

const look = (obj) => {
  console.log(obj)
  taskDetail.value = JSON.stringify(obj, null, '\t')
  drawerShow.value = true
}

const drawerHandleClose = (done) => {
  done()
}

const timestampToTime = (timeStamp) => {
  const stamp = new Date(timeStamp)
  return Moment(stamp).format('YYYY-MM-DD HH:mm:ss')
}

const search = async (clear = false) => {
  if (clear) {
    input.value = {
      task_id: [],
      node_id: [],
      parent_task_id: '',
      actions: []
    }
  }

  const inputParams = {
    es_connect: sdk.GetSelectEsConnID()
  }

  const { data, code, msg } = await ListAction(inputParams)

  if (code !== 0) {
    ElMessage({
      type: 'error',
      message: msg
    })
    return
  } else {
    ElMessage({
      type: 'success',
      message: msg
    })

    tableData.value = []
    taskIdList.value = []
    parentTaskIdList.value = []
    nodeIdList.value = []

    const nodeSet = new Set()
    const actionsSet = new Set()

    for (const taskId in data) {
      const v = data[taskId]
      v.taskID = taskId
      tableData.value.push(v)
      taskIdList.value.push(taskId)
      if (v.parent_task_id !== '') {
        parentTaskIdList.value.push(v.parent_task_id)
      }
      nodeSet.add(v.node)
      actionsSet.add(v.action)
    }

    nodeIdList.value = Array.from(nodeSet)
    actions.value = Array.from(actionsSet)
  }
}

const refresh = async () => {
  await search()
  searchByFilter()
}

onMounted(() => {
  console.log(123)
  search()
})
</script>

<style scoped>

</style>
