<template>
  <el-select
      collapse-tags-tooltip
      v-model="selectVal"
      filterable
      placeholder="请输入筛选值"
      style="width: 240px"
      multiple
      default-first-option
      allow-create
      collapse-tags
      :reserve-keyword="false"
      clearable
  >
    <el-option
        v-for="item in selectOptions"
        :key="item"
        :label="item"
        :value="item"
    />
  </el-select>
</template>

<script>
import { ref, watch, onMounted, computed } from 'vue';
import {sdk} from "@elasticview/plugin-sdk"

export default {
  name: 'Values',

  props: {
    data: {
      type: Array,
      default: () => []
    },
    value: {
      type: [String, Number, Array],
      default: ''
    },
  },
  setup(props, { emit }) {
    const options = ref([]);
    const selectVal = ref(Array.isArray(props.value) ? [...props.value] : []);

    // 生成选项列表，确保包含当前选中的值
    const selectOptions = computed(() => {
      // 合并选项和当前选中值，去重
      const allOptions = [...new Set([...options.value, ...(Array.isArray(selectVal.value) ? selectVal.value : [])])];
      return allOptions;
    });

    const cleanValues = () => {
      selectVal.value = [];
      emit('update:value', selectVal.value);
    };

    const changeValue = () => {
      emit('update:value', selectVal.value);
    };

    const initValue = (data) => {
      // 保留原有选项
      // options.value = []; // 不清空选项
    };

    // 批量导入值的方法
    const batchImportValues = (values) => {
      if (!values || !Array.isArray(values)) return;

      // 如果selectVal不是数组，初始化为空数组
      if (!Array.isArray(selectVal.value)) {
        selectVal.value = [];
      }

      // 将所有新值添加到选中值中，避免重复添加
      const newValues = [];
      values.forEach(val => {
        const trimmedVal = val.trim();
        if (trimmedVal && !selectVal.value.includes(trimmedVal)) {
          newValues.push(trimmedVal);
        }
      });

      // 更新选中值
      selectVal.value = [...selectVal.value, ...newValues];

      // 更新options以显示这些值
      options.value = [...new Set([...options.value, ...newValues])];

      // 发出值变更事件
      emit('update:value', selectVal.value);
    };

    watch(() => props.data, (newV) => {
      initValue(newV);
    });

    watch(selectVal, (newV) => {
      emit('update:value', newV);
    });

    onMounted(() => {
      initValue(props.data);
      // 确保值是数组
      if (!Array.isArray(selectVal.value)) {
        selectVal.value = [];
      }
    });

    return {
      options,
      selectVal,
      selectOptions,
      cleanValues,
      changeValue,
      batchImportValues
    };
  }
};
</script>

<style scoped>
</style>
