<template>
  <div>
    <div class="app-container">

      <div class="search-container">

        <el-form :inline="true" >
          <el-form-item label="关键词">
            <el-input v-model="input"  style="width: 7rem" clearable @input="search"/>

          </el-form-item>
          <el-form-item label="">
            <el-button
                type="primary"  @click="search">{{ $t('搜索') }}
            </el-button>
          </el-form-item>
        </el-form>


      </div>
      <el-card shadow="never" class="table-container">
        <el-table
            v-loading="connectLoading"
            :data="list"
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

          <el-table-column show-overflow-tooltip v-for="(info,index) in tableInfo" :key="index" align="center" :label="tableInfo[index].desc"
                           :width="info.width" :prop="info.data.to$t()" :sortable="info.sort">
            <template #default="scope">
              {{ scope.row[info.data.split('->').join('.')] }}
            </template>
          </el-table-column>

        </el-table>
        <el-pagination
            v-if="pageshow"
            class="pagination-container"
            :current-page="page"
            :page-sizes="[10, 20, 30, 50,100,150,200,500,1000]"
            :page-size="limit"
            layout="total, sizes, prev, pager, next, jumper"
            :total="total"
            @size-change="handleSizeChange"
            @current-change="handleCurrentChange"
        />
      </el-card>
    </div>
  </div>
</template>

<script lang="ts">

import {Component, Prop, toNative, Vue} from "vue-facing-decorator";
import {filterData} from "../../../utils/table";
import {getCurrentInstance} from "vue";
import {CatAction} from "../../../api/es";
import {ElMessage} from "element-plus";
import {sdk} from "@/plugin_sdk/sdk"

@Component({
  name: 'CatIndex',
  setup() {
    const ctx = getCurrentInstance().appContext.config.globalProperties
    return {ctx}
  },
})
class CatIndex extends Vue {

  @Prop
  catType = ''

  @Prop
  tableInfo = []

  ctx = null

  total= 0
  connectLoading=false
  page=1
  limit=10
  pageshow= true
  list=[]
  input= ''
  allList= []

  pageLimit() {
    this.list = this.allList.filter((item, index) =>
        index < this.page * this.limit && index >= this.limit * (this.page - 1)
    )
  }
  search() {
    this.page = 1
    this.pageshow = false
    this.searchData()
    this.$nextTick(() => {
      this.pageshow = true
    })
  }

  // 当每页数量改变
  handleSizeChange(val) {
    this.limit = val
    this.pageLimit()
  }
  // 当当前页改变
  handleCurrentChange(val) {
    this.page = val
    this.pageLimit()
  }
  searchData() {

    this.connectLoading = true
    const form = {
      cat: this.catType,
      es_connect: sdk.GetSelectEsConnID()
    }
    CatAction(form).then(res => {
      if (res.code == 0) {
        let list = res.data
        if (list == null) {
          return
        }
        for (const index in list) {
          const obj = list[index]
          // 把 . 转成 ->
          for (const key in obj) {
            let value = Number(obj[key])
            if (isNaN(value)) {
              value = obj[key]
            }
            list[index][key.split('.').join('->')] = value
          }
        }

        list = filterData(list, this.input.trim())
        this.allList = list
        this.total = list.length
        this.pageLimit()

      } else {
        ElMessage.error({
          type: 'error',
          message: res.msg
        })
      }
      this.connectLoading = false
    }).catch(err => {
      console.log(err)

      this.connectLoading = false
    })
  }

  mounted() {
    this.searchData()
  }
}

export default toNative(CatIndex)
</script>

