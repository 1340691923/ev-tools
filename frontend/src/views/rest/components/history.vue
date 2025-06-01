<template>
  <div>
    <el-drawer

        size="80%"
        v-model="dialogVisible"
        :title="$t('历史记录')"
        @close="close"
    >
      <div class="search-container">
        <el-form :inline="true">
          <el-form-item label="索引名">
            <index-select style="width:140px"  :clearable="true" :placeholder="$t('请选择索引名')" @change="changeIndex" />
          </el-form-item>
          <el-form-item label="时间">
            <date  v-model="input.date"  />
          </el-form-item>

          <el-form-item>
            <el-button type="primary" @click="searchHistory(1)">查询</el-button>
          </el-form-item>
        </el-form>

      </div>
      <el-table v-loading="loading" :data="list">

        <el-table-column align="center" label="Method" width="220">
          <template #default="{ row }">
            {{ row.method }}
          </template>
        </el-table-column>
        <el-table-column align="center" label="Path" width="300">
          <template #default="{ row }">
            {{ row.path }}
          </template>
        </el-table-column>
        <el-table-column align="center" label="Body" width="300">
          <template #default="{ row }">
            <el-popover placement="top-start" title="body" width="600" trigger="hover">
              <div>{{ row.body }}</div>
              <template #reference>{{ row.body.substr(0, 50) + "..." }}</template>
            </el-popover>
          </template>
        </el-table-column>

        <el-table-column align="center" :label="$t('创建时间')" width="220">
          <template #default="{ row }">
            {{ row.created }}
          </template>
        </el-table-column>

        <el-table-column align="center" :label="$t('操作')" fixed="right" width="100">
          <template #header>
            <el-button type="danger" @click="clean">
              {{ $t('清空') }}
            </el-button>
          </template>
          <template #default="{ row }">
            <el-button type="success"  @click="getHistoryData(row)">
              {{ $t('搜索') }}
            </el-button>
          </template>
        </el-table-column>
      </el-table>
      <div class="pagination-container">
        <el-pagination
            hide-on-single-page
            :current-page="input.page"
            :page-sizes="[10, 20, 30, 50,100,150,200]"
            :page-size="input.limit"
            layout="total, sizes, prev, pager, next"
            :total="total"
            @size-change="handleSizeChange"
            @current-change="searchHistory"
        />
      </div>
    </el-drawer>
  </div>
</template>

<script>
import {ElMessage} from "element-plus";
import dayjs from "dayjs";
import {ref, onMounted, onActivated, watch, } from 'vue';
import { CleanAction, ListAction } from '@/api/dsl-history';
import IndexSelect from '@/components/index/select.vue'
import Date from '@/components/Date/index.vue'
import {sdk} from "@elasticview/plugin-sdk";

export default {
  name: 'History',
  components: {
    IndexSelect,
    Date
  },
  props: {
    dialogVisible: {
      type: Boolean,
      default: false,
    },
  },
  setup(props, { emit }) {
    const input = ref({
      indexName: '',
      date: [
        dayjs().format('YYYY-MM-DD 00:00:00'),
        dayjs().format('YYYY-MM-DD 23:59:59')
      ],
      page: 1,
      limit: 10,
    });
    const list = ref([]);
    const total = ref(0);
    const dialogVisible = ref(props.dialogVisible)
    const loading = ref(false)
    const searchHistory = async (page = 1) => {
      input.value.page = page
      input.value.date = [
        dayjs(input.value.date[0]).format('YYYY-MM-DD HH:mm:ss'),
        dayjs(input.value.date[1]).format('YYYY-MM-DD HH:mm:ss')
      ]
      loading.value = true
      const { data, code, msg } = await ListAction(input.value);
      loading.value = false
      if (code !== 0) {
        ElMessage({ type: 'error', message: msg });
      } else {
        list.value = data.list;
        total.value = data.count;
      }
    };

    const handleSizeChange = (v) => {
      input.value.limit = v;
      searchHistory(1);
    };

    const changeDate = (v) => {
      input.value.date = v;
      searchHistory(1);
    };

    const changeIndex = (v) => {
      input.value.indexName = v;
      searchHistory(1);
    };

    const getHistoryData = (row) => {
      emit('getHistoryData', row);
      emit('close', false);
    };

    const clean = async () => {
      const { code, msg } = await CleanAction(input.value);
      if (code !== 0) {
        ElMessage({ type: 'error', message: msg });
      } else {
        ElMessage({ type: 'success', message: msg });
        searchHistory(1);
      }
    };

    const close = () => {
      emit('close', false);
    };

    onMounted(() => {
      console.log('history mounted')
      searchHistory(1);
    });

    onActivated(() => {
      console.log('history activated')
    });

    watch(() => props.dialogVisible, (newVal) => {
      if (newVal) searchHistory(1);
    });

    return {
      input,
      list,
      total,
      handleSizeChange,
      changeDate,
      changeIndex,
      getHistoryData,
      clean,
      close,
      searchHistory,
      dialogVisible,
      loading,
    };
  },
};
</script>

<style scoped>
/* Add any styles here */
</style>
