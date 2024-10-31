<template>
  <div>
    <el-row :style="'padding-left: ' + (newdeep - 1) * 15 + 'px; margin-top: 5px'">
      <template v-if="rootFlag">
        <div  class="item">
          <div
              :class="sizeClass"
              @click="hidden = !hidden"
          >
            <template  v-if="item.childs != null">
              <el-link :underline="false">
                <el-icon v-if="hidden"><ArrowRight></ArrowRight></el-icon>
                <el-icon v-else><ArrowDown></ArrowDown></el-icon>
              </el-link>
            </template>
            <template v-else>
              <i  style="margin-left: 14px" />
            </template>

          </div>
          <div class="item-cell" @click="hidden = !hidden">
            <el-link :underline="false">
              <el-tag :effect="'plain'" :class="tagClass">
                <span>{{ getName(item.key) ? getName(item.key) : item.key }}</span>
              </el-tag>
            </el-link>
          </div>
          <div class="item-key item-cell" :style="'width: ' + (300 - (newdeep - 1) * 15) + 'px'">
            <el-autocomplete
                v-model="item.key"
                clearable
                :fetch-suggestions="querySearch"
                :placeholder="parent.type == 'Array' ? '' : '请输入键'"
                :size="size"
                :disabled="item.isRoot || parent.type == 'Array'"
                :style="'width: ' + (290 - (newdeep - 1) * 15) + 'px'"
            >
              <template #default="{ item }">
                <span>{{ item.value }}-{{ item.data }}</span>
              </template>
            </el-autocomplete>
          </div>
          <div class="item-type item-cell">
            <el-select v-model="item.type" :size="size" :placeholder="'请选择'" class="select-body" @change="changeSelect">
              <el-option
                  v-for="type in item.isRoot ? rootOptions : options"
                  :key="type.value"
                  :label="type.label"
                  :value="type.value"
              />
            </el-select>
          </div>
          <div class="item-value item-cell">
            <el-autocomplete
                v-if="item.type != 'Number' && item.type != 'Boolean'"
                v-model="item.value"
                :fetch-suggestions="querySearch2"
                clearable
                :size="size"
                :placeholder="item.type == 'Array' || item.type == 'Object' ? '' : '请输入值'"
                :disabled="item.type == 'Array' || item.type == 'Object'"
                :class="widthClass"
            >
              <template #default="{ item }">
                <span>{{ item.value }}-{{ item.data }}</span>
              </template>
            </el-autocomplete>
            <el-input-number
                v-else-if="item.type == 'Number'"
                v-model="item.value"
                :size="size"
                :class="widthClass"
            />
            <el-radio-group v-else v-model="item.value" :class="widthClass">
              <el-radio class="el-radio" :label="true">是</el-radio>
              <el-radio class="el-radio" :label="false">否</el-radio>
            </el-radio-group>
          </div>
          <div :class="sizeClass">
            <el-tooltip
                v-if="item.type == 'Array' || item.type == 'Object'"
                class="item-control-cell"
                content="添加子元素"
                placement="top"
            >
              <el-link :underline="false" @click="addItem">
                <el-icon><Plus /></el-icon>
              </el-link>
            </el-tooltip>
            <el-popconfirm v-if="!item.isRoot" :title="'确定删除当前节点吗？'" @confirm="delItem">
              <el-link slot="reference" :underline="false" class="item-control-cell">
                <el-icon><Delete /></el-icon>
              </el-link>
            </el-popconfirm>
          </div>
        </div>
      </template>

    </el-row>
    <div :style="hidden ? 'display: none' : 'display: block'">
      <template v-if="item.childs && item.type == 'Object'">
        <span v-for="(child, index) in item.childs" :key="index">
          <VueJsonItem
              :size="size"
              :item="child"
              :parent="item"
              :names="names"
              :deep="newdeep"
              :open-flag="openFlag"
          />
        </span>
      </template>
      <template v-if="item.childs && item.type == 'Array'">
        <span v-for="(child, index) in item.childs" :key="index">
          <VueJsonItem
              :size="size"
              :item="child"
              :parent="item"
              :index="index"
              :names="names"
              :deep="newdeep"
              :open-flag="openFlag"
          />
        </span>
      </template>
    </div>
  </div>
</template>

<script setup>
import { reactive, ref } from 'vue';
import { filterData } from '@/utils/table';
import { Plus,Delete,ArrowRight,ArrowDown } from '@element-plus/icons-vue'


import VueJsonItem from '@/views/indices/components/Item.vue'


const props = defineProps({
  item: {
    type: Object,
    default: {
      key: '',
      value: '',
      type: ''
    }
  },
  names: {
    type: Array
  },
  parent: {
    type: Object
  },
  index: {
    type: Number,
    default: 0
  },
  deep: {
    type: Number,
    default: 0
  },
  size: {
    type: String,
    default: 'small'
  },
  rootFlag: {
    type: Boolean,
    default: true
  },
  openFlag: {
    type: Boolean,
    default: true
  }
});

const hidden = ref(!props.openFlag);
const newdeep = ref(props.deep + 1);
const max = 8;

const options = [
  { value: 'String', label: 'String' },
  { value: 'Object', label: 'Object' },
  { value: 'Array', label: 'Array' },
  { value: 'Number', label: 'Number' },
  { value: 'Boolean', label: 'Boolean' }
]

const rootOptions = [
  { value: 'Object', label: 'Object' },
  { value: 'Array', label: 'Array' }
];

const queryData = [
  { 'value': 'type', 'data': ' 数据类型' },
  { 'value': 'format', 'data': ' 时间格式化' },
  { 'value': 'analyzer', 'data': ' 分词器' },
  { 'value': 'normalizer', 'data': ' 分析器' },
  { 'value': 'boost', 'data': ' 权重' },
  { 'value': 'coerce', 'data': ' 强制类型转换' },

  { 'value': 'copy_to', 'data': ' 合并参数' },
  { 'value': 'doc_values', 'data': ' 文档值' },
  { 'value': 'dynamic', 'data': ' 动态设置' },
  { 'value': 'enabled', 'data': ' 是否开启字段' },
  { 'value': 'fielddata', 'data': ' 字段数据' },
  { 'value': 'ignore_above', 'data': ' 字段保存最大长度' },
  { 'value': 'ignore_malformed', 'data': ' 忽略格式不对的数据' },
  { 'value': 'include_in_all', 'data': ' _all 查询包含字段' },
  { 'value': 'index_options', 'data': ' 索引设置' },
  { 'value': 'index', 'data': ' 是否可以被搜索' },
  { 'value': 'fields', 'data': ' 字段' },

  { 'value': 'norms', 'data': ' 标准信息' },
  { 'value': 'null_value', 'data': ' 空值' },
  { 'value': 'position_increment_gap', 'data': ' 短语位置间隙' },
  { 'value': 'properties', 'data': ' 属性' },
  { 'value': 'search_analyzer', 'data': ' 搜索分析器' },
  { 'value': 'similarity', 'data': ' 匹配算法' },
  { 'value': 'store', 'data': ' 字段是否被存储' },
  { 'value': 'term_vector', 'data': ' 词根信息' }
];

const queryData2 = [
  { 'value': 'text', 'data': '字符串类型(可分词) ' },
  { 'value': 'keyword', 'data': '字符串类型(不可分词) ' },
  { 'value': 'long', 'data': '数字类型 ' },
  { 'value': 'integer', 'data': '数字类型 ' },
  { 'value': 'short', 'data': '数字类型 ' },
  { 'value': 'byte', 'data': '数字类型 ' },
  { 'value': 'double', 'data': '数字类型 ' },
  { 'value': 'float', 'data': '数字类型 ' },
  { 'value': 'half_float', 'data': '数字类型 ' },
  { 'value': 'scaled_float', 'data': '数字类型 ' },
  { 'value': 'date', 'data': '时间类型 ' },
  { 'value': 'boolean', 'data': '布尔类型 ' },
  { 'value': 'binary', 'data': '二进制类型 ' },
  { 'value': 'Array', 'data': '数组类型 ' },
  { 'value': 'object', 'data': '对象类型  ' },
  { 'value': 'nested', 'data': '嵌套类型 ' },
  { 'value': 'geo_point', 'data': '地理类型 ' },
  { 'value': 'geo_shape', 'data': '多边形类型 ' },
  { 'value': 'ip', 'data': 'ip类型 ' },
  { 'value': 'completion', 'data': '补全类型 ' },
  { 'value': 'token_count', 'data': '令牌计数类型 ' },
  { 'value': 'yyyy-MM-dd HH:mm:ss', 'data': '时间格式化 ' }
];

const iconClass = ['el-icon-arrow-right', 'el-icon-arrow-down'];
const tagClass = ['el-tag-blue', 'el-tag-orange'];
const sizeClass = ['item-cell item-control-cell', 'item-cell'];
const widthClass = ['el-input el-input--small'];

const querySearch = (queryString, cb) => {
  let queryDataTmp = JSON.parse(JSON.stringify(queryData))
  if (queryString == undefined) queryString = ''
  if (queryString.trim() == '') {
    if (queryDataTmp.length > max) {
      cb(queryDataTmp.slice(0, max))
    } else {
      cb(queryDataTmp)
    }
    return
  }

  queryDataTmp = filterData(queryDataTmp, queryString.trim())

  if (queryDataTmp.length > max) {
    cb(queryDataTmp.slice(0, max))
  } else {
    cb(queryDataTmp)
  }
}

const querySearch2 = (queryString, cb) => {
  let queryDataTmp = JSON.parse(JSON.stringify(queryData2))
  if (queryString == undefined) queryString = ''
  if (queryString.trim() == '') {
    if (queryDataTmp.length > max) {
      cb(queryDataTmp.slice(0, max))
    } else {
      cb(queryDataTmp)
    }
    return
  }

  queryDataTmp = filterData(queryDataTmp, queryString.trim())

  if (queryDataTmp.length > max) {
    cb(queryDataTmp.slice(0, max))
  } else {
    cb(queryDataTmp)
  }
}

const getName = (key) => {
  const result = queryData.find(item => item.value === key);
  return result ? result.data : '';
}

const changeSelect = () => {
  if (['Object', 'Array'].includes(props.item.type)) {
    hidden.value = false;
  }
}

const addItem = () => {
  if (!props.item.childs) {
    props.item.childs = [];
  }
  props.item.childs.push({ key: '', value: '', type: '' });
}

const delItem = () => {
  const index = props.parent.childs.indexOf(props.item);
  if (index > -1) {
    props.parent.childs.splice(index, 1);
  }
}
</script>

<style scoped>
.item {
  display: flex;
  align-items: center;
}

.item-cell {
  padding: 0 5px;
}

.item-key {
  width: 290px;
}

.item-type {
  width: 120px;
}

.item-value {
  width: 250px;
}

.select-body {
  width: 100%;
}

.el-icon-blue {
  color: #409eff;
}

.el-icon-dim {
  color: #c0c4cc;
}

.el-tag-blue {
  background-color: #409eff;
}

.el-tag-orange {
  background-color: cornsilk;
}
</style>
