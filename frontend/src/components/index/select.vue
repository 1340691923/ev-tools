<template>
  <el-select

      v-model="indexName"
      :reserve-keyword="multiple"
      :collapse-tags="multiple"
      :disabled="disabled"
      :placeholder="placeholder"
      clearable
      :multiple="multiple"
      :clearable="clearable"
      filterable
      @change="handleChange"
  >
    <el-option v-if="haveAll" label="全部" value="*" />
    <el-option
        v-for="(name, index) in indexList"
        :key="index"
        :label="name"
        :value="name"
    />
  </el-select>
</template>

<script setup>
import {ref, onMounted, getCurrentInstance} from 'vue'
import { IndexNamesAction } from '@/api/es-index'
import {ElMessage} from "element-plus";
import {sdk} from "@elasticview/plugin-sdk"

const ctx = getCurrentInstance().appContext.config.globalProperties

const props = defineProps({
  indexName: {
    type: [String, Array],
    default: ''
  },
  multiple: {
    type: Boolean,
    default: false
  },
  clearable: {
    type: Boolean,
    default: false
  },
  disabled: {
    type: Boolean,
    default: false
  },
  placeholder: {
    type: String,
    default: ''
  },
  haveAll: {
    type: Boolean,
    default: false
  }
})

const emit = defineEmits(['change'])

const indexName = ref(props.indexName)
const indexList = ref([])

onMounted(() => {
  getIndexList()
})

const handleChange = () => {
  emit('change', indexName.value)
}

const getIndexList = () => {
  const input = { es_connect: sdk.GetSelectEsConnID() }
  IndexNamesAction(input)
      .then(res => {
        if (res.code === 0) {
          indexList.value = res.data
        } else {
          ElMessage({
            type: 'error',
            message: res.msg
          })
        }
      })
      .catch(err => {
        console.log(err)
      })
}
</script>

<style scoped>
/* Your styles here */
</style>
