<template>
  <div>
    <el-drawer size="80%" v-model="open" :title="$t('创建快照')" @close="closeDialog">
      <el-card class="box-card">
        <el-form label-width="500px" label-position="left">
          <el-form-item :label="$t('存储库')">
            <el-select
              v-model="form.repositoryName"
              clearable
              filterable
              :placeholder="$t('请选择存储库')"
            >
              <el-option
                v-for="(v,k,index) of repositoryNameList"
                :key="index"
                :label="k"
                :value="k"
              />
            </el-select>
          </el-form-item>
          <el-form-item :label="$t('快照名')">
            <el-input v-model="form.snapshotName" :placeholder="$t('快照名')" />
          </el-form-item>
          <el-form-item :label="$t('需要备份的索引')">
            <index-select
              :multiple="true"
              :have-all="true"
              :clearable="true"
              style="width: 210px"
              :placeholder="$t('需要备份的索引')"
              @change="changeIndex"
            />
          </el-form-item>

          <!--          <el-form-item
            :label="$t('ignore_unavailable   【把这个选项设置为 true 的时候在创建快照的过程中会忽略不存在的索引,如果没有设置ignore_unavailable，在索引不存在的情况下快照请求将会失败。】')"
          >
            <el-select v-model="form.ignore_unavailable" filterable>
              <el-option :label="$t('不设置')" :value="null" />
              <el-option :label="$t('是')" :value="true" />
              <el-option :label="$t('否')" :value="false" />
            </el-select>
          </el-form-item>-->
          <el-form-item
            :label="$t('include_global_state 【通过设置 include_global_state 为false 能够防止 集群的全局状态被作为快照的一部分存储起来】')"
          >
            <el-select v-model="form.include_global_state" filterable>
              <el-option :label="$t('不设置')" :value="null" />
              <el-option :label="$t('是')" :value="true" />
              <el-option :label="$t('否')" :value="false" />
            </el-select>
          </el-form-item>
          <el-form-item :label="$t('partial  【如果快照中的1个或多个索引不是全部主分片都可用，就会导致整个创建快照的过程失败。 通过设置 partial 为 true 可以改变这个行为】')">
            <el-select v-model="form.partial" clearable filterable>
              <el-option :label="$t('不设置')" :value="null" />
              <el-option :label="$t('是')" :value="true" />
              <el-option :label="$t('否')" :value="false" />
            </el-select>
          </el-form-item>
          <el-form-item :label="$t('创建方式')">
            <el-select v-model="form.wait" clearable filterable>
              <el-option :label="$t('不设置')" :value="null" />
              <el-option :label="$t('异步创建')" :value="true" />
              <el-option :label="$t('同步创建')" :value="false" />
            </el-select>
          </el-form-item>
        </el-form>
        <template #footer>
          <el-button

            type="danger"

            @click="closeDialog"
          >{{ $t('取消') }}
          </el-button>
          <el-button

            type="primary"

            @click="confirm"
          >{{ $t('确认') }}
          </el-button>
        </template>
      </el-card>
    </el-drawer>
  </div>
</template>

<script setup>
import {ref, onMounted, getCurrentInstance} from 'vue'
import { CreateSnapshotAction, SnapshotRepositoryListAction } from '@/api/es-backup'
import { ElMessage, ElNotification } from 'element-plus'
const ctx = getCurrentInstance().appContext.config.globalProperties
import {sdk} from "@/plugin_sdk/sdk"
import IndexSelect from "@/components/index/select.vue";
// Props
const props = defineProps({
  open: {
    type: Boolean,
    default: false
  },
  snapshotData: {
    type: Object,
    default: null
  }
})

const open = ref(props.open)

// Local state
const isOpen = ref(false)
const form = ref({
  snapshotName: '',
  repositoryName: '',
  indexList: [],
  ignore_unavailable: null,
  include_global_state: null,
  partial: null,
  wait: null
})
const repositoryNameList = ref({})
const loading = ref(false)

// Methods
const changeIndex = (index) => {
  console.log(index)
  form.value.indexList = index
}

const initRepositoryName = async () => {
  loading.value = true
  const input = { es_connect: sdk.GetSelectEsConnID() }
  const { data, code, msg } = await SnapshotRepositoryListAction(input)

  if (code === 0) {
    repositoryNameList.value = data.res
  } else if (code === 199999) {
    ElNotification({
      title: 'Error',
      dangerouslyUseHTMLString: true,
      message: `
        <strong>
          <i style="color: orange">path.repo没有设置</i><br>
          <i>在elasticsearch.yml 配置文件中配置仓库base目录</i><br>
          <i>添加path.repo: /tmp/tmp (/tmp/tmp 为快照备份所在文件夹, <br><i style="color: orange">注意</i>首先要先创建这个文件夹)</i>
        </strong>
      `,
      type: 'error'
    })
  } else {
    ElMessage({
      type: 'error',
      message: msg
    })
  }

  loading.value = false
}

const closeDialog = () => {
  emit('close', false)
}

const confirm = async () => {
  const input = { ...form.value, es_connect: sdk.GetSelectEsConnID() }
  const { code, msg } = await CreateSnapshotAction(input)

  if (code === 0) {
    emit('close', true)
    ElMessage({
      type: 'success',
      message: msg
    })
  } else {
    ElMessage({
      type: 'error',
      message: msg
    })
  }
}

const emit = defineEmits(['close'])

// Lifecycle hooks
onMounted(() => {
  initRepositoryName()
})
</script>


<style scoped>

</style>
