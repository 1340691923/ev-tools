<template>
  <div class="app-container" v-loading="connectLoading">
    <div class="search-container">
      <el-form>
        <el-form-item>
          <el-button
              type="primary"

              @click="searchData"
          >刷新
          </el-button>
        </el-form-item>
      </el-form>
    </div>




    <el-row :gutter="40">
      <el-col v-for="(v, k) in list" :key="k" :xs="40" :sm="11" :lg="11" class="card-panel-col">

        <el-card shadow="never" class="table-container" >


          <el-form  class="responsive-form">

            <el-row  :gutter="20">
              <el-col :span="12">
                <el-form-item label="节点名:" >
                  <el-tag>节点名:{{ v["name"] }}</el-tag>
                </el-form-item>
              </el-col>
              <el-col :span="12">
                <el-form-item label="">
                  <div>
                    <el-tooltip content="主节点" v-if="v.master === '*'" placement="top" style="cursor: pointer">
                      <el-tag type="primary"><el-icon><Star /></el-icon></el-tag>
                    </el-tooltip>
                    <el-tooltip content="仅选举节点" v-else-if="v['node.role'].includes('m') && v['node.role'].includes('v')" placement="top" style="cursor: pointer">
                      <el-tag type="primary"><el-icon><Tickets /></el-icon></el-tag>
                    </el-tooltip>
                    <el-tooltip content="主节点候选" v-else-if="v['node.role'].includes('m')" placement="top" style="cursor: pointer">
                      <el-tag type="primary"><el-icon><StarFilled /></el-icon></el-tag>
                    </el-tooltip>
                    <el-tooltip content="数据节点" v-if="v['node.role'].includes('d')" placement="top" style="cursor: pointer">
                      <el-tag type="success">
                        <el-icon><CreditCard /></el-icon>
                      </el-tag>
                    </el-tooltip>
                    <el-tooltip content="预处理节点" v-if="v['node.role'].includes('i')" placement="top" style="cursor: pointer">
                      <el-tag type="warning">
                        <el-icon><Postcard /></el-icon>
                      </el-tag>
                    </el-tooltip>
                    <el-tooltip content="仅协调节点" v-if="v['node.role'] === '-'" placement="top" style="cursor: pointer">
                      <el-tag type="warning">
                        <el-icon><Wallet /></el-icon>
                      </el-tag>
                    </el-tooltip>
                  </div>
                </el-form-item>
              </el-col>
            </el-row>

            <el-row  :gutter="20">
              <el-col :span="12">
                <el-form-item label="IP:" >
                  <el-tag>{{ v.ip }}</el-tag>
                </el-form-item>
              </el-col>
              <el-col :span="12">
                <el-form-item label="主节点:">
                  <div>
                    <template v-if="v.master">
                      <el-tag>yes</el-tag>
                    </template>
                    <template v-else-if="v.masterEligible">
                      <el-tag>eligible</el-tag>
                    </template>
                    <template v-else>
                      <el-tag>no</el-tag>
                    </template>
                  </div>
                </el-form-item>
              </el-col>
            </el-row>

            <el-row :gutter="20">
              <el-col :span="12">
                <el-form-item label="节点角色:" >
                  <el-tag>{{ v["node.role"] }}</el-tag>
                </el-form-item>
              </el-col>
              <el-col :span="12">
                <el-form-item label="负载:">
                  <el-tag>1m:{{ v.load_1m }} / 5m:{{ v.load_5m }} / 15m:{{ v.load_15m }}</el-tag>
                </el-form-item>
              </el-col>
            </el-row>

            <el-row :gutter="20">
              <el-col :span="12">
                <el-form-item label="内存:" >
                  <el-tag>{{ v["ram.current"] }}/{{ v["ram.max"] }}</el-tag>
                </el-form-item>
              </el-col>
              <el-col :span="12">
                <el-form-item label="堆内存:">
                  <el-tag>{{ v["heap.current"] }}/ {{ v["heap.max"] }}</el-tag>
                </el-form-item>
              </el-col>
            </el-row>

            <el-row :gutter="20">
              <el-col :span="12">
                <el-form-item label="磁盘:" >
                  <el-tag>{{ v["disk.used"] }}/ {{ v["disk.total"] }}</el-tag>
                </el-form-item>
              </el-col>
              <el-col :span="12">
                <el-form-item label="磁盘可用:">
                  <el-tag>
                    {{ v['disk.avail'] }}
                  </el-tag>
                </el-form-item>
              </el-col>
            </el-row>

            <el-row :gutter="20">
              <el-col :span="12">
                <el-form-item label="分片数:" >
                  <el-tag>{{ v["shards"] }}</el-tag>
                </el-form-item>
              </el-col>
              <el-col :span="12">
                <el-form-item label="索引所占空间大小:">

                    <el-tag>{{ v["disk.indices"] }}</el-tag>

                </el-form-item>
              </el-col>
            </el-row>

            <el-row :gutter="20">
              <el-col :span="12">
                <el-form-item label="CPU:" >
                  <el-progress
                      :stroke-width="24"
                      :color="getOpt"
                      :text-inside="true"
                      style="margin-right: 1px;width: 8rem"
                      :percentage="Number(v.cpu)"
                  />
                </el-form-item>
              </el-col>
              <el-col :span="12">
                <el-form-item label="内存:">
                    <el-progress
                        :stroke-width="24"
                        :color="getOpt"
                        :text-inside="true"
                        style="margin-right: 1px;width: 8rem"
                        :percentage="Number(v['ram.percent'])"
                    />

                </el-form-item>
              </el-col>
            </el-row>

            <el-row :gutter="20">
              <el-col :span="12">
                <el-form-item label="堆内存:" >
                  <el-progress
                      :stroke-width="24"
                      :color="getOpt"
                      :text-inside="true"
                      style="margin-right: 1px;width: 8rem"
                      :percentage="Number(v['heap.percent'])"
                  />
                </el-form-item>
              </el-col>
              <el-col :span="12">
                <el-form-item label="磁盘:">

                    <el-progress

                        :stroke-width="24"
                        :color="getOpt"
                        :text-inside="true"
                        style="width: 8rem"
                        :percentage="Number(v['disk.used_percent'])"
                    />

                </el-form-item>
              </el-col>
            </el-row>
          </el-form>

        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script lang="ts">

import {Component, Prop, toNative, Vue} from "vue-facing-decorator";

import {ElMessage} from "element-plus";
import {CatAction} from "../../../api/es";
import {sdk} from "@elasticview/plugin-sdk"

@Component({
  name: 'Node',
})
class Node extends Vue {
  list: any[] = []
  connectLoading = false
  value = 90

  activated() {

  }

  mounted() {
    console.log('Node mounted')
    this.searchData()
  }

  getOpt(v: any) {
    const numV = Number(v)

    if (numV < 60) {
      return '#4ad47f'
    } else if (numV >= 60 && numV <= 80) {
      return '#ff9800'
    }
    return 'red'
  }

  forceRefresh() {
    console.log('Force refresh triggered')
    this.searchData()
  }

  async searchData() {
    console.log('searchData called')
    this.connectLoading = true

    const res = await CatAction({
      cat: 'Node',
      es_connect: sdk.GetSelectEsConnID()
    })
    const res2 = await CatAction({
      cat: 'CatAllocation',
      es_connect: sdk.GetSelectEsConnID()
    })
    if (res.code != 0) {
      ElMessage.error({
        type: 'error',
        message: res.msg
      })

      this.connectLoading = false
      return
    }
    for (const k in res.data) {
      const v = res.data[k]
      for (const k2 in res2.data) {
        const v2 = res2.data[k2]
        if (v2['node'] == v['name']) {
          res.data[k]['disk.indices'] = v2['disk.indices']
          res.data[k]['shards'] = v2['shards']
          res.data[k]['disk.avail'] = v2['disk.avail']

          break
        }
      }
    }
    this.list = res.data
    this.connectLoading = false
  }

}
export default toNative(Node)

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
