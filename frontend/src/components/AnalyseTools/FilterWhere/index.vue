<template>
  <div>
    <div class="xwl_main" style="margin-top: 20px">
      <div class="relation-editor globalFilters_xwl">
        <div class="relation-relation">
          <em class="relation-relation-line" />
          <div
              v-if="userFilter.filts.length > 1"
              class="relation-relation-value"
              @click="changeRelationLine"
          >
            {{ userFilter.relation }}
          </div>
        </div>
        <div class="relation-main">
          <div
              v-for="(v, index) in userFilter.filts"
              :key="index"
              class="relation-row"
          >
            <div class="ta-multa-filter-condition">
              <div v-if="v.filterType === 'SIMPLE'" class="action-row row___xwl">

                <action-row
                    v-if="changeActionRow"
                    v-model:value="userFilter.filts[index]"
                    :options="props.options"
                    :table-typ="props.tableTyp"
                    :data-type-map="props.dataTypeMap"
                    class="action-left"
                />
                <div class="action-right">
                  <el-button-group>
                    <el-button
                        type="link"
                        class="actions_xwl_btn"
                        @click="() => addRelationSimple(index)"
                        :icon="Filter"
                    >

                    </el-button>
                    <el-button
                        type="link"
                        class="actions_xwl_btn"
                        :icon="Delete"
                        @click="() => deleteRelationSimple(index)"
                    >

                    </el-button>
                  </el-button-group>
                </div>
              </div>
              <div v-else class="relation-editor">
                <div class="relation-relation">
                  <div class="relation-relation-sub">◆</div>
                  <em class="relation-relation-line-sub" />
                  <div
                      class="relation-relation-value"
                      @click="() => changeRelationLine1(index)"
                  >
                    {{ v.relation }}
                  </div>
                </div>
                <div class="relation-main">
                  <div class="relation-row">
                    <div
                        v-for="(v2, index2) in v.filts"
                        :key="index2"
                        class="action-row row___xwl"
                    >
                      <div class="action-left">

                        <action-row
                            :key="(index + '_' + index2).to$t()"
                            v-model:value="v.filts[index2]"
                            :options="props.options"
                            :table-typ="props.tableTyp"
                            :data-type-map="props.dataTypeMap"
                            class="action-left"
                        />
                      </div>
                      <div class="action-right">
                        <el-button-group>
                          <el-button
                              type="link"
                              class="actions_xwl_btn"
                              @click="() => addRelationCOMPOUND(index)"
                              :icon="Filter"
                          >

                          </el-button>
                          <el-button
                              type="link"
                              class="actions_xwl_btn"
                              :icon="Delete"
                              @click="() => deleteRelationCOMPOUND(index, index2)"
                          >

                          </el-button>
                        </el-button-group>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
    <div style="padding: 0 12px;">
      <span
          :style="{ color: props.fontColor }"
          class="footadd___2D4YB"
          @click="addRelation1"
      >
        {{ $t('增加条件') }}
      </span>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive, nextTick, toRefs } from 'vue';
import { ElMessage } from 'element-plus';
import ActionRow from '@/components/AnalyseTools/FilterWhere/ActionRow.vue';
import {Filter,Delete} from '@element-plus/icons-vue'

// Props
const props = defineProps({
  value: {
    type: Object,
    default: () => ({
      relationLine: '且',
      relationArr: [],
      ftv: [],
    }),
  },
  fontColor: {
    type: String,
    default: '#3d90ff',
  },
  tableTyp: {
    type: Number,
    default: 0,
  },
  dataTypeMap: {
    type: Array,
    default: () => [],
  },
  options: {
    type: Array,
    default: () => [],
  },
});

// Emits
const emit = defineEmits(['update:value']);

// Reactive state
const changeActionRow = ref(true);
const userFilter = reactive({ ...props.value });

// Methods
const refreshChangeActionRow = () => {
  changeActionRow.value = false;
  nextTick(() => {
    changeActionRow.value = true;
  });
};

const addRelationSimple = (index) => {
  const obj = userFilter.filts[index];
  const COMPOUNDObj = {
    filterType: 'COMPOUND',
    filts: [],
    relation: '且',
  };

  COMPOUNDObj.filts.push(obj);
  COMPOUNDObj.filts.push({
    columnName: props.options[0].options[0].value,
    comparator: '=',
    filterType: 'SIMPLE',
    ftv: '',
  });

  userFilter.filts[index] = COMPOUNDObj;
  emit('update:value', userFilter);
};

const addRelationCOMPOUND = (index) => {
  userFilter.filts[index].filts.push({
    columnName: props.options[0].options[0].value,
    comparator: '=',
    filterType: 'SIMPLE',
    ftv: '',
  });
  emit('update:value', userFilter);
};

const deleteRelationCOMPOUND = (index1, index2) => {
  if (userFilter.filts[index1].filts.length > 2) {
    userFilter.filts[index1].filts.splice(index2, 1);
  } else {
    const arr = [...userFilter.filts[index1].filts];
    arr.splice(index2, 1);
    const obj = arr[0];
    userFilter.filts[index1] = obj;
  }

  emit('update:value', userFilter);
};

const addRelation1 = () => {

  if (props.options.length == 0 ) {
    emit('update:value', userFilter);
    ElMessage({
      offset: 60,
      message: '该索引没有字段，无法筛选',
      type: 'error',
    });
    return;
  }

  if (props.options[0].options.length === 0) {
    emit('update:value', userFilter);
    ElMessage({
      offset: 60,
      message: '该索引没有字段，无法筛选',
      type: 'error',
    });
    return;
  }
  userFilter.filts.push({
    columnName: props.options[0].options[0].value,
    comparator: '=',
    filterType: 'SIMPLE',
    ftv: '',
  });
  emit('update:value', userFilter);
};

const deleteRelationSimple = (i) => {
  userFilter.filts.splice(i, 1);
  emit('update:value', userFilter);
  refreshChangeActionRow();
};

const changeRelationLine = () => {
  userFilter.relation = userFilter.relation === '或' ? '且' : '或';
  emit('update:value', userFilter);
};

const changeRelationLine1 = (index) => {
  userFilter.filts[index].relation =
      userFilter.filts[index].relation === '或' ? '且' : '或';
  emit('update:value', userFilter);
};
</script>


<style lang="scss" scoped>
.sider_xwl {
  display: inline-block;
  width: 270px;
  height: calc(100vh - 50px);
  padding: 0;
  overflow: hidden;
  vertical-align: top;
}

.top_xwl {
  display: flex;
  align-items: center;
  justify-content: flex-start;
  width: 100%;
  height: 57px;
  border-bottom: 1px solid #f0f0f0;
  border-right: 1px solid #f0f0f0;
  padding-right: 20px;
}

.body_xwl {
  border-bottom: 1px solid #f0f0f0;
  border-right: 1px solid #f0f0f0;
  height: calc(100vh - 120px);
  overflow-y: auto;
  background-color: white;
}

.content_xwl {
  display: inline-block;
  width: 100%;
  height: calc(100vh - 50px);
  padding: 0;
  overflow: hidden;
  vertical-align: top;
  background: #f0f2f5;
  border-bottom: 2px solid #e6e6e6;
}

.header_xwl {
  position: relative;
  z-index: 1000;
  width: 100%;
  height: 56px;
  padding: 0;
  line-height: 56px;

}

.main_xwl {
  display: flex;
  align-items: center;
}

.title_xwl {
  display: inline-block;
  margin-right: 16px;
  color: black;
  font-weight: 500;
  font-size: 18px;
}

.dashbordName_xwl {
  display: flex;
  align-items: center;
  justify-content: space-between;
  height: 40px;

  border-radius: 2px;
  cursor: pointer;
}

.root_xwl {
  display: flex;
  align-items: center;
  justify-content: space-between;
  width: 100%;
  height: 56px;
  padding: 0 16px;
  background-color: #fff;
  border-left: 1px solid #f0f0f0;
}

.actions_xwl {
  display: flex;
  align-items: center;
}

.actions_xwl_btn:hover {
  color: orangered;
}

.actions_xwl_btn {
  color: #67729d;
}

.echartBox {
  width: 48%;
  height: 400px;
  padding: 5px;
  margin-bottom: 5px;
}

.echartBox:hover {
  cursor: pointer;
}

.echartBox_title {
  cursor: pointer;
}

.echartBox_title:hover {
  color: #3d90ff;
}

.footadd___2D4YB {
  display: inline-block;
  margin-right: 16px;
  padding: 4px;
  color: #3d90ff;
  font-size: 13px;
  line-height: 20px;
  border-radius: 2px;
  cursor: pointer;
  transition: all .3s;
  transition-property: all;
  transition-duration: 0.3s;
  transition-timing-function: ease;
  transition-delay: 0s;
}

.right_res {
  position: relative;
  width: 100%;
  height: calc(100% - 32px);

  overflow-x: hidden;
  overflow-y: auto;
  white-space: normal;
  transition: width .3s;
}

.row___xwl {
  min-height: 40px;
  padding: 0 4px 0 8px;
  transition-property: all;
  transition-duration: 0.3s;
  transition-timing-function: ease;
  transition-delay: 0s;
  padding: 10px;
}

.row___xwl:hover {
  box-shadow: 0 0 3px 0 #1890ff;
  transition-property: all;
  transition-duration: 0.3s;
  transition-timing-function: ease;
  transition-delay: 0s;
}

.action-row {
  position: relative;
  display: flex;
  align-items: center;
  justify-content: space-between;
  width: 100%;
  min-height: 24px;
}

.relation-relation-line {
  position: absolute;
  top: 0;
  left: 12px;
  width: 1px;
  height: 100%;
  background-color: #d9dfe6;
  transition-property: all;
  transition-duration: 0.3s;
  transition-timing-function: ease;
  transition-delay: 0s;
}

.relation-relation-value {
  position: absolute;
  top: 50%;
  left: 0;
  width: 24px;
  height: 24px;
  margin-top: -12px;
  color: #3d90ff;
  font-size: 12px;
  line-height: 22px;
  text-align: center;
  background-color: #fff;
  border: 1px solid #d9dfe6;
  border-radius: 2px;
  cursor: pointer;
  transition-property: all;
  transition-duration: 0.3s;
  transition-timing-function: ease;
  transition-delay: 0s;
}

.xwl_main {

}

.globalFilters_xwl {
  padding: 0 8px 0 12px;
}

.relation-editor {
  position: relative;
  display: flex;
  width: 100%;
}

.relation-editor .relation-relation {
  position: relative;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: flex-start;
  width: 24px;
  margin-right: 4px;
}

.relation-relation-line {
  position: absolute;
  top: 0;
  left: 12px;
  width: 1px;
  height: 100%;
  background-color: #d9dfe6;
  transition-property: all;
  transition-duration: 0.3s;
  transition-timing-function: ease;
  transition-delay: 0s;
}

.relation-editor .relation-relation .relation-relation-value {
  position: absolute;
  top: 50%;
  left: 0;
  width: 24px;
  height: 24px;
  margin-top: -12px;
  color: #3d90ff;
  font-size: 12px;
  line-height: 22px;
  text-align: center;
  background-color: #fff;
  border: 1px solid #d9dfe6;
  border-radius: 2px;
  cursor: pointer;
  transition-property: all;
  transition-duration: 0.3s;
  transition-timing-function: ease;
  transition-delay: 0s;
}

.relation-editor .relation-main {
  flex-grow: 1;
  flex-shrink: 1;
  flex-basis: 0%;
}

.relation-editor .relation-relation .relation-relation-sub {
  position: relative;
  top: -7px;
  left: .5px;
  z-index: 100;
  color: #a3acc5;
  font-size: 11px;
}

.relation-editor .relation-relation .relation-relation-line-sub {
  position: absolute;
  top: 13px;
  left: 12px;
  width: 1px;
  height: calc(100% - 10px);
  background-color: #d9dfe6;
  transition-property: all;
  transition-duration: 0.3s;
  transition-timing-function: ease;
  transition-delay: 0s;
}

.ta-multa-filter-condition {
  padding: 0;
}

.relation-row .action-row {
  align-items: center;
  justify-content: flex-start;
}

.row___xwl {
  min-height: 32px;
  padding: 2px 4px;
}

.action-row .action-left {
  display: flex;
  align-items: center;
}

.ta-filter-condition {
  display: inline-block;
  min-height: 32px;
  padding-bottom: 2px;
  white-space: normal;
}

.action-row .action-right {
  display: flex;
  align-items: center;
}

.relation-relation:hover {
  cursor: pointer;
}

.relation-relation:hover .relation-relation-line {
  background-color: #3d90ff;
}

.relation-relation:hover .relation-relation-value {
  background-color: #3d90ff;
  color: white;
}

.relation-relation:hover .relation-relation-line-sub {
  background-color: #3d90ff;
}

.relation-relation:hover .relation-relation-sub {
  color: #3d90ff;
}

::-webkit-scrollbar {
  width: 8px; /*对垂直流动条有效*/
  height: 10px; /*对水平流动条有效*/
}

/*定义滚动条的轨道颜色、内阴影及圆角*/
::-webkit-scrollbar-track {
  border-radius: 3px;
}

/*定义滑块颜色、内阴影及圆角*/
::-webkit-scrollbar-thumb {
  border-radius: 7px;
  background-color: #e6e6e6;
}

/*定义两端按钮的样式*/
::-webkit-scrollbar-button {
}

/*定义右下角汇合处的样式*/
::-webkit-scrollbar-corner {
  background: khaki;
}

.eventRow_xwl {
  position: relative;
}

.rename_xwl {
  width: auto;
  min-width: 50px;
  max-width: 260px;
  height: 24px;
  margin: 8px 0 0;
  padding: 0;
  line-height: 24px;
  background: inherit;
}

.rename_xwl span {
  max-width: 260px;
  font-weight: 500;
  font-size: 13px;
}

.rename_xwl input {
  font-weight: 500;
  font-size: 13px;
}

.eventRow_xwl .eventItem_xwl {
  max-width: 400px;
  padding: 2px 0 6px;
  overflow: hidden;
  white-space: normal;
}

.filters_xwl {
  padding-left: 32px;
}

.formula_xwl {
  padding-right: 8px;
  padding-left: 32px;
}

</style>
