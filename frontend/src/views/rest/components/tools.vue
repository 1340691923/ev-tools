<template>
  <div>
    <div class="search-container">
      <el-form>
        <el-form-item>
          <el-container >
            <el-autocomplete
                style="width: 30rem"
                ref="autocomplete"
                v-model="input.path"
                clearable
                class="filter-item select-path autocomplete"
                :placeholder="$t('请输入内容')"
                :fetch-suggestions="querySearch"
                @clear="clear"
                @keyup.enter.native="go"
                @select="mySelect"
            >
              <template #prepend>
                <el-select

                    v-model="input.method"
                    style="width:10rem"
                    :placeholder="$t('请选择Http Method')"
                    filterable
                >
                  <el-option :label="$t('PUT【更新或创建】')" value="PUT"/>
                  <el-option :label="$t('GET【查询】')" value="GET"/>
                  <el-option :label="'DELETE【'+$t('删除')+'】'" value="DELETE"/>
                  <el-option :label="$t('POST【创建】')" value="POST"/>
                  <el-option :label="$t('HEAD【是否存在】')" value="HEAD"/>
                </el-select>
              </template>

            </el-autocomplete>
            <el-button slot="append"  type="text">
              <a href="https://www.elastic.co/guide/en/elasticsearch/reference/current/index.html" target="_blank">{{ $t('官方文档') }}</a>
            </el-button>
            <el-button
                class="filter-item go"
                style="display: inline;"
                :loading="loading"
                type="primary"
                @click="go"
            >GO</el-button>
            <el-button
                v-show="input.method == 'GET'"
                class="filter-item sql-format"
                style="display: inline;"
                type="warning"
                @click="openDrag"
            >{{ $t('SQL转换') }}
            </el-button>
            <el-button
                class="filter-item search-history"
                style="display: inline;"
                type="danger"

                @click.native="dialogVisible = true"
            >
              {{ $t('搜索历史') }}
            </el-button>
            <el-button
                v-if="canExport && input.path.trim() != '_search'"

                style="display: inline;"
                type="success"

                @click.native="tableDialogVisible = true"
            >
              {{ $t('返回结果转表格') }}
            </el-button>

            <download-excel
                v-if="canExport"
                ref="download"
                class="download"
                :fields="json_fields"
                :data="json_data"
                :name="$t(input.path+'.xls')"
                :before-generate="startDownload"
                :before-finish="endDownload"
            >
              <el-button
                  :disabled="downloadLoading"
                  v-loading="downloadLoading"
                  type="primary"
                  class="filter-item download"
              >
                {{ $t('下载') }}
              </el-button>
            </download-excel>

          </el-container>
        </el-form-item>
      </el-form>


    </div>

    <page-split style="margin-top:10px"
                @resizeLineEndMove="onResizeLineMove"
                @resizeLineMove="onResizeLineMove"
                :distribute="0.5" :lineThickness="6" :isVertical="true"  >
      <template v-slot:first>

        <json-editor

            ref="jsonEditorRef1"
            v-model:value="input.body"
            font-size="15"
            height="500"
            class="req-body"
            :point-out="pointOut"
            :read="false"
            :title="$t('请求Body')"

        />
      </template>
      <template v-slot:second>
        <json-editor
            ref="jsonEditorRef2"
            v-model:value="resData"
            font-size="15"
            height="500"
            class="res-body"
            :read="true"
            :title="$t('返回信息')"
        />
      </template>
    </page-split>
    <el-drawer
        ref="drawer"
        title="Edit SQL"
        :before-close="drawerHandleClose"
        v-model="drawerShow"
        direction="rtl"
        close-on-press-escape
        destroy-on-close
        size="80%"
    >
      <el-button
          style="margin: 20px"
          type="warning"
          @click="sqlToDsl"
      >
        {{ $t('开始转换为DSL') }}
      </el-button>
      <el-link type="success" disabled>{{ $t('表名可用索引名代替') }}</el-link>
      <sql-editor
          :point-out="props.indices"
          v-if="drawerShowSqlEditor"
          v-model:value="sqlStr"
          styles="width: 100%"
      />
    </el-drawer>

    <history
        v-if="dialogVisible"
        :dialog-visible="dialogVisible"
        @getHistoryData="getHistoryData"
        @close="closeHistory"
    />

    <res-table
        v-if="tableDialogVisible"
        :search-path="input.path"
        :dialog-visible="tableDialogVisible"
        :json-data="JSON.parse(resData)"
        @close="closeResTable"
    />
  </div>
</template>

<script setup>
import {ref, computed, onMounted, onBeforeUnmount, getCurrentInstance, nextTick} from 'vue'
import {sdk} from "@/plugin_sdk/sdk";
import SqlEditor from '@/components/SqlEditor/index.vue'
import JsonEditor from '@/components/JsonEditor/index.vue'
import History from '@/views/rest/components/history.vue'
import ResTable from '@/views/rest/components/res-table.vue'

import PageSplit  from "vue3-page-split";
import "vue3-page-split/dist/style.css";

import {ElMessage,ElMessageBox} from "element-plus";



import { clone } from '@/utils/index';
import { RunDslAction, SqlToDslAction } from '@/api/es';
import { filterData } from '@/utils/table';
import { esBodyKeyWords } from '@/utils/base-data';
const ctx = getCurrentInstance().appContext.config.globalProperties
const props = defineProps({
  queryData: {
    type: Array,
    default: () => []
  },
  indices:{
   type:Array,
   default:()=>[]
  },
  uniqueId: {
    type: String,
    default: ''
  },
  input: {
    type: Object,
    default: () => ({})
  },
  sqlStr: {
    type: String,
    default: ''
  },
  title: {
    type: String,
    default: ''
  },
  max: {
    type: Number,
    default: 8
  }
});
const sqlStr = ref(props.sqlStr)
const tableDialogVisible = ref(false);
const cancelToken = ref('');
const dialogVisible = ref(false);
const jsonEditorRef1 = ref()
const jsonEditorRef2 = ref()
const drawerShow = ref(false);
const drawerShowSqlEditor = ref(false);
const loading = ref(false);
const json_fields = ref({});
const json_data = ref('');
const pointOut = ref(esBodyKeyWords);
const resData = ref('{}');
const downloadLoading = ref(false);
const input = ref(props.input)

const startDownload = ()=>{
  ElMessage({
    type: 'success',
    message: '开始下载'
  })
}

const endDownload = ()=>{
  ElMessage({
    type: 'success',
    message: '下载成功'
  })
}


const canExport = computed(() => {
  json_data.value = '';
  json_fields.value = {};
  const resDataObj = JSON.parse(resData.value);
  if (!resDataObj) return false;

  if (Array.isArray(resDataObj)) {
    if (resDataObj.length <= 0) return false;
    json_data.value = replaceArrSpece(resDataObj);
    Object.keys(resDataObj[0]).forEach((key) => {
      json_fields.value[key] = key;
    });
    return true;
  } else {
    if (resDataObj.hasOwnProperty('hits')) {
      if (resDataObj['hits']['hits'].length > 0) {
        const jsonData = resDataObj['hits']['hits'];
        const defaultKeys = ['_index', '_type', '_id'];
        defaultKeys.forEach((key) => {
          json_fields.value[key] = key;
        });
        const arrayColumns = [];
        jsonData.forEach((v) => {
          const sourceMap = v['_source'];
          if (sourceMap == null) return;
          Object.keys(sourceMap).forEach((sourceVal) => {
            if (Object.prototype.toString.call(sourceMap[sourceVal]) === '[object Object]') {
              Object.keys(sourceMap[sourceVal]).forEach((key) => {
                json_fields.value[`${sourceVal}->${key}`] = `_source.${sourceVal}.${key}`;
              });
            } else if (Array.isArray(sourceMap[sourceVal])) {
              if (Object.prototype.toString.call(sourceMap[sourceVal][0]) === '[object Object]') {
                arrayColumns.push(sourceVal);
              }
              json_fields.value[sourceVal] = `_source.${sourceVal}`;
            } else {
              json_fields.value[sourceVal] = `_source.${sourceVal}`;
            }
          });
        });
        arrayColumns.forEach((arrayColumn) => {
          jsonData.forEach((item) => {
            if (item['_source'][arrayColumn]) {
              item['_source'][arrayColumn] = JSON.stringify(item['_source'][arrayColumn]);
            }
          });
        });
        json_data.value = jsonData;
        return true;
      }
    }
  }
  return false;
});

const startGuid = async () => {

};

const drawerHandleClose = (done) => {
  drawerShowSqlEditor.value = false
  done()
}

const openDrag = () => {
  drawerShow.value = true;
  nextTick(()=>{
    drawerShowSqlEditor.value = true
  })
};

const clear = () => {
  input.value.body = '{}';
  resData.value = '{}';
  jsonEditorRef1.value.SetText('{}')
  jsonEditorRef2.value.SetText('{}')
};

const mySelect = (obj) => {
  input.value.path = obj.data
}

const MeetingConfirmBox = (title) => {
  return new Promise((resolve, reject) => {
    ElMessageBox.confirm(title, '警告', {
      confirmButtonText: '确认',
      cancelButtonText: '取消',
      type: 'warning'
    })
        .then(() => resolve(true))
        .catch(err => resolve(false))
  })
}

const go = async () => {
  const inputGo = clone(input.value)

  if (inputGo['method'] == 'DELETE' || inputGo['path'].indexOf('_delete_by_query') != -1) {
    const isFinish = await MeetingConfirmBox('确定执行删除操作吗')
    if (!isFinish) {
      ElMessage({
        type: 'success',
        message: '已取消'
      })
      return
    }
  }

  if (inputGo['path'].trim().length > 0 || inputGo['path'].trim() == '') {
    if (inputGo['path'].trim().substr(0, 1) != '/') {
      inputGo['path'] = '/' + inputGo['path'].trim()
    }
  }

  loading.value = true

  inputGo['es_connect'] = sdk.GetSelectEsConnID()

  RunDslAction(inputGo).then(res => {
    loading.value = false
    if (res.code == 0 || res.code == 200 || res.code == 201) {
      ElMessage({
        type: 'success',
        message: res.msg
      })
    } else {

      if(res.data != null && res.data.hasOwnProperty("error") &&
          res.data["error"].hasOwnProperty("reason") &&
          res.data["error"].reason.indexOf("does not support having a body") !== -1
      ){
        ElMessage({
          type: 'error',
          message: "请清空body，然后重新点击【->GO】获取数据"
        })
      }else{

        ElMessage({
          type: 'error',
          message: res.msg
        })
      }

    }


    resData.value = JSON.stringify(res.data, null, '\t')
    jsonEditorRef2.value.SetText(resData.value)
  }).catch(err => {
    console.log(err)
    loading.value = false

  })
}

const onResizeLineMove = ()=>{
  jsonEditorRef1.value.updateEditorWidth()
  jsonEditorRef2.value.updateEditorWidth()
}


const cancelReq = () => {
  cancelToken.value();
  cancelToken.value = '';
};



const sqlToDsl = async () => {

  const {data, code, msg} = await SqlToDslAction({sql: sqlStr.value})
  console.log(data,msg)
  if (code == 0) {
    input.value.body = data.dsl
    input.value.path = data.tableName + '/_search'
    jsonEditorRef1.value.SetText(data.dsl)
    go()
  } else {
    ElMessage({
      type: 'error',
      message: msg
    })

  }

};

const getSql = (value) => {
  sqlStr.value = value;
};

const getHistoryData = (history) => {
  input.value.path = history.path;
  input.value.body = history.body;
  jsonEditorRef1.value.SetText(history.body)
  input.value.method = history.method;
  go();
};

const closeHistory = () => {
  dialogVisible.value = false;
};

const closeResTable = () => {
  tableDialogVisible.value = false;
};

const replaceArrSpece = (arr) => {
  const result = [];
  for (const item of arr) {
    const newObj = {};
    for (const key in item) {
      const newKey = key.replace(/^\s+|\s+$/g, '');
      newObj[newKey] = item[key];
    }
    result.push(newObj);
  }
  return result;
};

const querySearch = (queryString, cb) => {
  let queryDataSearch = JSON.parse(JSON.stringify(props.queryData))
  if (queryString.trim() == '') {
    if (queryDataSearch.length > props.max) {
      cb(queryDataSearch.slice(0, props.max))
    } else {
      cb(queryDataSearch)
    }
    return;
  }

  queryDataSearch = filterData(queryDataSearch, queryString.trim())

  if (queryDataSearch.length > props.max) {
    cb(queryDataSearch.slice(0, props.max))
  } else {
    cb(queryDataSearch)
  }
}

onMounted(() => {
  startGuid();
});

onBeforeUnmount(() => {
});
</script>

<style scoped>
/* Add your styles here */
</style>
