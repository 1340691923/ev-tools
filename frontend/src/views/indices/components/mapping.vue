<template>
  <div>
    <el-drawer

        v-model="open"
        :title="title.concat(`【${indexName}】`)"
        size="80%"
        @close="closeDialog"
    >
      <div class="app-container">
        <div class="search-container">
          <el-form :inline="true">
            <template v-if="showTypeName">
              <el-form-item  label="type">
                <el-input
                    v-model="type_name"
                    :readonly="typeReadonly"
                    style="width: 200px"
                />
              </el-form-item>
            </template>

            <el-form-item label="dynamic">
              <el-select style="width:100px" v-model="dynamic" >
                <el-option :label="$t('动态映射')" value="true" />
                <el-option :label="$t('静态映射')" value="false" />
                <el-option :label="$t('严格映射')" value="strict" />
              </el-select>
            </el-form-item>
            <el-form-item>
              <el-button
                  v-loading="loading"
                  :disabled="loading"
                  type="success"

                  @click="saveMappinng"
              >
                {{ $t('保存/修改映射') }}
              </el-button>
            </el-form-item>
          </el-form>

        </div>

        <el-tabs>

        </el-tabs>
        <el-tabs  v-model="activeName"  stretch >
          <el-tab-pane  name="1" :lazy="true" :label="$t('表单')">
            <div v-loading="!showVueJsonHelper" >
              <template v-if="showVueJsonHelper">
                <vue-json-helper

                    :size="size"
                    :names="names"
                    :json-str="jsonStr"
                    :root-flag="rootFlag"
                    :open-flag="openFlag"
                    :back-top-flag="backTopFlag"
                    :shadow-flag="false"
                    :border-flag="false"
                    @jsonListener="jsonListener"
                />
              </template>

            </div>
          </el-tab-pane>
          <el-tab-pane name="2" :lazy="true" :label="$t('JSON')">

            <div>
              <json-editor
                  v-model:value="jsonStrData"
                  height="720"
                  styles="width: 100%"
                  :title="type_name"
              />
            </div>
          </el-tab-pane>
        </el-tabs>

      </div>
    </el-drawer>
  </div>
</template>

<script setup>
import {ref, onMounted, nextTick, getCurrentInstance,computed} from 'vue';
import { ElMessage } from 'element-plus';
import { ListAction, UpdateMappingAction } from '@/api/es-map';
import JsonEditor from '@/components/JsonEditor/index.vue'

import VueJsonHelper from '@/views/indices/components/Helper.vue'
import {sdk} from "@elasticview/plugin-sdk";
// Define props
const props = defineProps({
  indexName: {
    type: String,
    default: ''
  },
  open: {
    type: Boolean,
    default: false
  },
  title: {
    type: String,
    default: '新增映射结构'
  }
});

const open = ref(props.open)
const title = ref(props.title)
const indexName = ref(props.indexName)
const  activeName = ref('1')

const ctx = getCurrentInstance().appContext.config.globalProperties

const jsonStrData = computed(()=>{
  return JSON.stringify(JSON.parse(jsonStr.value), null, '\t')
})

// Define emit
const emit = defineEmits(['close', 'finished']);

// Reactive data and state
const loading = ref(false);
const dynamic = ref('false');
const size = ref('small');
const names = [
  { key: 'Root', name: 'properties' },
  { key: 'type', name: '数据类型' },
  { key: 'format', name: '时间格式化' },
  { key: 'analyzer', name: '分词器' },
  { key: 'normalizer', name: '分析器' },
  { key: 'boost', name: '权重' },
  { key: 'coerce', name: '强制类型转换' },
  { key: 'copy_to', name: '合并参数' },
  { key: 'doc_values', name: '文档值' },
  { key: 'dynamic', name: '动态设置' },
  { key: 'enabled', name: '是否开启字段' },
  { key: 'fielddata', name: '字段数据' },
  { key: 'ignore_above', name: '字段保存最大长度' },
  { key: 'ignore_malformed', name: '忽略格式不对的数据' },
  { key: 'include_in_all', name: '_all 查询包含字段' },
  { key: 'index_options', name: '索引设置' },
  { key: 'index', name: '是否可以被搜索' },
  { key: 'fields', name: '字段' },
  { key: 'norms', name: '标准信息' },
  { key: 'null_value', name: '空值' },
  { key: 'position_increment_gap', name: '短语位置间隙' },
  { key: 'properties', name: '属性' },
  { key: 'search_analyzer', name: '搜索分析器' },
  { key: 'similarity', name: '匹配算法' },
  { key: 'store', name: '字段是否被存储' },
  { key: 'term_vector', name: '词根信息' }
];
const rootFlag = ref(true);
const openFlag = ref(true);
const backTopFlag = ref(false);
const jsonStr = ref('{}');
const showVueJsonHelper = ref(false);
const type_name = ref('');
const ver = ref(6);
const showTypeName = ref(false);
const typeReadonly = ref(false);

// Methods
const saveMappinng = async () => {
  let properties = {};
  try {
    properties = JSON.parse(jsonStr.value);
  } catch (e) {
    ElMessage({
      type: 'error',
      message: 'JSON格式不正确'
    });
    return;
  }

  const input = {
    es_connect: sdk.GetSelectEsConnID(),
    index_name: props.indexName
  };
  const activeData = {
    properties,
    dynamic: dynamic.value
  };

  switch (ver.value) {
    case 6:
      input['properties'] = activeData;
      input['type_name'] = type_name.value;
      break;
    case 7:
    case 8:
      input['properties'] = activeData;
      break;
  }

  loading.value = true;
  const { data, code, msg } = await UpdateMappingAction(input);
  loading.value = false;

  if (code === 0) {
    ElMessage({
      type: 'success',
      message: msg
    });
    emit('finished');
  } else {
    ElMessage({
      type: 'error',
      message: msg
    });
  }
};



const init = async () => {
  const input = {
    es_connect: sdk.GetSelectEsConnID(),
    index_name: props.indexName
  };

  const { data, code, msg } = await ListAction(input);

  if (code !== 0) {
    ElMessage({
      type: 'error',
      message: msg
    });
    return;
  }

  ver.value = data.ver;

  switch (ver.value) {
    case 6:
      const mappings = Object.keys(data.list[props.indexName].mappings);
      showTypeName.value = true;

      if (mappings.length === 0) {
        typeReadonly.value = false;
      } else {
        type_name.value = mappings[0];
        typeReadonly.value = true;
        dynamic.value = data.list[props.indexName].mappings[type_name.value]?.dynamic ?? 'false';
        jsonStr.value = JSON.stringify(data.list[props.indexName].mappings[type_name.value]?.properties ?? '{}');
      }
      break;
    case 7:
    case 8:
      dynamic.value = data.list[props.indexName].mappings?.dynamic ?? 'false';
      jsonStr.value = JSON.stringify(data.list[props.indexName].mappings?.properties ?? '{}');
      break;
  }
  refreshVueJsonHelper();
};

const jsonListener = (data) => {
  jsonStr.value = JSON.stringify(data);
};

const closeDialog = () => {
  emit('close');
};

const refreshVueJsonHelper = () => {
  showVueJsonHelper.value = false;
  nextTick(() => {
    showVueJsonHelper.value = true;
  });
};

// Mounted hook
onMounted(() => {
  init();
});


</script>

<style scoped>
/* Add your styles here */
</style>
