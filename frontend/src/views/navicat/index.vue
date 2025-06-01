<template>
  <div class="navicat-container">
    <div class="content-wrapper">

      <page-split class="page-split" :distribute="0.2" :lineThickness="3" :isVertical="true">
        <template v-slot:first>
          <div class="left-panel">
            <div class="search-container">
              <el-autocomplete
                  v-model="filterStr"
                  class="search-input"
                  clearable
                  :fetch-suggestions="querySearch"
                  :placeholder="$t('请输入索引名')"
              >
                <template #default="{ item }">
                  <span>{{ item.value }}</span>
                </template>
              </el-autocomplete>

              <el-menu v-loading="loadingMenu" class="index-menu">

                <template #default  >

                  <el-menu-item v-for="(v, index2) in getIndexList" :index="index2.toString()" @click="clickItem(v['index'])" class="menu-item">
                    <el-dropdown>
                      <span class="el-dropdown-link">
                        <el-icon v-if="v.health == 'red'" style="color: red" ><Grid /></el-icon>
                        <el-icon v-if="v.health == 'green'" style="color: #13ce66" ><Grid /></el-icon>
                        <el-icon v-if="v.health == 'yellow'" style="color: #ffba00" ><Grid /></el-icon>
                      </span>
                      <template #dropdown>
                        <el-dropdown-menu>
                          <el-dropdown-item
                              @click="deleteIndex(getIndexList[index2].k, getIndexList[index2].index)"
                          >
                            {{ $t('删除') }}
                          </el-dropdown-item>
                          <el-dropdown-item
                              @click.native="showIndexSettings = true"
                          >
                            {{ $t('结构') }}
                          </el-dropdown-item>
                        </el-dropdown-menu>
                      </template>
                    </el-dropdown>



                    <span slot="title">{{ v.index }}【{{ v.storeSize }}】</span>

                  </el-menu-item>
                </template>
              </el-menu>
            </div>
          </div>
        </template>
        <template v-slot:second>
          <div class="right-panel">

            <crud v-if="refreshTab"
                  :attr-map-prop="attrMap"
                  :event-attr-options-prop="eventAttrOptions"
                  :index-name="currentIndexName" />
          </div>
          <el-drawer
              ref="drawer"
              :title="$t('索引结构')"
              v-model="showIndexSettings"
              direction="rtl"
              close-on-press-escape
              destroy-on-close
              size="80%"
          >
            <div style="height: 95%;width: 100px;display: inline-block; height: 100%;vertical-align: top;width: 100%;">
              <el-tabs v-model="activeName" v-loading="tabLoading" type="border-card">
                <el-tab-pane :lazy="true" :label="$t('索引设置')" name="settings">
                  <json-editor
                      v-if="refreshTab"
                      v-model:value="currentSettingsJson"
                      height="720"
                      styles="width: 100%"
                      :read="true"
                      :title="currentIndexName"
                  />
                </el-tab-pane>
                <el-tab-pane :lazy="true" :label="$t('映射结构')" name="mappings">
                  <json-editor
                      v-if="refreshTab"
                      v-model:value="currentMappingJson"
                      height="720"
                      styles="width: 100%"
                      :read="true"
                      :title="currentIndexName"
                  />
                </el-tab-pane>
              </el-tabs>
            </div>
          </el-drawer>
        </template>
      </page-split>
    </div>
  </div>
</template>

<script setup>
import {ref, computed, onBeforeMount, nextTick, onMounted, onActivated} from 'vue'
import { ElMessageBox, ElMessage } from 'element-plus'

import { CatAction } from '@/api/es'
import { filterData } from '@/utils/table'
import { DeleteAction, GetSettingsAction } from '@/api/es-index'
import { ListAction } from '@/api/es-map'

import PageSplit  from "vue3-page-split";

import "vue3-page-split/dist/style.css";


import {sdk} from "@elasticview/plugin-sdk";
// Importing components
import JsonEditor from '@/components/JsonEditor/index.vue'
import Crud from '@/views/navicat/crud.vue'

// Setup instance

// Define reactive variables
const indexTishiList = ref([])
const currentMappings = ref({})
const tabLoading = ref(false)
const refreshTab = ref(true)
const activeName = ref('settings')
const loadingMenu = ref(false)
const indexList = ref([])
const filterStr = ref('')
const currentIndexName = ref('')
const currentSettings = ref({})
const eventAttrOptions = ref([])
const showIndexSettings = ref(false)
const attrMap = ref([])

// Computed properties
const getIndexList = computed(() => {
  if (filterStr.value === '') {
    return indexList.value
  }
  return filterData(indexList.value, filterStr.value)
})

const currentSettingsJson = computed(() => {
  return JSON.stringify(currentSettings.value, null, '\t')
})

const currentMappingJson = computed(() => {
  return JSON.stringify(currentMappings.value, null, '\t')
})

// Methods
const reloadTab = () => {
  refreshTab.value = false
  nextTick(() => {
    refreshTab.value = true
  })
}

const clickItem = async (index) => {
  currentIndexName.value = index
  const input = {
    es_connect: sdk.GetSelectEsConnID(),
    index_name: currentIndexName.value
  }
  tabLoading.value = true
  try {
    const res = await ListAction( input)
    if (res.code !== 0) {
      tabLoading.value = false
      ElMessage.error(res.msg)
      return
    }

    const res2 = await GetSettingsAction( input)
    if (res2.code !== 0) {
      tabLoading.value = false
      ElMessage.error(res2.msg)
      return
    }

    currentMappings.value = res.data.list
    const eventAttrOptionsData = [{label: '筛选字段', options: [{value: "_id", label: "_id"}]}]
    const attrMapData = {'2': [{
      "attribute_name": "_id",
      "show_name": "_id",
      "data_type": 3
  }]}
    let propertiesObj = {}
    switch (res.data.ver) {
      case 6:

          for(let doc of Object.keys(currentMappings.value[currentIndexName.value].mappings)){
            for(let col in currentMappings.value[currentIndexName.value].mappings[doc].properties){
              propertiesObj[col] = currentMappings.value[currentIndexName.value].mappings[doc].properties[col]
            }
          }

        console.log("propertiesObj",propertiesObj)

        break
      case 7:
      case 8:
        propertiesObj = currentMappings.value[currentIndexName.value].mappings.properties
        break
    }

    const Int = 1
    const Float = 2
    const String = 3
    const Text = 5

    for (const key in propertiesObj) {
      if (propertiesObj[key].type) {
        eventAttrOptionsData[0].options.push({value: key, label: key})
        const obj = {attribute_name: key, show_name: key}

        switch (propertiesObj[key].type) {
          case 'text':
            obj.data_type = Text
            break
          case 'keyword':
            obj.data_type = String
            break
          case 'byte':
          case 'short':
          case 'integer':
          case 'long':
            obj.data_type = Int
            break
          case 'float':
          case 'half_float':
          case 'scaled_float':
          case 'double':
            obj.data_type = Float
            break
        }

        attrMapData['2'].push(obj)
      }
    }



    attrMap.value = attrMapData
    eventAttrOptions.value = eventAttrOptionsData
    currentSettings.value = res2.data
    tabLoading.value = false
    reloadTab()
  } catch (error) {
    tabLoading.value = false
    ElMessage.error('An error occurred while fetching data.',error)
  }
}

const deleteListByK = (k) => {
  const index = indexList.value.findIndex((v) => v.k.toString() === k.toString())
  if (index !== -1) {
    indexList.value.splice(index, 1)
  }
}

const deleteIndex = async (key, indexName) => {
  try {
    await ElMessageBox.confirm(('删除该索引不可恢复，确认删除吗？'), $t('警告'), {
      confirmButtonText: $t('确定'),
      cancelButtonText: $t('取消'),
      type: 'warning',
    })
  } catch (error) {
    return
  }
  const res = await DeleteAction( {
    es_connect: sdk.GetSelectEsConnID(),
    index_name: indexName,
  })

  if (res.code !== 0) {
    ElMessage({ type: 'error', message: res.msg })
  } else {
    deleteListByK(key)
    ElMessage({ type: 'success', message: res.msg })
  }
}

// onBeforeMount hook to fetch index list
onMounted(async () => {
  console.log('navicat mounted')
  init()
})

onActivated(async () => {
  console.log('navicat activated')
})

const init= async () => {
  const form = {
    cat: 'CatIndices',
    es_connect: sdk.GetSelectEsConnID()
  }
  indexList.value = []
  loadingMenu.value = true
  indexTishiList.value = []
  const res = await CatAction(form)

  if (res.code != 0) {
    ElMessage({
      type: 'error',
      message: res.msg
    })
    loadingMenu.value = false
    return
  }
  for (const k in res.data) {
    const v = res.data[k]
    const obj = { health: v.health, index: v.index, k: k, storeSize: v['store.size'], docsCount: v['docs.count'] }
    indexList.value.push(obj)
    indexTishiList.value.push({ 'value': v.index, 'data': v.index })
  }
  loadingMenu.value = false
}

// Autocomplete search function
const querySearch = (queryString, cb) => {
  const results = queryString ? indexTishiList.value.filter(createFilter(queryString)) : indexTishiList.value
  cb(results)
}

const createFilter = (queryString) => {
  return (item) => {
    return item.value.toLowerCase().includes(queryString.toLowerCase())
  }
}
</script>


<style scoped>
.navicat-container {
  height: 100%;
  background-color: var(--el-bg-color);
}

.content-wrapper {
  height: 100%;
  padding: 16px;
}

.page-split {
  height: calc(100vh - 100px);
  margin-top: 8px;
  background: var(--el-bg-color-page);
  border-radius: 8px;
  box-shadow: 0 2px 12px 0 rgba(0,0,0,0.1);
}

.left-panel {
  height: 100%;
  padding: 16px;
  background: var(--el-bg-color);
  border-right: 1px solid var(--el-border-color-light);
}

.search-container {
  height: calc(100% - 16px);
  overflow: hidden;
}

.search-input {
  width: 100%;
  margin-bottom: 16px;
}

.index-menu {
  height: calc(100% - 56px);
  overflow-y: auto;
  border-right: none;
}

.menu-item {
  margin: 4px 0;
  border-radius: 4px;
  transition: all 0.3s;
}

.menu-item:hover {
  background-color: var(--el-color-primary-light-9);
}

.el-dropdown-link {
  margin-right: 8px;
}

.right-panel {
  height: 100%;
  padding: 16px;
  overflow-x: hidden;
  overflow-y: auto;
  background: var(--el-bg-color);
}

::-webkit-scrollbar {
  width: 6px;
  height: 6px;
}

::-webkit-scrollbar-thumb {
  background: var(--el-border-color);
  border-radius: 3px;
}

::-webkit-scrollbar-track {
  background: var(--el-fill-color-lighter);
  border-radius: 3px;
}
</style>

<style>
.eventNameDisplayInput .ant-input {
  resize: none;
  border: none;
}

.eventNameDisplayInput .ant-input:focus {
  border: none;
  box-shadow: none;
}
</style>

