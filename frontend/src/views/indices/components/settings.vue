<template>
  <div>
    <el-drawer size="90%"  v-model="open" :title="title" @close="closeDialog">
      <el-card class="box-card">
        <el-form label-width="500px" label-position="left">
          <el-form-item :label="$t('索引名称：')">
            <el-input v-model="indexName" :placeholder="$t('多个索引名用,隔开')" :disabled="settingsType !== 'add'" />
          </el-form-item>
          <el-form-item label="number_of_shards (分片数)：">
            <el-input v-model="form.number_of_shards" type="number" style="width: 300px" :disabled="isOpen" />
          </el-form-item>
          <el-form-item label="number_of_replicas (副本数)：">
            <el-input v-model="form.number_of_replicas" type="number" style="width: 300px" />
          </el-form-item>
          <el-form-item label="refresh_interval (索引的刷新时间间隔)：">
            <el-input v-model="form.refresh_interval" :placeholder="$t('索引的刷新时间间隔')" />
          </el-form-item>
          <el-form-item label="max_result_window  (from+size 的最大值)：">
            <el-input v-model="form['index.max_result_window']" type="number" style="width: 300px" />
          </el-form-item>
        </el-form>
        <template #footer>
          <el-button type="danger" @click="closeDialog">{{ $t('取消') }}</el-button>
          <el-button type="primary" @click="confirmSettings">{{ $t('确认') }}</el-button>
        </template>
      </el-card>
    </el-drawer>
  </div>
</template>

<script setup>
import {ref, computed, onMounted, getCurrentInstance} from 'vue';
import { ElMessage, ElLoading } from 'element-plus';
import { StatsAction, CreateAction, GetSettingsAction } from '@/api/es-index';
import { clone } from '@/utils/index';
const ctx = getCurrentInstance().appContext.config.globalProperties
import {sdk} from "@elasticview/plugin-sdk";
const props = defineProps({
  indexName: {
    type: String,
    default: '',
  },
  settingsType: {
    type: String,
    default: 'add',
  },
  open: {
    type: Boolean,
    default: false,
  },
});

const indexName = ref(props.indexName)
const settingsType = ref(props.settingsType)
const open = ref(props.open)

const emit = defineEmits(['finished', 'close']);

const staticConfig = ref([
  'number_of_shards',
  'index.translog.sync_interval',
  'index.shard.check_on_startup',
  'index.routing_partition_size',
  'index.codec',
]);

const isOpen = ref(false);

const form = ref({
  'index.shard.check_on_startup': '',
  'index.routing_partition_size': '',
  'index.codec': '',
  'index.auto_expand_replicas': '',
  'index.max_result_window': '',
  'index.max_rescore_window': '',
  'index.blocks.read_only': '',
  'index.blocks.read': '',
  'index.blocks.write': '',
  'index.blocks.metadata': '',
  'index.blocks.read_only_allow_delete': '',
  'index.max_refresh_listeners': '',
  'number_of_shards': '',
  'number_of_replicas': '',
  'refresh_interval': '',
  'index.translog.sync_interval': '',
  'index.translog.durability': '',
  'index.translog.flush_threshold_size': '',
  'index.merge.scheduler.max_thread_count': '',
  'index.merge.policy.max_merged_segment': '',
  'index.merge.policy.segments_per_tier': '',
});

const title = computed(() => {
  return settingsType.value === 'add' ? '新增索引配置' : '修改索引配置';
});

const confirmSettings = () => {
  if (form.value['index.routing_partition_size'] >= form.value['index.number_of_shards']) {
    ElMessage({
      type: 'error',
      message: '自定义路由值可以转发的目的分片数 必须小于 主分片数',
    });
    return;
  }

  if (indexName.value === '') {
    ElMessage({
      type: 'error',
      message: '索引名不能为空',
    });
    return;
  }

  const input = {};
  input['es_connect'] = sdk.GetSelectEsConnID();
  input['settings'] = clone(form.value);
  input['index_name'] = indexName.value;
  input['types'] = settingsType.value;

  if (settingsType.value === 'update') {
    for (const config of staticConfig.value) {
      delete input['settings'][config];
    }
  }

  for (const settingsKey in input['settings']) {
    const settingsVal = input['settings'][settingsKey];
    if (settingsVal === '') {
      delete input['settings'][settingsKey];
    }
  }

  const loading = ElLoading.service({
    lock: true,
    text: 'Loading',
    spinner: 'el-icon-loading',
    background: 'rgba(0, 0, 0, 0.7)',
  });

  CreateAction( input)
      .then((res) => {
        if (res.code === 0 || res.code === 200) {
          ElMessage({
            type: 'success',
            message: res.msg,
          });
          props.settingsType = 'update';
          emit('finished');
        } else {
          ElMessage({
            type: 'error',
            message: res.msg,
          });
        }
        loading.close();
      })
      .catch((err) => {
        loading.close();
      });
};

const catIndexSettings = () => {
  const input = {};
  input['es_connect'] = sdk.GetSelectEsConnID();
  input['index_name'] = indexName.value;

  GetSettingsAction( input)
      .then((res) => {
        if (res.code === 0 || res.code === 200) {
          const index = res.data['index'];
          form.value.number_of_replicas = index['number_of_replicas'];
          form.value.number_of_shards = index['number_of_shards'];
          if (index.hasOwnProperty('auto_expand_replicas')) {
            form.value['index.auto_expand_replicas'] = index['auto_expand_replicas'];
          }
          if (index.hasOwnProperty('blocks')) {
            const blocks = index['blocks'];
            if (blocks.hasOwnProperty('metadata')) {
              form.value['index.blocks.metadata'] = blocks['metadata'];
            }
            if (blocks.hasOwnProperty('read')) {
              form.value['index.blocks.read'] = blocks['read'];
            }
            if (blocks.hasOwnProperty('read_only')) {
              form.value['index.blocks.read_only'] = blocks['read_only'];
            }
            if (blocks.hasOwnProperty('read_only_allow_delete')) {
              form.value['index.blocks.read_only_allow_delete'] = blocks['read_only_allow_delete'];
            }
            if (blocks.hasOwnProperty('write')) {
              form.value['index.blocks.write'] = blocks['write'];
            }
          }
          if (index.hasOwnProperty('merge')) {
            const merge = index['merge'];
            if (merge.hasOwnProperty('policy')) {
              const policy = merge['policy'];
              if (policy.hasOwnProperty('max_merged_segment')) {
                form.value['index.merge.policy.max_merged_segment'] = policy['max_merged_segment'];
              }
              if (policy.hasOwnProperty('max_merged_segment')) {
                form.value['index.merge.policy.segments_per_tier'] = policy['segments_per_tier'];
              }
            }
            if (merge.hasOwnProperty('scheduler')) {
              const scheduler = merge['scheduler'];
              if (scheduler.hasOwnProperty('max_thread_count')) {
                form.value['index.merge.scheduler.max_thread_count'] = scheduler['max_thread_count'];
              }
            }
          }

          if (index.hasOwnProperty('max_refresh_listeners')) {
            form.value['index.max_refresh_listeners'] = index['max_refresh_listeners'];
          }
          if (index.hasOwnProperty('max_rescore_window')) {
            form.value['index.max_rescore_window'] = index['max_rescore_window'];
          }
          if (index.hasOwnProperty('max_result_window')) {
            form.value['index.max_result_window'] = index['max_result_window'];
          }

          if (index.hasOwnProperty('refresh_interval')) {
            form.value['refresh_interval'] = index['refresh_interval'];
          }
          if (index.hasOwnProperty('translog')) {
            const translog = index['translog'];
            if (translog.hasOwnProperty('durability')) {
              form.value['index.translog.durability'] = translog['durability'];
            }
            if (translog.hasOwnProperty('flush_threshold_size')) {
              form.value['index.translog.flush_threshold_size'] = translog['flush_threshold_size'];
            }
          }
          console.log(form.value, 'form');
        } else {
          ElMessage({
            type: 'error',
            message: res.msg,
          });
        }
      })
      .catch((err) => {
        console.log(err);
      });
};

const catIndexStatus = () => {
  const input = {}
  input['es_connect'] = sdk.GetSelectEsConnID()
  input['index_name'] = indexName.value

  StatsAction(input).then(res => {
    if (res.code == 0 || res.code == 200) {

      isOpen.value = (res.data[0].status == 'open')
    } else {
      ElMessage({
        type: 'error',
        message: res.msg
      })
    }
  }).catch(err => {
    console.error(err)
  })
}

const closeDialog = () => {
  emit('close')
}

onMounted(()=>{
  if (indexName.value != '') {
    catIndexStatus()
    catIndexSettings()
  }
})

</script>
