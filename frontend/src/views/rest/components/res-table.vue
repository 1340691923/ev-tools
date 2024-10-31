<template>
  <div>
    <el-drawer  size="95%" v-model="dialogVisible" title="查询结果" @close="close">
      <div class="search-container">
        <el-tag >{{ $t('请输入关键词') }}</el-tag>
        <el-input v-model="input"  style="width: 300px" clearable @input="search" />
        <el-button type="success"  @click="search">{{ $t('搜索') }}</el-button>
        <el-button
            v-if="ISDoc"
            type="primary"

            @click="openAddDialog = true"
        >
          {{ $t('添加文档') }}
        </el-button>
      </div>

      <el-table
          v-if="showTable"
          :data="getTableData"
          use-virtual
          :row-height="30"
      >
        <el-table-column label="ID" align="center" fixed min-width="100">
          <template #default="{ row, $index }">
            <template v-if="ISDoc">
              {{ row["_id"] }}
            </template>
            <template v-else>
              {{ $index + 1 }}
            </template>
          </template>
        </el-table-column>

        <el-table-column
            v-for="(val, key) in tableHeader"
            v-if="val !== 'xwl_index' && val !== '_id'"
            :key="key"
            :prop="val"
            :sortable="true"
            align="center"
            :label="val"
        >
          <template #default="{ row }">
            <el-popover
                v-if="row[val] !== undefined"
                placement="top-start"
                trigger="hover"
            >
              <div>{{ row[val].toString() }}</div>
              <template #reference>
                <span
                    v-if="row[val].toString().length >= 20"

                >
                {{ row[val].toString().substr(0, 20) + "..." }}
              </span>
                <span v-else >{{ row[val].toString() }}</span>
              </template>

            </el-popover>
          </template>
        </el-table-column>

        <el-table-column align="center" :label="$t('操作')" fixed="right" width="200">
          <template #default="{ row, $index }">
            <el-button-group>
              <el-button
                  v-if="ISDoc"
                  type="primary"
                  @click="look($index)"
              >
                {{ $t('编辑') }}
              </el-button>
              <el-button
                  v-else
                  type="success"
                  @click="look($index)"
              >
                {{ $t('查看') }}
              </el-button>
              <el-button
                  v-if="ISDoc"
                  type="danger"
                  @click="deleteByID(row, $index)"
              >
                {{ $t('删除') }}
              </el-button>
            </el-button-group>
          </template>
        </el-table-column>
      </el-table>
    </el-drawer>

    <el-drawer
        ref="drawer"
        :title="$t('详细数据')"
        :before-close="drawerHandleClose"
        v-model="drawerShow"
        direction="rtl"
        close-on-press-escape
        destroy-on-close
        size="80%"
    >
      <div v-if="ISDoc" class="search-container">
        <el-tag >{{ $t('操作') }}</el-tag>
        <el-button
            type="primary"

            @click="updateByID"
        >
          {{ $t('修改') }}
        </el-button>
      </div>

      <json-editor
          v-if="Array.isArray(jsonData)"
          v-model:value="jsonDataJSON"
          height="900"
          class="res-body"
          styles="width: 100%"
          :read="true"
          :title="$t('详细数据')"
      />

      <json-editor
          v-else
          v-model:value="HitsJson"
          height="900"
          class="res-body"
          styles="width: 100%"
          :read="!ISDoc"
          :title="$t('详细数据')"
          @getValue="getEditDoc"
      />
    </el-drawer>

    <el-drawer
        ref="drawer"
        :title="$t('新增文档')"
        :before-close="drawerHandleClose"
        v-model="openAddDialog"
        direction="rtl"
        close-on-press-escape
        destroy-on-close
        size="80%"
    >
      <div class="search-container">
        <el-tag >{{ $t('操作') }}</el-tag>
        <el-button
            type="primary"

            @click="add"
        >
          {{ $t('提交') }}
        </el-button>
      </div>

      <json-editor
          v-if="openAddDialog"
          v-model:value="propertiesJSON"
          height="900"
          class="res-body"
          styles="width: 100%"
          :title="$t('新增文档')"
          @getValue="getNewDoc"
      />
    </el-drawer>
  </div>
</template>

<script setup lang="ts">
import {ref, reactive, computed, onMounted, getCurrentInstance, nextTick} from 'vue';
import { ElMessage, ElMessageBox } from 'element-plus';
import { filterData } from '@/utils/table';
import { clone } from '@/utils/index';
import { DeleteRowByIDAction, InsertAction, UpdateByIDAction } from '@/api/es-doc';
import { ListAction } from '@/api/es-map';
import JsonEditor from '@/components/JsonEditor/index.vue';
import {sdk} from "@/plugin_sdk/sdk";
const props = defineProps({
  dialogVisible: Boolean,
  jsonData: Array,
  searchPath: String,
});

const emit = defineEmits(['close']);

const ctx = getCurrentInstance().appContext.config.globalProperties

const jsonData = ref(props.jsonData)

const dialogVisible = ref(props.dialogVisible)

const openAddDialog = ref(false);
const showTable = ref(true);
const drawerShow = ref(false);
const tableData = ref([]);
const index = ref(0);
const isArray = ref(false);
const tableHeader = ref([]);
const ISDoc = ref(false);
const input = ref('');
const properties = ref({});
const newDoc = ref({});

const getTableData = computed(() => {
  if (input.value === '') {
    return tableData.value;
  } else {
    return filterData(tableData.value, input.value.trim());
  }
});

const search = () => {
  showTable.value = false;
  nextTick(() => {
    showTable.value = true;
  });
};

const propertiesJSON = computed(()=>{
  return JSON.stringify(properties.value, null, '\t')
})

const HitsJson = computed(()=>{
  return JSON.stringify(jsonData.value['hits']['hits'][index], null, '\t')
})

const jsonDataJSON = computed(()=>{
  return JSON.stringify(jsonData.value[index], null, '\t')
})

const initTableData = async () => {
  const resData = clone(jsonData.value);
  if (Array.isArray(resData)) {
    if (resData.length > 500) {
      ElMessage.error('请减少数据条数');
      emit('close', false);
      return;
    }
    tableData.value = resData;

    for (const index in tableData.value) {
      const obj = tableData.value[index];

      for (const key in obj) {
        let value = strToNum(obj[key]);
        if (value === false) {
          value = obj[key];
        }
        if (key.includes('.')) {
          tableData.value[index][key.split('.').join('->')] = value;
          delete tableData.value[index][key];
        } else {
          tableData.value[index][key] = value;
        }

        tableData.value[index]['xwl_index'] = index;
      }
    }

    tableHeader.value = Object.keys(tableData.value[0]);
    isArray.value = true;
  } else {
    if (resData !=undefined && resData.hasOwnProperty('hits')) {
      if (resData['hits']['hits'].length > 0) {
        ISDoc.value = true;
        if (resData['hits']['hits'].length > 1000) {
          ElMessage.error('请减少查詢的数据条数');
          emit('close', false);
          return;
        }

        const sourceArr = [];

        for (const index in resData['hits']['hits']) {
          const _source = resData['hits']['hits'][index]['_source'];
          _source['_id'] = resData['hits']['hits'][index]['_id'];
          _source['_score'] = resData['hits']['hits'][index]['_score'];
          _source['xwl_index'] = index;
          sourceArr.push(_source);
        }

        isArray.value = false;
        tableData.value = sourceArr;

        for (const index in tableData.value) {
          for (const key in tableData.value[index]) {
            if (typeof tableData.value[index][key] === 'object' || Array.isArray(tableData.value[index][key])) {
              tableData.value[index][key] = JSON.stringify(tableData.value[index][key]);
            }
          }
        }
      }

      const input = {
        es_connect: sdk.GetSelectEsConnID(),
        index_name: resData['hits']['hits'][0]['_index']
      };

      const { data, code, msg } = await ListAction( input);

      if (code === 0) {
        if (data.ver === 6) {
          const mappings = Object.keys(data.list[input.index_name].mappings);
          const tableHeaderValue = Object.keys(data.list[input.index_name].mappings[mappings[0]].properties);
          properties.value = data.list[input.index_name].mappings[mappings[0]].properties;
          const tmpTableHeader = ['_id', ...tableHeaderValue.filter(header => header !== '_id')];
          tableHeader.value = tmpTableHeader;
        } else if (data.ver === 7 || data.ver === 8) {
          const tableHeaderValue = Object.keys(data.list[input.index_name].mappings.properties);
          properties.value = data.list[input.index_name].mappings.properties;
          const tmpTableHeader = ['_id', ...tableHeaderValue.filter(header => header !== '_id')];
          tableHeader.value = tmpTableHeader;
        }
      } else {
        ElMessage.error(msg);
      }
    }
  }
};

const getNewDoc = (doc) => {
  try {
    newDoc.value = JSON.parse(doc);
  } catch (e) {
    console.error(e);
  }
};

const add = async () => {
  const editData = jsonData.value['hits']['hits'][index.value];
  const doc = newDoc.value;

  const input = {
    es_connect: sdk.GetSelectEsConnID(),
    index: editData['_index'],
    type_name: editData['_type'],
    json: doc
  };

  const res = await InsertAction( input);
  if (res.code === 0) {
    ElMessage.success(res.msg + '(_id为:' + res.data._id + ')');
  } else {
    ElMessage.error(res.msg);
  }
};

const getEditDoc = (v) => {
  try {
    const editDoc = JSON.parse(v);
    jsonData.value['hits']['hits'][index.value] = editDoc;
  } catch (e) {
    console.error(e);
  }
};

const updateByID = async () => {
  const editData = jsonData.value['hits']['hits'][index.value];
  const doc = editData['_source'];

  const input = {
    es_connect: sdk.GetSelectEsConnID(),
    index: editData['_index'],
    type_name: editData['_type'],
    id: editData['_id'],
    json: doc
  };

  const res = await UpdateByIDAction( input);
  if (res.code === 0) {
    const _source = clone(jsonData.value['hits']['hits'][index.value]['_source']);
    _source['_id'] = jsonData.value['hits']['hits'][index.value]['_id'];
    _source['_score'] = jsonData.value['hits']['hits'][index.value]['_score'];
    _source['xwl_index'] = index.value;
    for (const key in _source) {
      if (typeof _source[key] === 'object' || Array.isArray(_source[key])) {
        _source[key] = JSON.stringify(_source[key]);
      }
    }
    tableData.value[index.value] = _source;
    showTable.value = false;
    nextTick(() => {
      showTable.value = true;
    });
    ElMessage.success(res.msg);
  } else {
    ElMessage.error(res.msg);
  }
};

const deleteByID = async (row, index) => {
  ElMessageBox.confirm('确定删除该条文档吗?', '警告', {
    confirmButtonText: '确认',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    const input = {
      es_connect: evUtils.GetSelectEsConnID(),
      index_name: jsonData.value['hits']['hits'][0]['_index'],
      type: jsonData.value['hits']['hits'][0]['_type'],
      id: row['_id']
    };

    const res = await DeleteRowByIDAction(props, input);
    if (res.code === 0) {
      tableData.value.splice(index, 1);
      ElMessage.success(res.msg);
    } else {
      ElMessage.error(res.msg);
    }
  }).catch(err => {
    console.error(err);
  });
};

const strToNum = (str) => {
  const convertNum = Number(str);
  if (str === '') return false;
  if (str.includes(' ')) return false;
  if (isNaN(convertNum)) return false;
  return convertNum;
};

const look = (i) => {
  index.value = i;
  drawerShow.value = true;
};

const drawerHandleClose = (done) => {
  done();
};

const close = () => {
  emit('close', false);
};

onMounted(() => {
  initTableData();
});
</script>

<style scoped>
 .el-table .sort-caret.descending {
  bottom: 0px;
}
</style>
