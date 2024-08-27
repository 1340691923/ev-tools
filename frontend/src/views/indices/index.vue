<template>
  <div class="app-container">

    <div class="search-container">

      <el-collapse v-model="activeNames">
        <el-collapse-item name="1">
          <template #title>
            <span id="lixianglong">筛选

            </span>
          </template>
          <el-form>
            <el-row :gutter="20">
              <el-col :lg="3" :md="12" :sm="4" :xl="6" :xs="24">
                <el-form-item label="状态">
                  <el-select
                      id="index-health-status"
                      v-model="status"
                      class="filter-item width150"
                      filterable
                      @change="search"
                  >
                    <el-option :label="$t('所有')" value=""/>
                    <el-option label="green" value="green"/>
                    <el-option label="yellow" value="yellow"/>
                    <el-option label="red" value="red"/>
                  </el-select>
                </el-form-item>
              </el-col>
              <el-col :lg="8" :md="12" :sm="12" :xl="6" :xs="24">
                <el-form-item label="索引名">
                  <el-autocomplete
                      class="filter-item width300"
                      id="index-keyword"
                      clearable
                      :fetch-suggestions="querySearch"
                      style="width: 300px"
                      v-model="input"
                      :placeholder="$t('请输入索引名')"
                  >

                    <template #default="{ item }">
                      <span>{{ item.value }}</span>
                    </template>

                  </el-autocomplete>
                  <el-button id="index-search" type="primary" :icon="Search"  @click="search"></el-button>
                </el-form-item>
              </el-col>

            </el-row>
          </el-form>

        </el-collapse-item>
      </el-collapse>

      <el-form style="margin-top: 15px">
        <el-row :gutter="20">
          <el-col :lg="20" :md="12" :sm="4" :xl="6" :xs="24">
            <el-form-item label="索引操作">
              <el-container>
                <el-tooltip content="关闭" placement="top">
                  <el-button
                      :disabled="loadingGroup['close']"
                      id="patchCloseIndex"
                      v-loading="loadingGroup['close']"
                      type="danger"
                      :icon="Hide"
                      @click="Close('close',selectIndexList.join(','))"
                  >
                  </el-button>

                </el-tooltip>
                <el-tooltip content="打开" placement="top">
                  <el-button
                      id="patchOpenIndex"
                      v-loading="loadingGroup['open']"
                      :disabled="loadingGroup['open']"
                      type="success"
                      :icon="View"

                      @click="Open('open',selectIndexList.join(','))"
                  >
                  </el-button>
                </el-tooltip>
                <el-tooltip content="新建" placement="top">
                  <el-button
                      :icon="Plus"
                      id="new-index"
                      @click="openSettingDialog('','add')"
                  >
                  </el-button>
                </el-tooltip>
                <el-tooltip content="删除" placement="top">
                  <el-button
                      :icon="Delete"
                      id="patchDeleteIndex"
                      v-loading="loadingGroup['deleteIndex']"
                      :disabled="loadingGroup['deleteIndex']"

                      type="danger"
                      @click="deleteIndex(selectIndexList.join(','),'deleteIndex')"
                  >
                  </el-button>
                </el-tooltip>
                <el-tooltip content="将节点切换为可读写状态" placement="top">
                  <el-button
                      :icon="Sort"
                      :disabled="readOnlyAllowDeleteLoading"
                      id="readOnlyAllowDelete"
                      v-loading="readOnlyAllowDeleteLoading"
                      type="warning"


                      @click="readOnlyAllowDelete()"
                  >

                  </el-button>
                </el-tooltip>

                <el-tooltip content="将所选索引刷新到磁盘" placement="top">
                  <el-button
                      id="patchFlushIndex"
                      v-loading="loadingGroup['_flush']"
                      :disabled="loadingGroup['_flush']"
                      type="warning"
                      :icon="BrushFilled"
                      @click="Flush('_flush',selectIndexList.join(','))"
                  >
                  </el-button>
                </el-tooltip>
                <el-tooltip content="强制合并索引" placement="top">
                  <el-button
                      :icon="Connection"

                      id="patchForcemergeIndex"
                      v-loading="loadingGroup['_forcemerge']"
                      :disabled="loadingGroup['_forcemerge']"


                      @click="Forcemerge('_forcemerge',selectIndexList.join(','))"
                  >
                  </el-button>
                </el-tooltip>
                <el-tooltip content="刷新索引" placement="top">
                  <el-button
                      id="patchRefreshIndex"
                      #reference
                      v-loading="loadingGroup['_refresh']"
                      :disabled="loadingGroup['_refresh']"
                      :icon="RefreshRight"

                      type="primary"

                      @click="Refresh('_refresh',selectIndexList.join(','))"
                  >
                  </el-button>
                </el-tooltip>

                <el-tooltip content="清空索引" placement="top">
                  <el-button
                      id="patchEmptyIndex"
                      v-loading="loadingGroup['empty']"
                      :disabled="loadingGroup['empty']"
                      :icon="DocumentDelete"
                      type="danger"
                      @click="Empty('empty',selectIndexList.join(','))"
                  >
                  </el-button>
                </el-tooltip>
                <el-tooltip content="清理缓存" placement="top">
                  <el-button
                      :icon="ToiletPaper"

                      id="patchCacheClear"
                      v-loading="loadingGroup['_cache/clear']"
                      :disabled="loadingGroup['_cache/clear']"


                      type="warning"
                      @click="CacheClear('_cache/clear',selectIndexList.join(','))"
                  >
                  </el-button>
                </el-tooltip>

              </el-container>
            </el-form-item>
          </el-col>

        </el-row>
      </el-form>



    </div>


    <el-table
        v-loading="connectLoading"
        :data="list"
        @selection-change="selectChange"
    >
      <el-table-column
          type="selection"
          width="55"
      />


      <el-table-column align="center" :label="$t('健康')" width="100">
        <template #default="scope">
          <el-button
              v-if="scope.row.health == 'green'" type="success" circle/>
          <el-button
              v-if="scope.row.health == 'yellow'" type="warning" circle/>
          <el-button
              v-if="scope.row.health == 'red'" type="danger" circle/>
        </template>
      </el-table-column>
      <el-table-column align="center" :label="$t('状态')" width="100">
        <template #default="scope">
          <el-button

              v-show="scope.row.status == 'open'"
              type="success"

              @click="Close('close',scope.row.index)"
          >{{ $t('开启') }}
          </el-button>
          <el-button

              v-show="scope.row.status == 'close'"
              type="danger"

              @click="Open('open',scope.row.index)"
          >{{ $t('关闭') }}
          </el-button>
        </template>
      </el-table-column>
      <el-table-column align="center" :label="$t('索引名')" width="180">
        <template #default="scope">
          {{ scope.row.index }}
        </template>
      </el-table-column>
      <el-table-column align="center" :label="$t('uuid')" width="180">
        <template #default="scope">
          {{ scope.row.uuid }}
        </template>
      </el-table-column>
      <el-table-column align="center" :label="$t('主分片数')" width="80" prop="pri" sortable>
        <template #default="scope">
          {{ scope.row.pri }}
        </template>
      </el-table-column>
      <el-table-column align="center" :label="$t('副本分片数量')" width="80" prop="rep" sortable>
        <template #default="scope">
          {{ scope.row.rep }}
        </template>
      </el-table-column>
      <el-table-column align="center" :label="$t('文档总数')" width="80" prop="docs->count" sortable>
        <template #default="scope">
          {{ bigNumberTransform(scope.row["docs.count"]) }}
        </template>
      </el-table-column>
      <el-table-column align="center" :label="$t('删除状态的文档')" width="80" prop="docs->deleted" sortable>
        <template #default="scope">
          {{ scope.row["docs.deleted"] }}
        </template>
      </el-table-column>
      <el-table-column align="center" :label="$t('主+副本分片大小')" width="120" prop="store->size" sortable>
        <template #default="scope">
          {{ scope.row["store.size"] }}
        </template>
      </el-table-column>
      <el-table-column align="center" :label="$t('主分片大小')" width="150" prop="pri->store->size" sortable>
        <template #default="scope">
          {{ scope.row["pri.store.size"] }}
        </template>
      </el-table-column>

      <el-table-column align="center" width="190" :label="$t('操作')" fixed="right" >
        <template #default="scope">
            <el-tooltip content="修改配置" placement="top">
            <el-button
                :icon="Setting"
                type="info"
                @click="openSettingDialog(scope.row.index,'update')"
            >
            </el-button>
            </el-tooltip>
            <el-tooltip content="修改映射" placement="top">
            <el-button
                :icon="EditPen"
                type="waring"
                @click="openMappingEditDialog(scope.row.index,false)"
            >
            </el-button>
            </el-tooltip>
            <el-tooltip content="更多" placement="top">
            <el-button
                :icon="More"
                type="primary"
                @click="openDrawer(scope.row.index)"
            >
            </el-button>
            </el-tooltip>
        </template>
      </el-table-column>
    </el-table>
    <el-pagination
        v-if="pageshow"
        class="pagination-container"
        :current-page="page"
        :page-sizes="[10, 20, 30, 50,100,150,200,500,1000]"
        :page-size="limit"
        layout="total, sizes, prev, pager, next"
        :total="total"
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
    />

    <el-drawer
        ref="drawer"
        :title="indexName"
        :before-close="drawerHandleClose"
        v-model="drawerShow"
        direction="rtl"
        custom-class="demo-drawer"
        close-on-press-escape
        destroy-on-close
        :size="isMobile?'100%':'50%'"
    >
      <el-tabs v-model="activeName" @tab-click="changeTab">
        <el-tab-pane  :label="$t('设置')" name="Settings">
          <json-editor
              v-if="showSettings"
              v-model:value="settingData"
              styles="width: 100%"
              :read="true"
              :title="$t('设置')"
          />

        </el-tab-pane>
        <el-tab-pane  :label="$t('映射')" name="Mapping">
          <div class="search-container operate">

            <el-tag type="warning" >{{ $t('切换为其它索引的映射') }}</el-tag>

            <index-select
                :clearable="true"
                style="width: 150px"
                :placeholder="$t('请选择索引名')"
                @change="changeMapToAnotherIndex"
            />
            <el-tag type="primary" >{{ $t('操作') }}</el-tag>
            <el-button
                :disabled="loadingGroup['saveMappinng']"
                v-loading="loadingGroup['saveMappinng']"
                type="primary"
                @click="saveMappinng"
            >{{ $t('修改') }}
            </el-button>
            <el-link type="danger">{{ $t('【注意：只能新增映射字段不可修改映射字段类型】') }}</el-link>
          </div>
          <json-editor
              v-if="showMapping"
              v-model:value="mappingData"
              styles="width: 100%"
              :read="false"
              :title="$t('映射')"
              @getValue="getMapping"
          />
        </el-tab-pane>

        <el-tab-pane  label="Stats" name="Stats">
          <json-editor
              v-if="showStats"
              v-model:value="statsData"

              styles="width: 100%;"
              :read="true"
              title="Stats"
          />
        </el-tab-pane>
        <el-tab-pane  :label="$t('编辑索引配置')" name="editSettings">
          <el-form>
            <el-form-item :label="$t('编辑索引配置')">
              <el-button

                  type="primary"  @click="submitSettings()">{{ $t('提交') }}
              </el-button>
              <el-button

              @click="resetSettings">{{ $t('重置') }}
              </el-button>
            </el-form-item>
            <el-form-item>
              <json-editor
                  v-if="showEditSettings"
                  v-model:value="editSettingsData"
                  :point-out="pointOut"
                  styles="width: 100%;"
                  :read="false"
                  :title="$t('编辑配置')"
                  @getValue="getSettings"
              />
            </el-form-item>
          </el-form>
        </el-tab-pane>
        <el-tab-pane  :label="$t('修改别名')" name="alias">
          <alias v-if="activeName == 'alias'" :index-name="indexName"/>
        </el-tab-pane>
        <div class="search-container operate">
          <el-form>
            <el-form-item label="操作">
              <el-tooltip content="关闭" placement="top">
                <el-button
                    :icon="Hide"
                    :disabled="loadingGroup['close']"
                    v-loading="loadingGroup['close']"
                    type="danger"
                    @click="Close('close',indexName)"
                >
                </el-button>
              </el-tooltip>
              <el-tooltip content="打开" placement="top">
                <el-button
                    :icon="View"
                    :disabled="loadingGroup['open']"
                    v-loading="loadingGroup['open']"
                    type="success"
                    @click="Open('open',indexName)"
                >
                </el-button>
              </el-tooltip>
              <el-tooltip content="强制合并索引【forcemerge操作,手动释放磁盘空间】" placement="top">
                <el-button
                    :disabled="loadingGroup['_forcemerge']"
                    v-loading="loadingGroup['_forcemerge']"
                    :icon="Connection"
                    @click="Forcemerge('_forcemerge',indexName)"
                >
                </el-button>
              </el-tooltip>
              <el-tooltip content="刷新索引【为了让最新的数据可以立即被搜索到】" placement="top">

                <el-button

                    :disabled="loadingGroup['_refresh']"
                    v-loading="loadingGroup['_refresh']"
                    :icon="RefreshRight"
                    type="primary"
                    @click="Refresh('_refresh',indexName)"
                >
                </el-button>
              </el-tooltip>

              <el-tooltip content="将索引刷新到磁盘【让数据持久化到磁盘中】" placement="top">
                <el-button
                    :icon="BrushFilled"
                    :disabled="loadingGroup['_flush']"
                    v-loading="loadingGroup['_flush']"
                    type="info"
                    @click="Flush('_flush',indexName)"
                >
                </el-button>
              </el-tooltip>
              <el-tooltip content="清理缓存" placement="top">
                <el-button
                    :icon="ToiletPaper"
                    :disabled="loadingGroup['_cache/clear']"
                    v-loading="loadingGroup['_cache/clear']"


                    type="warning"
                    @click="CacheClear('_cache/clear',indexName)"
                >
                </el-button>
              </el-tooltip>
              <el-tooltip content="删除索引" placement="top">
                <el-button
                    :icon="Delete"
                    :disabled="loadingGroup['deleteIndex']"
                    v-loading="loadingGroup['deleteIndex']"

                    type="danger"

                    @click="deleteIndex(indexName,'deleteIndex')"
                >
                </el-button>
              </el-tooltip>
              <el-tooltip content="清空索引" placement="top">
                <el-button
                    :disabled="loadingGroup['empty']"
                    v-loading="loadingGroup['empty']"

                    type="danger"
                    :icon="DocumentDelete"
                    @click="Empty('empty',indexName)"
                >
                </el-button>
              </el-tooltip>
            </el-form-item>
          </el-form>
        </div>
      </el-tabs>

    </el-drawer>
    <settings
        v-if="openSettings"
        :index-name="indexName"
        :settings-type="settingsType"
        @finished="search"
        :open="openSettings"
        @close="closeSettings"
    />
    <mappings
        v-if="openMappings"
        :index-name="indexName"
        :title="mappingTitle"
        :open="openMappings"
        @close="closeMappings"
    />

  </div>
</template>

<script>
import {clone} from '@/utils/index'

import {filterData} from '@/utils/table'
import {
  CatAction,
  RecoverCanWrite,
} from '@/api/es'
import {
  CacheClear,
  Close,
  Empty,
  Flush,
  Forcemerge,
  Open,
  Refresh
} from '@/api/es-index'
import {bigNumberTransform} from '@/utils/format'
import {CreateAction, DeleteAction, GetSettingsAction, GetSettingsInfoAction, StatsAction} from '@/api/es-index'
import {esSettingsWords} from '@/utils/base-data'
import {ListAction, UpdateMappingAction} from '@/api/es-map'
import {ElMessage} from "element-plus";
import Settings from "@/views/indices/components/settings.vue";
import Mappings from "@/views/indices/components/mapping.vue";
import JsonEditor from "@/components/JsonEditor/index.vue";
import Alias from "@/views/indices/components/alias.vue";
import IndexSelect from "@/components/index/select.vue";
import {sdk} from "@/plugin_sdk/sdk"
import * as Icon from '@element-plus/icons-vue'

export default {
  name: 'indices',
  setup(){
    return Icon
  },
  components: {
    Settings,
    Mappings,
    JsonEditor,
    Alias,
    IndexSelect,
  },

  data() {
    return {
      activeNames:'1',
      indexTishiList: [],
      modName: '索引管理',
      aliasList: [],
      pointOut: esSettingsWords,
      settings: {},
      readOnlyAllowDeleteLoading: false,
      loadingGroup: {
        'empty': false,
        'close': false,
        'open': false,
        '_forcemerge': false,
        '_refresh': false,
        '_flush': false,
        '_cache/clear': false,
        'deleteIndex': false,
        '_all/_flush': false,
        'saveMappinng': false
      },

      forceMergeLoading: false,
      tabLoading: false,
      settingData:'{}',
      mappingData:'{}',
      statsData:'{}',
      editSettingsData:'{}',
      activeName: 'Settings',
      showEditSettings:true,
      showSettings:true,
      showMapping:true,
      showStats:true,
      drawerShow: false,
      settingsType: 'add',
      mappingTitle: '',
      indexName: '',
      openSettings: false,
      total: 0,
      connectLoading: false,
      page: 1,
      limit: 10,
      pageshow: true,
      list: [],
      input: '',
      status: '',
      mappings: {},
      selectIndexList: [],
      openMappings: false,
      allList: [],
      max: 8,
    }
  },
  destroyed() {
    sessionStorage.setItem('CatIndices', this.input)
  },
  computed:{
    isMobile(){
      return sdk.IsMobile()
    }
  },
  mounted() {
    const input = sessionStorage.getItem('CatIndices')
    if (input != null) {
      this.input = input
    }
    this.startGuid()
    this.GetMapAction()
    this.searchData()
  },
  methods: {
    openMappingDialog(index) {
      this.indexName = index
      this.openMappings = true
    },
    async finishGuid() {

    },
    async startGuid() {},
    selectChange(row) {
      this.selectIndexList = []
      for (const v of row) {
        this.selectIndexList.push(v.index)
      }
    },
    async changeMapToAnotherIndex(indexName) {
      if (indexName == '') {
        indexName = this.indexName
      }
      const input = {}
      input['es_connect'] = sdk.GetSelectEsConnID()
      input['index_name'] = indexName

      const res = await ListAction(input)

      if (res.code == 0) {
        this.mappingData = JSON.stringify(res.data['list'][indexName].mappings, null, '\t')
      }
    },
    getMapping(v) {
      this.mappingData = v
    },
    async saveMappinng() {
      let activeData = clone(this.mappingData)
      try {
        activeData = JSON.parse(activeData)
      } catch (e) {
        ElMessage({
          type: 'error',
          message: 'JSON格式不正确'
        })
        return
      }
      const activeDataKeys = Object.keys(activeData)
      if (activeDataKeys.length == 0) {
        ElMessage({
          type: 'error',
          message: '请按格式写type名字'
        })
        return
      }

      const input = {}
      input['es_connect'] = sdk.GetSelectEsConnID()
      input['index_name'] = this.indexName

      if (activeData[activeDataKeys[0]].to$t() != 'false') {
        input['properties'] = activeData[activeDataKeys[0]]
        input['type_name'] = activeDataKeys[0]
      } else {
        input['properties'] = activeData
      }

      this.loadingGroup['saveMappinng'] = true
      const {data, code, msg} = await UpdateMappingAction(input)
      this.loadingGroup['saveMappinng'] = false
      if (code == 0) {
        ElMessage({
          type: 'success',
          message: msg
        })
      } else {
        ElMessage({
          type: 'error',
          message: msg
        })
      }
    },
    openMappingEditDialog(indexName, haveMapping) {
      if (haveMapping) {
        this.mappingTitle = $t('新增字段')
      } else {
        this.mappingTitle = $t('新增映射结构')
      }
      this.indexName = indexName

      this.openMappings = true
    },
    async GetMapAction() {
      const input = {}
      input['es_connect'] = sdk.GetSelectEsConnID()
      const {code, data, msg} = await ListAction(input)
      if (code == 0) {
        this.mappings = data['list']
      } else {
        ElMessage({
          type: 'error',
          message: msg
        })
        return
      }
    },
    submitSettings() {
      const input = {}
      input['es_connect'] = sdk.GetSelectEsConnID()
      try {
        input['settings'] = JSON.parse(this.editSettingsData)
      } catch (e) {
        console.log(e)
        ElMessage({
          type: 'error',
          message: 'JSON 解析异常'
        })
        return
      }
      input['index_name'] = this.indexName
      input['types'] = 'update'
      const loading = this.$loading({
        lock: true,
        text: 'Loading',
        spinner: 'el-button-loading',
        background: 'rgba(0, 0, 0, 0.7)'
      })
      CreateAction(input).then(res => {
        if (res.code == 0 || res.code == 200) {
          ElMessage({
            type: 'success',
            message: res.msg
          })
          this.search()
        } else {
          ElMessage({
            type: 'error',
            message: res.msg
          })
        }
        loading.close()
      }).catch(err => {
        loading.close()
      })
    },
    getSettings(value) {
      this.editSettingsData = value
    },
    resetSettings() {
      this.changeTab()
    },
    Forcemerge(command, indexName) {
      const input = {}
      input['es_connect'] = sdk.GetSelectEsConnID()
      input['index_name'] = indexName
      input['command'] = command

      this.loadingGroup[command] = true
      Forcemerge(input).then(res => {
        if (res.code == 0 || res.code == 200) {
          ElMessage({
            type: 'success',
            message: res.msg
          })
          this.search()
        } else {
          ElMessage({
            type: 'error',
            message: res.msg
          })
        }
        this.loadingGroup[command] = false
      })
          .catch(err => {
            this.loadingGroup[command] = false
          })
    },
    Empty(command, indexName) {
      const input = {}
      input['es_connect'] = sdk.GetSelectEsConnID()
      input['index_name'] = indexName
      input['command'] = command

      this.loadingGroup[command] = true
      Empty(input).then(res => {
        if (res.code == 0 || res.code == 200) {
          ElMessage({
            type: 'success',
            message: res.msg
          })
          this.search()
        } else {
          ElMessage({
            type: 'error',
            message: res.msg
          })
        }
        this.loadingGroup[command] = false
      })
          .catch(err => {
            this.loadingGroup[command] = false
          })
    },
    Flush(command, indexName) {
      const input = {}
      input['es_connect'] = sdk.GetSelectEsConnID()
      input['index_name'] = indexName
      input['command'] = command

      this.loadingGroup[command] = true
      Flush(input).then(res => {
        if (res.code == 0 || res.code == 200) {
          ElMessage({
            type: 'success',
            message: res.msg
          })
          this.search()
        } else {
          ElMessage({
            type: 'error',
            message: res.msg
          })
        }
        this.loadingGroup[command] = false
      })
          .catch(err => {
            this.loadingGroup[command] = false
          })
    },
    CacheClear(command, indexName) {
      const input = {}
      input['es_connect'] = sdk.GetSelectEsConnID()
      input['index_name'] = indexName
      input['command'] = command

      this.loadingGroup[command] = true
      CacheClear(input).then(res => {
        if (res.code == 0 || res.code == 200) {
          ElMessage({
            type: 'success',
            message: res.msg
          })
          this.search()
        } else {
          ElMessage({
            type: 'error',
            message: res.msg
          })
        }
        this.loadingGroup[command] = false
      })
          .catch(err => {
            this.loadingGroup[command] = false
          })
    },
    Open(command, indexName) {
      const input = {}
      input['es_connect'] = sdk.GetSelectEsConnID()
      input['index_name'] = indexName
      input['command'] = command

      this.loadingGroup[command] = true
      Open(input).then(res => {
        if (res.code == 0 || res.code == 200) {
          ElMessage({
            type: 'success',
            message: res.msg
          })
          this.search()
        } else {
          ElMessage({
            type: 'error',
            message: res.msg
          })
        }
        this.loadingGroup[command] = false
      })
          .catch(err => {
            this.loadingGroup[command] = false
          })
    },
    Close(command, indexName) {
      const input = {}
      input['es_connect'] = sdk.GetSelectEsConnID()
      input['index_name'] = indexName
      input['command'] = command

      this.loadingGroup[command] = true
      Close(input).then(res => {
        if (res.code == 0 || res.code == 200) {
          ElMessage({
            type: 'success',
            message: res.msg
          })
          this.search()
        } else {
          ElMessage({
            type: 'error',
            message: res.msg
          })
        }
        this.loadingGroup[command] = false
      })
          .catch(err => {
            this.loadingGroup[command] = false
          })
    },
    Refresh(command, indexName) {
      const input = {}
      input['es_connect'] = sdk.GetSelectEsConnID()
      input['index_name'] = indexName
      input['command'] = command

      this.loadingGroup[command] = true
      Refresh(input).then(res => {
        if (res.code == 0 || res.code == 200) {
          ElMessage({
            type: 'success',
            message: res.msg
          })
          this.search()
        } else {
          ElMessage({
            type: 'error',
            message: res.msg
          })
        }
        this.loadingGroup[command] = false
      })
          .catch(err => {
            this.loadingGroup[command] = false
          })
    },
    async changeTab(data) {

      let index = 0

      if(data !=undefined){
        index = Number(data.index)
      }

      const input = {}
      let res = {}
      input['es_connect'] = sdk.GetSelectEsConnID()
      input['index_name'] = this.indexName
      switch (index) {
        case 0:
          res = await GetSettingsInfoAction(input)

          if (res.code == 0) {
            this.settingData = JSON.stringify(res.data, null, '\t')
          }

          this.showSettings = false

          this.$nextTick(()=>{
            this.showSettings = true
          })

          return

        case 1:
          res = await ListAction(input)

          if (res.code == 0) {
            this.mappingData = JSON.stringify(res.data['list'][this.indexName].mappings, null, '\t')
            if (Object.keys(res.data['list'][this.indexName].mappings).length == 0) {
              ElMessage({
                type: 'error',
                message: '您还没有设置映射结构'
              })
            }
          }

          this.showMapping = false

          this.$nextTick(()=>{
            this.showMapping = true
          })
          return
        case 2:

          res = await StatsAction(input)
          if (res.code == 0) {

            this.statsData = JSON.stringify(res.data, null, '\t')
          }

          this.showStats = false

          this.$nextTick(()=>{
            this.showStats = true
          })

          return

        case 3:

          const {data} = await GetSettingsAction(input)

          const deleteKeyArr = [
            'creation_date', 'version', 'provided_name', 'uuid', 'format', 'number_of_shards'
          ]

          for (const key of deleteKeyArr) {
            if (data['index'].hasOwnProperty(key)) {
              delete data['index'][key]
            }
          }

          this.editSettingsData = JSON.stringify(data, null, '\t')

          this.showEditSettings = false

          this.$nextTick(()=>{
            this.showEditSettings = true
          })
          return
        case 4:
          return
        default:
          return
      }
    },
    openDrawer(indexName) {
      this.indexName = indexName
      this.changeTab()
      this.drawerShow = true
    },
    drawerHandleClose(done) {
      this.indexName = ''
      this.activeName = 'Settings'
      done()
    },
    bigNumberTransform(value) {
      return bigNumberTransform(value)
    },
    openSettingDialog(indexName, settingsType) {
      this.indexName = indexName
      this.settingsType = settingsType
      this.openSettings = true
    },
    readOnlyAllowDelete() {
      const input = {}
      input['es_connect'] = sdk.GetSelectEsConnID()
      this.readOnlyAllowDeleteLoading = true
      RecoverCanWrite(input).then(res => {
        if (res.code == 0 || res.code == 200) {
          ElMessage({
            type: 'success',
            message: res.msg
          })
        } else {
          ElMessage({
            type: 'error',
            message: res.msg
          })
        }
        this.readOnlyAllowDeleteLoading = false
      }).catch(err => {
        this.readOnlyAllowDeleteLoading = false
      })
    },
    querySearch(queryString, cb) {

      let queryData = JSON.parse(JSON.stringify(this.indexTishiList))
      if (queryString == undefined) queryString = ""
      if (queryString.trim() == '') {
        if (queryData.length > this.max) {
          cb(queryData.slice(0, this.max))
        } else {
          cb(queryData)
        }
        return;
      }


      queryData = filterData(queryData, queryString.trim())

      if (queryData.length > this.max) {
        cb(queryData.slice(0, this.max))
      } else {
        cb(queryData)
      }
    },
    deleteIndex(indexName, loadingType) {
      this.$confirm('确定删除该索引吗?', '警告', {
        confirmButtonText: '确认',
        cancelButtonText: '取消',
        type: 'warning'
      })
          .then(async () => {
            const input = {}
            input['es_connect'] = sdk.GetSelectEsConnID()
            input['index_name'] = indexName
            this.loadingGroup[loadingType] = true
            DeleteAction(input).then(res => {
              if (res.code == 0 || res.code == 200) {
                ElMessage({
                  type: 'success',
                  message: indexName + '已删除'
                })
                this.searchData()
                this.loadingGroup[loadingType] = false
              } else {
                ElMessage({
                  type: 'error',
                  message: res.msg
                })
                this.loadingGroup[loadingType] = false
              }
            }).catch(err => {

            })
          })
          .catch(err => {
            console.error(err)
          })
    },
    closeSettings() {
      this.indexName = ''
      this.settingsType = 'add'
      this.openSettings = false
    },
    closeMappings() {
      this.indexName = ''
      this.mappingTitle = 'add'
      this.openMappings = false
    },
    search() {
      this.page = 1
      this.pageshow = false
      this.searchData()
      this.$nextTick(() => {
        this.pageshow = true
      })
    },
    filterData(list, input) {
      return filterData(list, input)
    },
    // 当每页数量改变
    handleSizeChange(val) {
      console.log(`每页 ${val} 条`)
      this.limit = val
      this.pageLimit()
    },
    // 当当前页改变
    handleCurrentChange(val) {
      console.log(`当前页: ${val}`)
      this.page = val
      this.pageLimit()
    },
    pageLimit() {
      this.list = this.allList.filter((item, index) =>
          index < this.page * this.limit && index >= this.limit * (this.page - 1)
      )
    },
    searchData() {
      this.connectLoading = true
      const form = {
        cat: 'CatIndices',
        es_connect: sdk.GetSelectEsConnID()
      }
      this.indexTishiList = []
      CatAction(form).then(res => {
        if (res.code == 0) {
          const list = res.data

          for (const index in list) {
            const obj = list[index]
            // 把 . 转成 ->
            for (const key in obj) {
              let value = parseInt(obj[key])
              if (isNaN(value)) {
                value = obj[key]
              }
              list[index][key.split('.').join('->')] = value
            }
            this.indexTishiList.push({'value': obj["index"], 'data': obj["index"]})
          }

          let tmpList = []
          if (this.status.trim() != '') {
            for (const v of list) {
              if (v['health'] == this.status.trim()) {
                tmpList.push(v)
              }
            }
          } else {
            tmpList = list
          }
          tmpList = filterData(tmpList, this.input.trim())
          this.allList = tmpList
          this.total = tmpList.length
          this.pageLimit()
        } else {
          ElMessage({
            type: 'error',
            message: res.msg
          })
        }
        this.connectLoading = false
      }).catch(err => {
        console.log(err)

        this.connectLoading = false
      })
    }
  }
}
</script>

<style scoped>
.filter-item {

}

.aliasName {
  width: 400px;
}

.margin-left-10 {
  margin-left: 10px
}

.width300 {
  width: 300px;
}

.width150 {
  width: 150px;
}



 .el-table .sort-caret.descending{
  bottom: 0px;
}

</style>
