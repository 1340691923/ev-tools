<template>
  <div class="app-container">

    <div class="search-container">
      <el-form :inline="true">
        <el-form-item label="存储库">
          <el-select
              style="width:10rem"
              v-model="snapshotNameList"
              clearable
              filterable
              multiple
              :placeholder="$t('请选择存储库')"
              v-loading="loading"
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
          >{{ $t('新建存储库') }}
          </el-button>
        </el-form-item>
        <el-form-item>
          <el-button

              type="success"
              @click.native="search"
          >{{ $t('刷新') }}
          </el-button>
        </el-form-item>
      </el-form>


    </div>
    <el-card shadow="never" class="table-container">
      <el-table
          :data="tableData"
      >
        <el-table-column
            :label="$t('序号')"
            align="center"
            fixed
            width="100"
        >
          <template #default="scope">
            {{ scope.$index + 1 }}
          </template>
        </el-table-column>
        <el-table-column align="center" :label="$t('存储库名')" prop="name" width="200"/>
        <el-table-column align="center" :label="$t('存储库地址')" prop="location" width="300"/>
        <el-table-column align="center" :label="$t('类型')" prop="type" width="200"/>

        <el-table-column align="center" :label="$t('是否压缩')" prop="compress" width="100"/>
        <el-table-column align="center" :label="$t('分解块大小')" prop="chunk_size" width="100"/>
        <el-table-column align="center" :label="$t('是否只读(默认false)')" prop="readonly" width="100"/>

        <el-table-column align="center" :label="$t('制作快照的速度')" width="100">
          <template #default="scope">
            <div v-if="scope.row.max_snapshot_bytes_per_sec != ''">{{ scope.row.max_snapshot_bytes_per_sec }}/s</div>
            <div v-else>20mb/s</div>
          </template>
        </el-table-column>
        <el-table-column align="center" :label="$t('快照恢复的速度')" width="100">
          <template #default="scope">
            <div v-if="scope.row.max_restore_bytes_per_sec != ''">{{ scope.row.max_restore_bytes_per_sec }}/s</div>
            <div v-else>20mb/s</div>
          </template>
        </el-table-column>

        <el-table-column align="center" :label="$t('操作')" fixed="right" width="350">
          <template #default="scope">
            <el-button-group>
              <el-button

                  type="success"


                  @click="look(scope.row.name)"
              >
                {{ $t('查看') }}
              </el-button>
              <el-button

                  type="warning"

                  @click="updateSnapshotData(scope.row)"
              >
                {{ $t('修改') }}
              </el-button>
              <el-button

                  type="danger"


                  @click="SnapshotDeleteRepositoryAction(scope.row.name)"
              >{{ $t('删除') }}
              </el-button>
              <el-button

                  type="warning"


                  @click="CleanupeRepository(scope.row.name)"
              >{{ $t('清理') }}
              </el-button>

            </el-button-group>

          </template>
        </el-table-column>
      </el-table>
    </el-card>

    <el-drawer

        ref="drawer"
        :title="$t('JSON数据')"
        :before-close="drawerHandleClose"
        v-model="drawerShow"

        direction="rtl"
        close-on-press-escape
        destroy-on-close
        size="80%"
    >
      <json-editor
          v-if="drawerShow"
          v-model:value="jsonData"
          styles="width: 100%"
          :read="true"
          :title="$t('JSON数据')"
      />

    </el-drawer>

        <add
            v-if="openAddDialog"
            :snapshot-data="snapshotData"
            :open="openAddDialog"
            @close="closeAddDialog"
        />
  </div>
</template>

<script lang="ts">
import {sdk} from "@elasticview/plugin-sdk"
import Add from '@/views/back-up/components/addRepository.vue'
import {Component, Vue, toNative} from 'vue-facing-decorator'
import {getCurrentInstance} from "vue";
import {
  CleanupeRepositoryAction,
  SnapshotDeleteRepositoryAction,
  SnapshotRepositoryListAction
} from "../../api/es-backup";
import {ElMessage, ElNotification} from "element-plus";
import JsonEditor from '@/components/JsonEditor/index.vue'
@Component({
  name: 'Task',
  components: {
    Add,
    JsonEditor
  },
  setup() {
    const ctx = getCurrentInstance().appContext.config.globalProperties
    return {ctx}
  },
})
class Task extends Vue {
  ctx = null
  snapshotData = null
  openAddDialog = false
  loading = false
  snapshotNameList = []
  resData = {}
  jsonData = "{}"
  name = ''
  drawerShow = false
  tableData = []
  index = 0
  tableHeader = []

  mounted() {
    this.initSnapshotName()
    this.search()
  }

  async CleanupeRepository(name) {
    const input = {}
    input['es_connect'] = sdk.GetSelectEsConnID()
    input['name'] = name
    const {data, code, msg} = await CleanupeRepositoryAction(input)
    if (code != 0) {
      ElMessage.error({
        type: 'error',
        message: msg
      })
      return
    } else {
      this.initSnapshotName()
      this.search()
      ElMessage.success({
        type: 'success',
        message: msg
      })
      return
    }
  }

  async SnapshotDeleteRepositoryAction(name) {
    const input = {}
    input['es_connect'] = sdk.GetSelectEsConnID()
    input['name'] = name
    const {data, code, msg} = await SnapshotDeleteRepositoryAction(input)
    if (code != 0) {
      ElMessage.error({
        type: 'error',
        message: msg
      })
      return
    } else {
      this.initSnapshotName()
      this.search()
      ElMessage.success({
        type: 'success',
        message: msg
      })
      return
    }
  }

  updateSnapshotData(row) {
    this.snapshotData = row
    this.openAddDialog = true
  }

  closeAddDialog(addSuccess) {
    this.openAddDialog = false
    this.snapshotData = null
    if (addSuccess) {
      this.initSnapshotName()
      this.search()
    }
  }

  async initSnapshotName() {
    this.loading = true
    const input = {}
    input['es_connect'] = sdk.GetSelectEsConnID()
    const {data, code, msg} = await SnapshotRepositoryListAction(input)
    if (code == 0) {
      this.resData = data.res
      ElNotification({
        title: '成功',
        dangerouslyUseHTMLString: true,
        message: `
<strong> <b style="color: orange">您设置的path.repo为${data.pathRepo.join(',')}</b><br></strong>
          `,
        type: 'success'
      })
    } else if (code == 199999) {
      ElNotification({
        title: 'Error',
        dangerouslyUseHTMLString: true,
        message: `
<strong>
            <i style="color: orange">path.repo没有设置</i><br>
<i>在elasticsearch.yml 配置文件中配置仓库base目录</i><br>
<i>添加path.repo: /tmp/tmp (/tmp/tmp 为快照备份所在文件夹, <br><i style="color: orange">注意</i>首先要先创建这个文件夹)</i>

          `,
        type: 'error'
      })

    } else {
      ElMessage.error({
        type: 'error',
        message: msg
      })

    }
    this.loading = false
  }

  async search() {
    const input = {}
    input['es_connect'] = sdk.GetSelectEsConnID()
    input['snapshot_info_list'] = this.snapshotNameList
    input['repository'] = 'test3'
    const {data, code, msg} = await SnapshotRepositoryListAction(input)

    if (code != 0) {
      if (code != 199999) {
        ElMessage.error({
          type: 'error',
          message: msg
        })
        return
      }



    }else if (code === 199999) {
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
      this.tableData = data.list
    }
  }

  openDetail(row, index, tmp) {
    this.name = row.name
    this.drawerShow = true
  }

  look(index) {
    this.name = index
    this.jsonData =JSON.stringify( this.resData[index] ,null, '\t')
    this.drawerShow = true
  }

  drawerHandleClose(done) {
    done()
  }
}

export default toNative(Task)
</script>

<style scoped>
/* 默认两列布局 */
.responsive-form .el-row {
  margin-bottom: 20px;
}

/* 移动端一列布局 */
@media (max-width: 767px) {
  .responsive-form .el-col {
    flex: 0 0 100%;
    max-width: 100%;
  }
}
</style>
