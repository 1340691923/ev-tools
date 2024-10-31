<template>
  <div>
    <div class="app-container" style="height: 95%;">
      <template v-if="indexName != ''">
        <div class="search-container">
          <el-tabs v-model="crudTab">
            <el-tab-pane :lazy="true" :label="$t('筛选')" name="filter">
              <filter-where
                  v-model:value="whereFilter"
                  table-typ="2"
                  :data-type-map="attrMap"
                  :options="eventAttrOptions"
              />
            </el-tab-pane>
            <el-tab-pane lazy="true" :label="$t('排序')" name="sort">
              <div class="search-container">
                <div class="relation-row">
                  <div v-for="(v,index) in sortList" :key="index" class="action-row row___xwl">
                    <div class="action-left">
                      <el-select

                          v-model="sortList[index].col"

                          filterable

                          style="width: 300px;margin-left: 30px"
                          :placeholder="$t('请选择')"
                      >
                        <el-option-group
                            v-for="group in eventAttrOptionsProp"
                            :key="group.label"
                            :label="group.label"
                        >
                          <el-option
                              v-for="item in group.options"
                              :key="item.value"
                              :label="item.label"
                              :value="item.value"
                          >
                            <span style="float: left">{{ item.label }}</span>
                            <span style="float: right; color: #8492a6; font-size: 13px">{{ item.value }}</span>
                          </el-option>
                        </el-option-group>
                      </el-select>
                      <el-select
                          v-model="sortList[index].sortRule"


                          style="width: 100px;margin-left: 30px"
                          :placeholder="$t('请选择排序规则')"
                          filterable
                      >
                        <el-option
                            :label="$t('正序排')"
                            value="asc"
                        />
                        <el-option
                            :label="$t('倒序排')"
                            value="desc"
                        />
                      </el-select>
                      <el-button
                          style="margin-left: 30px"

                          class="actions_xwl_btn"
                          @click="deleteSort(index)"
                      >删除</el-button>
                    </div>

                  </div>
                </div>
              </div>
              <div style="padding: 0 12px;">
                <span style="color:#3d90ff" class="footadd___2D4YB" @click="pushSort">
                  {{ $t('增加排序') }}
                </span>
              </div>
            </el-tab-pane>
          </el-tabs>
        </div>
        <el-card class="box-card">
          <div  class="search-container">
           <el-form>
             <el-form-item>
               <!--<el-button

icon="el-icon-plus" type="primary" @click.native="openAddDialog = true" >添加文档</el-button>-->
               <el-button

                   type="warning"

                   @click="getdsl(1)"
               >查看查询语句
               </el-button>
               <el-button

                   type="success"

                   @click="search(1)"
               >查询
               </el-button>
               <el-button
                   :disabled="downloadLoading"
                   :loading="downloadLoading"

                   type="primary"

                   @click="download()"
               >下载
               </el-button>
             </el-form-item>
           </el-form>
          </div>

          <el-table
              v-loading="tableLoading"
              height="400"
              :data="tableData"
              use-virtual
              ref="multipleTable"
              :row-height="30"
              show-overflow-tooltip
          >
            <el-table-column
                label="ID"
                align="center"
                fixed
                min-width="100"
                prop="_id"
            >
            </el-table-column>
            <template
                v-for="(val,key) in tableHeader"

            >
              <el-table-column
                  v-if="val != 'xwl_index' && val != '_id'"
                  :prop="val"
                  :sortable="true"
                  align="center"
                  :label="val"
              >

              </el-table-column>
            </template>

            <el-table-column align="center" :label="$t('操作')" fixed="right" width="200">
              <template #default="scope">
                <el-button-group>
                  <el-button

                      type="primary"


                      @click="look(scope.$index)"
                  >
                    编辑
                  </el-button>

                  <el-button

                      type="danger"

                      @click="deleteByID(scope.row,scope.$index)"
                  >{{ $t('删除') }}
                  </el-button>
                </el-button-group>
              </template>
            </el-table-column>

          </el-table>
          <el-pagination
              :current-page="input.page"
              :page-sizes="[10, 20, 30, 50,100,150,200]"
              :page-size="input.limit"
              layout="total, sizes, prev, pager, next, jumper"
              :total="count"
              @size-change="handleSizeChange"
              @current-change="search"
          />
        </el-card>

        <el-drawer
            ref="drawer"
            title="详细数据"
            :before-close="drawerHandleClose"
            v-model="drawerShow"

            direction="rtl"
            close-on-press-escape
            destroy-on-close
            size="80%"
        >

          <el-tag >操作</el-tag>
          <el-button
              type="primary"
              @click="updateByID"
          >修改
          </el-button>

          <json-editor
              v-if="showJsonEditor1"
              v-model:value="jsonDataString"
              height="900"
              class="res-body"
              styles="width: 100%"
              title="详细数据"
              @getValue="getEditDoc"
          />

        </el-drawer>

        <el-drawer
            ref="drawer"
            title="查询DSL"
            :before-close="drawerHandleClose"
            v-model="queryDslShow"

            direction="rtl"
            close-on-press-escape
            destroy-on-close
            size="80%"
        >

          <json-editor
              v-if="showJsonEditor2"
              v-model:value="queryDslString"
              height="900"
              class="res-body"
              styles="width: 100%"
              title="DSL"
              :read="true"
          />

        </el-drawer>
      </template>

      <div
          v-else
          style="padding: 40px;width: 300px;height: 800px; text-align: center;margin: 0px auto;position: relative;top: 30%"
      >
        <el-empty content="请先点击左侧索引名称"></el-empty>
      </div>

    </div>
  </div>
</template>

<script>

import JsonEditor from '@/components/JsonEditor/index.vue'
import FilterWhere from '@/components/AnalyseTools/FilterWhere/index.vue'
import {sdk} from "@/plugin_sdk/sdk";
import { GetDSL, GetList, Download } from '@/api/es-crud'
import { ListAction } from '@/api/es-map'
import { DeleteRowByIDAction, InsertAction, UpdateByIDAction } from '@/api/es-doc'
import { clone } from '../../utils'
import {ElMessageBox,ElMessage } from "element-plus";
import {getCurrentInstance} from "vue";

export default {
  name: 'Crud',
  setup(){
    const ctx = getCurrentInstance().appContext.config.globalProperties
    return {ctx}
  },
  components: {
    JsonEditor,FilterWhere,

  },
  computed:{
    jsonDataString(){
      return JSON.stringify(this.jsonData,null, '\t')
    },
    queryDslString(){
      return JSON.stringify(this.queryDsl,null, '\t')
    }
  },
  props: {
    indexName: {
      type: String,
      default: 'test2'
    },
    attrMapProp: {
      type: Array,
      default: []
    },
    eventAttrOptionsProp: {
      type: Array,
      default: []
    }
  },
  data() {
    return {
      ctx:null,
      showJsonEditor1:false,
      showJsonEditor2:false,
      properties: {},
      openAddDialog: false,
      tableLoading: false,
      tableHeader: [],
      tableData: [],
      crudTab: 'filter',
      eventAttrOptions: this.eventAttrOptionsProp,
      attrMap: this.attrMapProp,
      whereFilter: {
        filterType: 'COMPOUND',
        filts: [],
        relation: '且'
      },
      sortList: [],
      input: {
        page: 1,
        limit: 20
      },
      count: 0,
      typName: '',
      newDoc: {},
      jsonData: {},
      tableDataClone: [],
      drawerShow: false,
      queryDslShow: false,
      queryDsl: {},
      downloadLoading: false
    }
  },
  mounted() {
    if (this.indexName != '') this.search(1)
  },
  methods: {
    download() {
      const form = {
        index_name: this.indexName,
        relation: this.whereFilter,
        sort_list: this.sortList,
        es_connect: sdk.GetSelectEsConnID()
      }
      const link = document.createElement('a')
      this.downloadLoading = true
      Download(form).then(res => {
        if (res) {
          const blob = new Blob([res], { type: 'application/vnd.ms-excel' })
          const objectUrl = URL.createObjectURL(blob) // 创建URL
          link.href = objectUrl
          link.download = this.indexName + '.csv' // 自定义文件名
          link.click() // 下载文件
          URL.revokeObjectURL(objectUrl) // 释放内存
        }
        this.downloadLoading = false
      })
          .catch((err) => {
            this.downloadLoading = false
            console.error(err)
          })
    },
    drawerShowFn() {
      this.drawerShow = false // 先关闭，
      this.$nextTick(() => {
        this.drawerShow = true
        this.$nextTick(()=>{
          this.showJsonEditor1 = true
        })
      })
    },
    look(index) {
      this.index = index
      console.log('this.tableDataClone', this.tableDataClone)
      this.jsonData = this.tableDataClone[index]
      this.drawerShowFn()
    },
    async updateByID() {
      console.log('this.jsonData', this.jsonData)
      const editData = this.jsonData
      const doc = editData['_source']

      const input = {}
      input['es_connect'] = sdk.GetSelectEsConnID()
      input['index'] = editData['_index']
      input['type_name'] = editData['_type']
      input['id'] = editData['_id']
      input['json'] = doc

      const res = await UpdateByIDAction(input)
      if (res.code == 0) {
        this.search(1)
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
    },
    getEditDoc(v) {
      try {
        const editDoc = JSON.parse(v)
        this.jsonData = editDoc
      } catch (e) {

      }
    },
    getNewDoc(doc) {
      try {
        this.newDoc = JSON.parse(doc)
      } catch (e) {

      }
    },
    async add() {
      const doc = this.newDoc

      const input = {}
      input['es_connect'] = sdk.GetSelectEsConnID()
      input['index'] = this.indexName
      input['type_name'] = this.typName
      input['json'] = doc

      const res = await InsertAction(input)
      if (res.code == 0) {
        ElMessage({
          type: 'success',
          message: res.msg + '(_id为:' + res.data._id + ')',
          duration: 20000
        })
      } else {
        ElMessage({
          type: 'error',
          message: res.msg
        })
      }
    },
    drawerHandleClose(done) {
      done()
    },
    async deleteByID(row, index) {
      ElMessageBox.confirm('确定删除ID为【' + row['_id'] + '】的这条文档吗?', '警告', {
        confirmButtonText: '确认',
        cancelButtonText: '取消',
        type: 'warning'
      })
          .then(async() => {
            const input = {}
            input['es_connect'] = sdk.GetSelectEsConnID()
            input['index_name'] = this.indexName
            input['type'] = this.typName
            input['id'] = row['_id']

            const res = await DeleteRowByIDAction(input)
            if (res.code == 0) {
              setTimeout(async() => {
                await this.search(1)
                ElMessage({
                  type: 'success',
                  message: res.msg
                })
              }, 1000)
            } else {
              ElMessage({
                type: 'error',
                message: res.msg
              })
            }
          })
          .catch(err => {
            console.error(err)
          })
    },
    async initTableData(resData) {
      if (resData.hasOwnProperty('hits')) {
        if (resData['hits']['hits'].length > 0) {
          this.typName = resData['hits']['hits'][0]['_type']

          const sourceArr = []
          this.tableDataClone = clone(resData['hits']['hits'])
          for (const index in resData['hits']['hits']) {
            const _source = resData['hits']['hits'][index]['_source']
            _source['_id'] = resData['hits']['hits'][index]['_id']
            _source['_score'] = resData['hits']['hits'][index]['_score']
            _source['xwl_index'] = index
            sourceArr.push(_source)
          }

          this.tableData = sourceArr

          for (const index in this.tableData) {
            for (const key in this.tableData[index]) {
              if (typeof this.tableData[index][key] === 'object' || typeof this.tableData[index][key] === 'array') {
                this.tableData[index][key] = JSON.stringify(this.tableData[index][key])
              }
            }
          }
        }
      }
      this.$nextTick(() => {
        this.$refs.multipleTable.doLayout();
      })
    },

    async getMapping() {
      const input = {}
      input['es_connect'] = sdk.GetSelectEsConnID()
      input['index_name'] = this.indexName

      const { data, code, msg } = await ListAction(input)

      if (code == 0) {
        if (data.ver == 6) {
          const mappings = Object.keys(data.list[input['index_name']].mappings)
          for (const k in mappings) {
            if (mappings[k] == '_default_') {
              mappings.splice(k, 1)
              break
            }
          }
          const tableHeader = Object.keys(data.list[input['index_name']].mappings[mappings[0]].properties)
          this.properties = data.list[input['index_name']].mappings[mappings[0]].properties
          const tmpTableHeader = []
          tmpTableHeader.push('_id')
          for (const i in tableHeader) {
            if (tableHeader[i] != '_id') {
              tmpTableHeader.push(tableHeader[i])
            }
          }
          this.tableHeader = tmpTableHeader
        } else if (data.ver == 7 || data.ver == 8) {
          if(JSON.stringify(data.list[input['index_name']].mappings) == "{}"){
            ElMessage({
              type: 'error',
              message: "该索引没有设置映射结构"
            })
            return
          }
          const tableHeader = Object.keys(data.list[input['index_name']].mappings.properties)
          this.properties = data.list[input['index_name']].mappings.properties
          const tmpTableHeader = []
          tmpTableHeader.push('_id')
          for (const i in tableHeader) {
            if (tableHeader[i] != '_id') {
              tmpTableHeader.push(tableHeader[i])
            }
          }
          this.tableHeader = tmpTableHeader
        }
      } else {
        ElMessage({
          type: 'error',
          message: msg
        })
      }
    },
    pushSort() {
      this.sortList.push({
        col: '',
        sortRule: 'desc'
      })
    },
    deleteSort(index) {
      this.sortList.splice(index, 1)
    },

    handleSizeChange(v) {
      this.input.limit = v
      this.search(1)
    },
    async getdsl() {
      const form = {
        index_name: this.indexName,
        relation: this.whereFilter,
        sort_list: this.sortList,
        es_connect: sdk.GetSelectEsConnID(),
        page: this.input.page,
        limit: this.input.limit
      }
      const res = await GetDSL(form)
      if (res.code != 0) {
        ElMessage({
          type: 'error',
          message: res.msg
        })
        return
      }
      this.queryDsl = res.data.list
      this.queryDslShow = true
      this.$nextTick(()=>{
        this.showJsonEditor2 = true
      })
    },
    async search(page) {
      this.tableLoading = true
      !page ? this.input.page = 1 : this.input.page = page
      const form = {
        index_name: this.indexName,
        relation: this.whereFilter,
        sort_list: this.sortList,
        es_connect: sdk.GetSelectEsConnID(),
        page: this.input.page,
        limit: this.input.limit
      }
      const res = await GetList(form)
      if (res.code != 0) {
        ElMessage({
          type: 'error',
          message: res.msg
        })
        this.tableData = []
        this.tableLoading = false
        return
      }
      this.count = res.data.count

      await this.getMapping()

      if (this.count > 0) {
        await this.initTableData(res.data.list)
      } else {
        this.tableData = []
      }
      this.tableLoading = false
    }
  }
}
</script>

<style scoped>
.footadd___2D4YB {
  display: inline-block;
  margin-right: 16px;
  padding: 4px;
  color: #3d90ff;
  font-size: 13px;
  line-height: 20px;
  border-radius: 2px;
  cursor: pointer;
  transition: all .3s;
  transition-property: all;
  transition-duration: 0.3s;
  transition-timing-function: ease;
  transition-delay: 0s;
}

.row___xwl {
  min-height: 40px;
  padding: 0 4px 0 8px;
  transition-property: all;
  transition-duration: 0.3s;
  transition-timing-function: ease;
  transition-delay: 0s;
  padding: 10px;
}

.row___xwl:hover {
  box-shadow: 0 0 3px 0 #1890ff;
  transition-property: all;
  transition-duration: 0.3s;
  transition-timing-function: ease;
  transition-delay: 0s;
}

.action-row .action-left {
  display: flex;
  align-items: center;
}

.action-row .action-right {
  display: flex;
  align-items: center;
  width: 100px;
}

.actions_xwl_btn:hover {
  color: orangered;
}
::v-deep .el-table .sort-caret.descending{
  bottom: 0px;
}

</style>
