<template>
  <div class="app-container">

    <div class="search-container">
      <el-form :inline="true">
        <el-form-item>
          <el-select
              v-model="repositoryName"
              clearable
              filterable
              style="width: 200px"
              :placeholder="$t('请选择存储库')"
              :loading="loading"
              @change="search"
          >
            <el-option
                v-for="(v,k,index) of resData"
                :key="index"
                :label="k"
                :value="k"
            />
          </el-select>

        </el-form-item>
        <el-form-item>
          <el-button

              type="warning"
              @click.native="openAddDialog = true"
          >{{ $t('新建快照') }}
          </el-button>
          <el-button

              type="success"
              @click="search"
          >{{ $t('刷新') }}
          </el-button>
        </el-form-item>
      </el-form>
    </div>


    <el-table
      :data="tableData"
    >
      <el-table-column
        :label="$t('序号')"
        align="center"
        fixed
        width="50"
      >
        <template slot-scope="scope">
          {{ scope.$index + 1 }}
        </template>
      </el-table-column>

      <el-table-column align="center" :label="$t('快照名')" prop="id" width="100" />
      <el-table-column align="center" :label="$t('备份索引数')" prop="indices" width="100" />

      <el-table-column align="center" :label="$t('状态')" width="100">
        <template slot-scope="scope">
          <el-tag v-if="scope.row.status == 'SUCCESS'" type="success">{{ $t('成功') }}</el-tag>
          <el-tag v-else-if="scope.row.status == 'IN_PROGRESS'" type="warning">{{ $t('还在进行中') }}</el-tag>
          <el-tag v-else>{{ scope.row.status }}</el-tag>
        </template>
      </el-table-column>

      <el-table-column align="center" :label="$t('开始详细时间')" width="180">
        <template slot-scope="scope">
          <div>{{ timestampToTime(scope.row.start_epoch) }}</div>
        </template>
      </el-table-column>

      <el-table-column align="center" :label="$t('结束详细时间')" width="180">
        <template slot-scope="scope">
          <div>{{ timestampToTime(scope.row.end_epoch) }}</div>
        </template>
      </el-table-column>

      <el-table-column align="center" :label="$t('耗费时长')" prop="duration" width="120" />
      <el-table-column align="center" :label="$t('分片总数')" prop="total_shards" width="120" />
      <el-table-column align="center" :label="$t('成功分片数')" prop="successful_shards" width="120" />
      <el-table-column align="center" :label="$t('失败分片数')" prop="failed_shards" width="120" />

      <el-table-column align="center" :label="$t('操作')" fixed="right" width="350">
        <template slot-scope="scope">
          <el-button-group>
            <el-button
              type="success"


              @click="look(scope.row.id)"
            >{{ $t('查看') }}
            </el-button>

            <el-button

              type="danger"

              @click="SnapshotDeleteAction(scope.row.id)"
            >{{ $t('删除') }}
            </el-button>
            <el-button

              type="warning"

              @click="openRestore(scope.row.id)"
            >{{ $t('恢复') }}
            </el-button>
            <el-button


              type="primary"


              @click="status(scope.row.id)"
            >{{ $t('状态') }}
            </el-button>

          </el-button-group>

        </template>
      </el-table-column>
    </el-table>

    <el-drawer
      ref="drawer"
      :title="$t('快照详细信息')"
      :before-close="drawerHandleClose"
      v-model="drawerShow"
      direction="rtl"
      close-on-press-escape
      destroy-on-close
      size="50%"
    >

      <json-editor

        v-model:value="snapshotDetailJSON"
        class="res-body"
        styles="width: 100%"
        :read="true"
        :title="$t('快照详细信息')"
      />
    </el-drawer>

    <add
      v-if="openAddDialog"
      :open="openAddDialog"
      @close="closeAddDialog"
    />
    <snapshot-restore
      v-if="openRestoreDialog"
      :repository="repositoryName"
      :snapshot="name"
      :open="openRestoreDialog"
      @close="closeoRestoreDialog"
    />

  </div>
</template>

<script setup>
import {ref, onMounted, computed, getCurrentInstance} from 'vue'
import {
  SnapshotDeleteAction,
  SnapshotDetailAction,
  SnapshotListAction,
  SnapshotRepositoryListAction,
  SnapshotStatusAction
} from '@/api/es-backup'
import { ElMessage, ElNotification } from 'element-plus'
import { timestampToTime } from '@/utils/time'

import Add from '@/views/back-up/components/addSnapshot.vue'
import SnapshotRestore from '@/views/back-up/components/snapshotRestore.vue'
import {sdk} from "@/plugin_sdk/sdk"

const ctx = getCurrentInstance().appContext.config.globalProperties

// State
const openRestoreDialog = ref(false)
const openAddDialog = ref(false)
const loading = ref(false)
const repositoryName = ref('')
const resData = ref({})
const name = ref('')
const drawerShow = ref(false)
const tableData = ref([])
const snapshotDetail = ref('')
const evUtils = {} // Assume this is defined globally or imported from a utility module

// Computed properties
const snapshotDetailJSON = computed(() => JSON.stringify(snapshotDetail.value, null, '\t'))

// Methods
const status = async (snapshotName) => {
  const input = {
    es_connect: sdk.GetSelectEsConnID(),
    snapshot: snapshotName,
    repository: repositoryName.value
  }

  const { data, code, msg } = await SnapshotStatusAction(input)
  if (code !== 0) {
    ElMessage.error(msg)
    return
  }
  snapshotDetail.value = data.snapshots
  drawerShow.value = true
  ElMessage.success(msg)
}

const openRestore = (snapshotName) => {
  name.value = snapshotName
  openRestoreDialog.value = true
}

const deleteSnapshot = async (snapshotName) => {
  const input = {
    es_connect: sdk.GetSelectEsConnID(),
    snapshot: snapshotName,
    repository: repositoryName.value
  }

  const { code, msg } = await SnapshotDeleteAction(input)
  if (code !== 0) {
    ElMessage.error(msg)
    return
  }
  initRepositoryName()
  search()
  ElMessage.success(msg)
}

const closeRestoreDialog = (addSuccess) => {
  openRestoreDialog.value = false
  if (addSuccess) {
    initRepositoryName()
    search()
  }
}

const closeAddDialog = (addSuccess) => {
  openAddDialog.value = false
  if (addSuccess) {
    initRepositoryName()
    search()
  }
}

const initRepositoryName = async () => {
  loading.value = true
  const input = { es_connect: sdk.GetSelectEsConnID() }
  const { data, code, msg } = await SnapshotRepositoryListAction(input)

  if (code === 0) {
    resData.value = data.res
    if (Object.keys(resData.value).length > 0) {
      repositoryName.value = Object.keys(resData.value)[0]
    }
    search()
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
    loading.value = false
    return
  } else {
    ElMessage.error(msg)
  }
  loading.value = false
}

const search = async () => {
  const input = {
    es_connect: sdk.GetSelectEsConnID(),
    repository: repositoryName.value
  }
  const { data, code, msg } = await SnapshotListAction(input)

  if (code !== 0) {
    ElMessage.error(msg)
    return
  }
  ElMessage.success(msg)
  tableData.value = data
}

const openDetail = (row) => {
  name.value = row.name
  drawerShow.value = true
}

const look = async (snapshotIndex) => {
  name.value = snapshotIndex
  const input = {
    es_connect: sdk.GetSelectEsConnID(),
    repository: repositoryName.value,
    snapshot: snapshotIndex
  }

  const { data, code, msg } = await SnapshotDetailAction(input)
  if (code !== 0) {
    ElMessage.error(msg)
    return
  }
  snapshotDetail.value = data.snapshots
  drawerShow.value = true
}

const drawerHandleClose = (done) => {
  done()
}

// Lifecycle hook
onMounted(() => {
  initRepositoryName()
})
</script>




<style scoped>

</style>
