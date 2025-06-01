<template>
  <div>
    <el-form>
      <el-form-item>
        <el-button type="success"  @click="addAlias">
          {{ $t('新增别名') }}
        </el-button>
        <el-button type="primary" @click="batchAdd">
          {{ $t('批量提交') }}
        </el-button>
        <el-button type="info" @click="getAlias">
          {{ $t('重置') }}
        </el-button>
        <IndexSelect style="width: 210px;margin-left:5px" :multiple="true" :clearable="true" :placeholder="$t('迁移别名到多个索引上')" @change="changeAnotherIndex" />
      </el-form-item>
      <el-form-item
          v-for="(alias, index) in aliasList"
          :key="index"
          :label="$t('别名') + Number(index + 1)"
      >
        <el-input v-model="alias.name" :readonly="alias.types != 'new'" style="width:300px" />

        <el-button v-show="anotherIndex.length > 0"  type="success" @click="moveAliasToIndex(index)">
          迁移
        </el-button>

        <el-button v-clipboard:copy="alias.name" v-clipboard:success="onCopy" v-clipboard:error="onError" type="success">
          {{ $t('复制') }}
        </el-button>
        <el-button v-show="alias.types == 'new'" type="primary"  @click="submitForm(index)">
          {{ $t('提交') }}
        </el-button>
        <el-button  type="danger" @click="removeAlias(index)">
          {{ $t('删除') }}
        </el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script setup>
import {ref, onMounted, onActivated} from 'vue';
import { AddAliasToIndex, BatchAddAliasToIndex, GetAliasAction, MoveAliasToIndex, RemoveAlias } from '@/api/es-index';
import IndexSelect from '@/components/index/select.vue';
import {ElMessage} from "element-plus";
import {sdk} from "@elasticview/plugin-sdk";


const props = defineProps({
  indexName: {
    type: String,
    default: ''
  }
});

const aliasList = ref([]);
const anotherIndex = ref([]);

const addAlias = () => {
  aliasList.value.push({ name: '', types: 'new' });
};

const getAlias = async () => {
  aliasList.value = [];
  const input = {
    es_connect: sdk.GetSelectEsConnID(),
    index_name: props.indexName
  };
  const { data, code, msg } = await GetAliasAction( input);
  console.log(data, code, msg);
  if (code === 0) {
    for (const k in data) {
      aliasList.value.push({ name: data[k].AliasName, types: 'old' });
    }
  } else {
    ElMessage({
      message: msg,
      type: 'error'
    });
  }
};

const submitForm = async (index) => {
  const alias = aliasList.value[index].name;
  const input = {
    types: 1,
    index_name: props.indexName,
    alias_name: alias.trim(),
    es_connect: sdk.GetSelectEsConnID()
  };
  const { data, code, msg } = await AddAliasToIndex( input);
  if (code === 0) {
    ElMessage({
      type: 'success',
      message: msg
    });
  } else {
    ElMessage({
      type: 'error',
      message: msg
    });
  }
};

const batchAdd = async () => {
  const aliasNames = aliasList.value.filter(alias => alias.types === 'new').map(alias => alias.name);
  const input = {
    types: 4,
    index_name: props.indexName,
    new_alias_name_list: aliasNames,
    es_connect: sdk.GetSelectEsConnID()
  };
  const { data, code, msg } = await BatchAddAliasToIndex( input);
  if (code === 0) {
    ElMessage({
      type: 'success',
      message: msg
    });
  } else {
    ElMessage({
      type: 'error',
      message: msg
    });
  }
};

const removeAlias = async (index) => {
  const alias = aliasList.value[index];
  if (alias.types === 'new') {
    aliasList.value.splice(index, 1);
  } else if (alias.types === 'old') {
    const input = {
      types: 2,
      index_name: props.indexName,
      alias_name: alias.name.trim(),
      es_connect: sdk.GetSelectEsConnID()
    };
    const { data, code, msg } = await RemoveAlias( input);
    if (code === 0) {
      ElMessage({
        type: 'success',
        message: msg
      });
      aliasList.value.splice(index, 1);
    } else {
      ElMessage({
        type: 'error',
        message: msg
      });
    }
  }
};

const changeAnotherIndex = (v) => {
  anotherIndex.value = v;
};

const moveAliasToIndex = async (index) => {
  const alias = aliasList.value[index].name;
  const input = {
    types: 3,
    new_index_list: anotherIndex.value,
    alias_name: alias.trim(),
    es_connect: sdk.GetSelectEsConnID()
  };
  const { data, code, msg } = await MoveAliasToIndex( input);
  if (code === 0) {
    ElMessage({
      type: 'success',
      message: msg
    });
  } else {
    ElMessage({
      type: 'error',
      message: msg
    });
  }
};

const onCopy = (e) => {
  ElMessage({
    message: '复制成功！',
    type: 'success'
  });
};

const onError = (e) => {
  ElMessage({
    message: '复制失败！',
    type: 'error'
  });
};

onMounted(() => {
  console.log('alias mounted')
  getAlias();
});

onActivated(() => {
  console.log('alias activated')
});
</script>

<style scoped>
</style>
