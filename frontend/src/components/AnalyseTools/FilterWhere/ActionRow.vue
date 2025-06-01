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
              ref="selectValuesRef"
              v-model="v.ftv"
              :table-typ="tableTyp"
              style="width: 300px"
              :data="v.columnName"
          />
          <el-button
              class="import-btn"
              type="primary"
              size="small"
              circle
              @click="openBatchImportDialog"
          >
            <el-icon><Plus /></el-icon>
          </el-button>
      </template>
    </div>

    <!-- 批量导入对话框 -->
    <el-dialog
        v-model="batchImportDialogVisible"
        title="批量导入"
        width="80%"
    >
      <el-form>
        <el-form-item label="">
          <el-input
              v-model="batchImportContent"
              type="textarea"
              :rows="10"
              placeholder="请输入内容，每行一个值"
          />
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="batchImportDialogVisible = false">取消</el-button>
          <el-button type="primary" @click="handleBatchImport">完成</el-button>
        </span>
      </template>
    </el-dialog>
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
import { Plus } from '@element-plus/icons-vue'
import SelectValues from '@/components/AnalyseTools/FilterWhere/values.vue'
import { ElMessage } from 'element-plus'

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
const selectValuesRef = ref(null)
const batchImportDialogVisible = ref(false)
const batchImportContent = ref('')

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
  if (props.dataTypeMap.hasOwnProperty(props.tableTyp.toString())) {
    for (const v of props.dataTypeMap[props.tableTyp.toString()]) {
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

// 打开批量导入对话框
const openBatchImportDialog = () => {
  batchImportDialogVisible.value = true
  batchImportContent.value = ''
}

// 处理批量导入
const handleBatchImport = () => {
  if (!batchImportContent.value.trim()) {
    ElMessage.warning('导入内容不能为空')
    return
  }

  // 按行分割输入内容
  const lines = batchImportContent.value.split('\n').filter(line => line.trim() !== '')

  if (lines.length === 0) {
    ElMessage.warning('没有有效的导入内容')
    return
  }

  // 如果select-values组件有批量导入方法，则调用
  if (selectValuesRef.value && typeof selectValuesRef.value.batchImportValues === 'function') {
    selectValuesRef.value.batchImportValues(lines)

    // 确保更新值到父组件
    if (Array.isArray(v.ftv)) {
      // 如果已经是数组，则添加新值
      const currentValues = new Set(v.ftv);
      lines.forEach(line => {
        const trimmed = line.trim();
        if (trimmed) currentValues.add(trimmed);
      });
      v.ftv = Array.from(currentValues);
    } else {
      // 如果不是数组，则创建新数组
      v.ftv = lines.filter(line => line.trim());
    }

    // 发送更新事件
    emit('update:value', v);

    ElMessage.success(`成功导入${lines.length}个值`)
  } else {
    // 如果组件没有提供批量导入方法，则将值设置到v.ftv
    if (!Array.isArray(v.ftv)) {
      v.ftv = [];
    }
    lines.forEach(line => {
      const trimmed = line.trim();
      if (trimmed && !v.ftv.includes(trimmed)) {
        v.ftv.push(trimmed);
      }
    });
    emit('update:value', v)
    ElMessage.success(`成功导入${lines.length}个值`)
  }

  // 关闭对话框
  batchImportDialogVisible.value = false
}
</script>

<style scoped>
.ta-filter-condition {
  display: flex;
  align-items: center;
}

.import-btn {
  margin-left: 5px;
  height: 24px;
  width: 24px;
  padding: 0;
  font-size: 12px;
}
</style>
