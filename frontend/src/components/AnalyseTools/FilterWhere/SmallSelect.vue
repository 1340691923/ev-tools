<template>
  <el-select
      v-model="currValue"
      multiple
      allow-create
      placeholder="请选择"
      style="width: 150px"
      class="yt-select"
      filterable
      v-bind="$attrs"
      :filter-method="userFilter"
      :disabled="disabled"
      :clearable="clearable"
      @change="change"
  >
    <el-option
        v-for="option in renderOption"
        :key="option.value"
        :value="option.value"
        :label="option.label"
    >{{ option.label }}
    </el-option>
  </el-select>
</template>

<script setup>
import {sdk} from "@/plugin_sdk/sdk"

import { ref, watch, computed } from 'vue';

const props = defineProps({
  value: {
    type: [String, Number],
    default: ''
  },
  max: {
    type: Number,
    default: 300
  },
  disabled: {
    type: Boolean,
    default: false
  },
  clearable: {
    type: Boolean,
    default: true
  },
  options: {
    type: Array,
    default: () => []
  }
});

const emit = defineEmits(['update:value', 'change']);

const renderOption = ref([]);

const currValue = computed({
  get() {
    return props.value || '';
  },
  set(value) {
    emit('update:value', value);
  }
});

const addValueOptions = () => {
  if (currValue.value) {
    const target = props.options.find(item => item.value === currValue.value);
    if (target && renderOption.value.every(item => item.value !== target.value)) {
      renderOption.value.unshift(target);
    }
  }
};

const addFilterOptions = (label) => {
  const target = props.options.find(item => item.label === label);
  if (target && renderOption.value.every(item => item.label !== target.label)) {
    renderOption.value.unshift(target);
  }
};

const userFilter = (query = '') => {
  let arr = [];
  if (query !== '') {
    arr = props.options.filter(item => item.label.includes(query) || item.value.includes(query));
  } else {
    arr = JSON.parse(JSON.stringify(props.options));
  }

  if (arr.length > props.max) {
    renderOption.value = arr.slice(0, props.max);
    addFilterOptions(query);
  } else {
    renderOption.value = arr;
  }
};

const change = (value) => {
  emit('change', value);
  if (!value) {
    userFilter();
  }
};

watch(
    [() => props.value, () => props.options],
    () => {
      addValueOptions();
    },
    { deep: true }
);

// Initialize options
userFilter();
addValueOptions();
</script>
