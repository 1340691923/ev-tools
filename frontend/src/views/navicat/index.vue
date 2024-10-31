<template>
  <div style="display:flex;justify-content:space-between">
    <div class="content_xwl">

      <page-split style="margin-top: 3px" :distribute="0.14" :lineThickness="6" :isVertical="true">
        <template v-slot:first>
          <div
              id="scollL"
              style="height: 95%;width: 100px;display: inline-block; height: 100%;vertical-align: top;width: 100%;"
          >
            <div style="width: 100%;height: calc(100% - 80px); overflow-x: hidden; overflow-y: auto;padding: 10px">
              <el-autocomplete
                  v-model="filterStr"
                  style="width: 90%"
                  clearable
                  :fetch-suggestions="querySearch"
                  :placeholder="$t('请输入索引名')"
              >
                <template #default="{ item }">
                  <span>{{ item.value }}</span>
                </template>
              </el-autocomplete>

              <el-menu v-loading="loadingMenu" style="margin-top: 10px" >

                <template #default  >

                  <el-menu-item v-for="(v, index2) in getIndexList"  :index="index2.toString()" @click="clickItem(v['index'])">
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
          <div style="width: 100%;height: calc(100% - 80px); overflow-x: hidden; overflow-y: auto;padding: 10px">

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
import { ref, computed, onBeforeMount, nextTick } from 'vue'
import { ElMessageBox, ElMessage } from 'element-plus'

import { CatAction } from '@/api/es'
import { filterData } from '@/utils/table'
import { DeleteAction, GetSettingsAction } from '@/api/es-index'
import { ListAction } from '@/api/es-map'

import PageSplit  from "vue3-page-split";

import "vue3-page-split/dist/style.css";

import { getCurrentInstance } from 'vue'
import {sdk} from "@/plugin_sdk/sdk";
// Importing components
import JsonEditor from '@/components/JsonEditor/index.vue'
import Crud from '@/views/navicat/crud.vue'

// Setup instance
const ctx = getCurrentInstance().appContext.config.globalProperties

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
  console.log("123")
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
    const eventAttrOptionsData = [{label: '筛选字段', options: []}]
    const attrMapData = {'2': []}
    let propertiesObj = {}

    switch (res.data.ver) {
      case 6:
        propertiesObj = currentMappings.value[currentIndexName.value].mappings[
            Object.keys(currentMappings.value[currentIndexName.value].mappings)[0]
            ].properties
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
onBeforeMount(async () => {
  init()
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


<style scoped src="@/styles/event.css"/>

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

