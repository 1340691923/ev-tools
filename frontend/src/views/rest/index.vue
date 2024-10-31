<template>
  <div class="app-container" style="overflow-x:hidden;">

    <el-card style="margin-top: 10px" class="box-card">
      <el-tabs
          editable
          @edit="handleTabsEdit"
          class="curd-tab"
          v-model="editableTabsValue" type="card" closable @tab-remove="removeTab">
        <template #add-icon>
          <el-icon style="font-size: 14px;font-weight: bolder" >
            <Plus/>
          </el-icon>
        </template>
        <el-tab-pane
            :lazy="true"
            v-for="(item, index) in editableTabs"
            :key="item.uniqueId"
            :label="item.title"
            :name="Number(item.uniqueId)"
        >

          <tools
              :indices="sqlPointOut"
              :title="item.title"
              :query-data="queryData"
              :input="item.input"
              :unique-id="item.uniqueId"
              :sql-str="item.sqlStr"
              @saveData="saveData"
          />
        </el-tab-pane>
      </el-tabs>
    </el-card>
  </div>
</template>

<script setup>
import {ref, onMounted, onBeforeUnmount, getCurrentInstance} from 'vue'
import { clone } from '@/utils/index'
import { esPathKeyWords } from '@/utils/base-data'
import { ListAction } from '@/api/es-map'
import {ElMessageBox} from "element-plus";
import Tools from "@/views/rest/components/tools.vue";
const ctx = getCurrentInstance().appContext.config.globalProperties
import {sdk} from "@/plugin_sdk/sdk";

const editableTabsValue = ref(1)
const editableTabs = ref([])
const queryData = ref([])
const sqlPointOut = ref([])

onMounted(() => {
  mergeEsPathKeyWords()

  const savedEditableTabs = sessionStorage.getItem('editableTabs')
  const savedEditableTabsValue = sessionStorage.getItem('editableTabsValue')

  if (savedEditableTabsValue != null && savedEditableTabs != '' && savedEditableTabs != 'null') {
    editableTabsValue.value = Number(savedEditableTabsValue)
  } else {
    editableTabsValue.value = 1
  }

  if (savedEditableTabs != null && savedEditableTabs != '' && savedEditableTabs != 'null') {
    editableTabs.value = JSON.parse(savedEditableTabs)
  } else {
    editableTabs.value.push({
      title: '新窗口1',
      uniqueId: 1,
      input: {
        body: '{}',
        method: 'POST',
        path: ''
      },
      sqlStr: 'select * from '
    })
  }
})


const handleTabsEdit = (targetName,action) => {
  if (action === 'add') {
    addTab()
  }
}

onBeforeUnmount(() => {
  const editableTabsString = JSON.stringify(editableTabs.value)
  sessionStorage.setItem('editableTabs', editableTabsString)
  sessionStorage.setItem('editableTabsValue', editableTabsValue.value)
})

const saveData = (uniqueId, input, sqlStr, title) => {
  for (const editableTab of editableTabs.value) {
    if (editableTab.uniqueId === uniqueId) {
      editableTab.input = input
      editableTab.sqlStr = sqlStr
      editableTab.title = title
      break
    }
  }
}

const mergeEsPathKeyWords = async () => {
  const input = {
    es_connect: sdk.GetSelectEsConnID()
  }
  const res = await ListAction( input)
  if (res.code === 0) {
    const list = res.data.list
    const indices = Object.keys(list)
    queryData.value = []
    const clonedQueryData = clone(esPathKeyWords)
    for (const indice of indices) {
      sqlPointOut.value.push({ caption: indice, meta: indice, value: indice, score: 1 })
    }
    for (const esPathKeyWord of clonedQueryData) {
      if (esPathKeyWord.value.includes('{indices}')) {
        for (const indice of indices) {
          const mappings = Object.keys(list[indice]['mappings']) || ['{type}']
          const obj = {
            value: esPathKeyWord.value.replace(/{indices}/g, indice).replace(/{type}/g, mappings[0]),
            data: esPathKeyWord.data.replace(/{indices}/g, indice).replace(/{type}/g, mappings[0])
          }
          queryData.value.push(obj)
        }
      }
    }
    queryData.value.push(...clonedQueryData)
  }
}

const addTab = () => {
  ElMessageBox.prompt('请输入新窗口标题', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消'
  })
      .then(({ value }) => {
        if (!value) {
          value = `新窗口${editableTabs.value.length + 1}`
        }
        const timestamp = Date.now()
        editableTabs.value.push({
          title: value,
          uniqueId: timestamp,
          input: {
            body: '{}',
            method: 'POST',
            path: ''
          },
          sqlStr: 'select * from '
        })
        editableTabsValue.value = timestamp
      })
      .catch(err => {
        console.log('err', err)
      })
}

const removeTab = targetName => {



  let activeName = editableTabsValue.value
  const tabs = editableTabs.value
  if (activeName === targetName) {
    tabs.forEach((tab, index) => {
      if (tab.uniqueId === targetName) {
        const nextTab = tabs[index + 1] || tabs[index - 1]
        if (nextTab) {
          activeName = nextTab.uniqueId
        }
      }
    })
  }

  editableTabsValue.value = activeName
  editableTabs.value = tabs.filter(tab => tab.uniqueId !== targetName)
}
</script>

<style>
.curd-tab .el-tabs__new-tab{
  width:34px;
  height:34px
}
</style>
