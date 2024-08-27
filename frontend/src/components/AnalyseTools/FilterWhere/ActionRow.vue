<template>
  <div>
    <div class="ta-filter-condition">

      <el-select
          v-model="v.columnName"
          filterable
          style="width: 300px"
          placeholder="请选择"
          @change="changeColumnNameSelect"
      >
        <el-option-group
            v-for="group in options"
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

      <el-select v-model="v.comparator" style="width: 80px">
        <el-option
            v-for="(v, k, index) in getDataTypeCalcuSymbol(v.columnName)"
            :key="index"
            :label="v"
            :value="k"
        />
      </el-select>

      <template v-if="noValueSymbolArr.includes(v.comparator)" />
      <template v-else-if="rangeSymbolArr.includes(v.comparator)">
        <el-input v-model="v.ftv[0]" clearable type="number" style="width: 150px" />
        ~
        <el-input v-model="v.ftv[1]" clearable type="number" style="width: 150px" />
      </template>

      <el-input
          v-else-if="inputSymbolArr.includes(v.comparator)"
          v-model="v.ftv"
          clearable
          style="width: 300px"
      />

      <template v-else-if="rangeTimeSymbolArr.includes(v.comparator)">
        <el-date-picker
            v-model="v.ftv"
            type="datetimerange"
            align="right"
            :default-time="['00:00:00', '23:59:59']"
            unlink-panels
            range-separator="至"
            start-placeholder="开始日期"
            end-placeholder="结束日期"
            format="yyyy-MM-dd HH:mm:ss"
            value-format="yyyy-MM-dd HH:mm:ss"
            :picker-options="pickerOptions"
        />
      </template>

      <el-input v-else-if="v.comparator == 'match'" v-model="v.ftv" clearable style="width: 300px" />

      <template v-else>

          <select-values
              ref="values"
              v-model="v.ftv"
              :table-typ="tableTyp"
              style="width: 300px"
              :data="v.columnName"
          />

      </template>
    </div>
  </div>
</template>

<script setup>

import { ref, reactive, watch, onMounted } from 'vue'
import { pickerOptions } from '@/utils/date'
import {
  dataTypeCalcuSymbol,
  inputSymbolArr,
  noValueSymbolArr,
  rangeSymbolArr,
  rangeTimeSymbolArr
} from '@/utils/base-data'

import SelectValues from '@/components/AnalyseTools/FilterWhere/values.vue'

const props = defineProps({
  value: {
    type: Object,
    default: () => ({})
  },
  datas: {
    type: Object,
    default: () => ({})
  },
  options: {
    type: Array,
    default: () => []
  },
  dataTypeMap: {
    type: Array,
    default: () => []
  },
  tableTyp: {
    type: Number,
    default: 0
  }
})

const emit = defineEmits(['update:value'])

const v = reactive(props.value)

const pickerOptionsRef = ref(pickerOptions)

watch(() => v.columnName, () => {
  emit('update:value', v)
})

watch(() => v.comparator, () => {
  emit('update:value', v)
})

watch(() => v.ftv, () => {
  emit('update:value', v)
})

onMounted(() => {
  pickerOptionsRef.value = pickerOptions
})

const getDataTypeCalcuSymbol = (data) => {
  let typ = 0
  if (props.dataTypeMap.hasOwnProperty(props.tableTyp.to$t())) {
    for (const v of props.dataTypeMap[props.tableTyp.to$t()]) {
      if (data === v.attribute_name) {
        typ = v.data_type
      }
    }
  }
  return dataTypeCalcuSymbol[typ]
}

const changeColumnNameSelect = () => {
  const tmp = getDataTypeCalcuSymbol(v.columnName)
  for (const i in tmp) {
    v.comparator = i
    break
  }
  emit('update:value', v)
}
</script>

<style scoped>

</style>
