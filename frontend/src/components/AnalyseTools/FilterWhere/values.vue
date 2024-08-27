<template>
  <el-select
      collapse-tags-tooltip
      v-model="selectVal"
      filterable
      :options="options"
      placeholder="请输入筛选值"
      style="width: 240px"
      multiple
      default-first-option
      allow-create
      collapse-tags
      :reserve-keyword="false"
      clearable
  />
</template>

<script>
import { ref, watch, onMounted } from 'vue';
import {sdk} from "@/plugin_sdk/sdk"

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
    const selectVal = ref(props.value);
    const cleanValues = () => {
      selectVal.value = [];
      emit('update:value', selectVal.value);
    };

    const changeValue = () => {
      emit('update:value', selectVal.value);
    };

    const initValue = (data) => {
      options.value = []; // Initialize options based on your logic
    };

    watch(() => props.data, (newV) => {
      initValue(newV);
    });

    watch(selectVal, (newV) => {
      emit('update:value', newV);
    });

    onMounted(() => {
      initValue(props.data);
    });

    return {
      options,
      selectVal,
      cleanValues,
      changeValue
    };
  }
};
</script>

<style scoped>
</style>
