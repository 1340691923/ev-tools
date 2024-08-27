<template>
  <div>
    <el-config-provider :size="size">
      <router-view />
    </el-config-provider>
  </div>
</template>

<script setup>

import {sdk} from '@/plugin_sdk/sdk'
import {computed} from "vue";

const size = computed(()=>{
  return sdk.getLayoutSize()
})


</script>
