<template>
  <div class="app-container">
    <div class="search-container">
      <el-form>
        <el-form-item>
        <el-button
          type="success"

          @click="refresh"
        >{{ $t('清空表单') }}
        </el-button>
        <el-button

          @click="commit"
        >{{ $t('确认重建索引') }}
        </el-button>
        </el-form-item>
      </el-form>
    </div>

    <el-tabs  v-model="activeName"  stretch tab-position="right">
      <el-tab-pane name="url参数" :lazy="true" :label="$t('url参数')">
        <el-card #default class="box-card card">
          <el-tag effect="dark" type="success">{{ $t('url参数') }}</el-tag>
          <div class="margin-20" />
          <el-form label-position="right">
            <el-form-item :label="$t('配置')">
              <el-checkbox-group v-model="urlParmasConfig">
                <el-checkbox v-for="(v,k,index) in form.urlParmas" :key="index" :label="k" />
              </el-checkbox-group>
            </el-form-item>
            <el-form-item
              v-if="urlParmasConfig.indexOf('requests_per_second') != -1"
              label="requests_per_second (每秒数据量阈值控制)"
            >
              <el-input v-model="form.urlParmas.requests_per_second" type="number" />
            </el-form-item>
            <el-form-item
              v-if="urlParmasConfig.indexOf('wait_for_active_shards') != -1"
              label="wait_for_active_shards (重建索引分片响应设置)"
            >
              <el-input v-model="form.urlParmas.wait_for_active_shards" />
            </el-form-item>
            <el-form-item v-if="urlParmasConfig.indexOf('scroll') != -1" label="scroll (快照查询时间)">
              <el-input v-model="form.urlParmas.scroll" clearable />
            </el-form-item>
            <el-form-item v-if="urlParmasConfig.indexOf('slices') != -1" label="slices (重建并行任务切片)">
              <el-input v-model="form.urlParmas.slices" type="number" />
            </el-form-item>

            <el-form-item v-if="urlParmasConfig.indexOf('refresh') != -1" label="refresh (目标索引是否立即刷新)">
              <el-select v-model="form.urlParmas.refresh" filterable>
                <el-option label="true" :value="true" />
                <el-option label="false" :value="false" />

              </el-select>
            </el-form-item>

            <el-form-item v-if="urlParmasConfig.indexOf('timeout') != -1" label="timeout (指定等待响应的时间段，单位秒)">
              <el-input v-model.number="form.urlParmas.timeout" style="width: 300px" clearable />
              <el-tag>秒</el-tag>
            </el-form-item>
            <el-form-item
              v-if="urlParmasConfig.indexOf('wait_for_completion') != -1"
              label="wait_for_completion (是否将请求阻塞，直到操作完成为止)"
            >
              <el-select v-model="form.urlParmas.wait_for_completion" filterable>
                <el-option :label="是" :value="true" />
                <el-option :label="否" :value="false" />
              </el-select>
            </el-form-item>

          </el-form>
        </el-card>
      </el-tab-pane>
      <el-tab-pane name="索引" :lazy="true" :label="$t('索引')">
        <template #default>

          <el-card class="box-card card" style="width: 45%;">

            <el-tag effect="dark">{{ $t('源索引') }}</el-tag>
            <div class="margin-20" />
            <el-form>

              <el-form-item :label="$t('配置')">
                <el-checkbox-group v-model="sourceConfig">
                  {{form.source}}
                  <el-checkbox v-for="(v,k,index) in form.source" v-if="k != 'index'" :key="index" :label="k" />
                </el-checkbox-group>
              </el-form-item>
              <el-form-item v-if="sourceConfig.indexOf('remote') != -1">
                <el-form>
                  <el-form-item :label="$t('从连接树中连接')">
                    <el-select
                        v-model="form.source.remote.link"
                        filterable
                        default-first-option
                        :placeholder="$t('请选择ES连接')"
                        @change="changeLink"
                    >
                      <el-option v-for="item in linkOpt" :key="item.id" :value="item.remark" :label="item.remark" />
                    </el-select>
                  </el-form-item>
                  <el-form-item label="host (主机)">
                    <el-input v-model="form.source.remote.host" placeholder="主机" clearable />
                  </el-form-item>
                  <el-form-item label="username (用户名)">
                    <el-input v-model="form.source.remote.username" placeholder="用户名" clearable />
                  </el-form-item>
                  <el-form-item label="password (密码)">
                    <el-input v-model="form.source.remote.password" placeholder="密码" clearable />
                  </el-form-item>
                  <el-form-item label="socket_timeout">
                    <el-input v-model="form.source.remote.socket_timeout" placeholder="socket_timeout" clearable />
                  </el-form-item>
                  <el-form-item label="connect_timeout">
                    <el-input
                        v-model="form.source.remote.connect_timeout"
                        placeholder="connect_timeout"
                        clearable
                    />
                  </el-form-item>
                </el-form>
              </el-form-item>
              <el-form-item :label="$t('索引名')">
                <el-select
                    v-model="form.source.index"
                    :placeholder="$t('请选择索引名')"
                    filterable
                    allow-create
                    multiple
                    clearable
                    @change="changeSourceIndex"
                >
                  <el-option

                      v-for="(item,index) of indexNameList"
                      :key="index"
                      :label="item"
                      :value="item"
                  />
                </el-select>
              </el-form-item>
              <el-form-item v-if="sourceConfig.indexOf('type') != -1" label="type (类型名)">
                <el-select
                    v-model="form.source.type"
                    :placeholder="$t('请选择类型名')"
                    filterable
                    allow-create
                    multiple
                    clearable
                >
                  <el-option
                      v-for="(item,index) of mappingConfig"
                      :key="index"
                      :label="item"
                      :value="item"
                  />
                </el-select>
              </el-form-item>
              <el-form-item v-if="sourceConfig.indexOf('_source') != -1" label="_source (字段列表)">
                <el-select
                    v-model="form.source._source.excludes"
                    :placeholder="$t('排除字段')"
                    filterable
                    allow-create
                    multiple
                    clearable
                />
                <el-select
                    v-model="form.source._source.includes"
                    :placeholder="$t('包含字段')"
                    filterable
                    allow-create
                    multiple
                    clearable
                />
              </el-form-item>
              <el-form-item v-if="sourceConfig.indexOf('sort') != -1" :label="$t('sort (排序)')">
                <el-tag effect="dark">{{ $t('字段') }}</el-tag>
                <el-input v-model="form.source.sort.key" class="sort-key" />
                <el-tag effect="dark">{{ $t('是否正序') }}</el-tag>
                <el-select
                    v-model="form.source.sort.sortType"
                    class="sort-key"
                    :placeholder="$t('是否正序')"
                    filterable
                >
                  <el-option
                      key="index"
                      :label="$t('是')"
                      value="asc"
                  />
                  <el-option
                      key="index"
                      :label="$t('否')"
                      value="desc"
                  />
                </el-select>
              </el-form-item>
              <el-form-item v-if="sourceConfig.indexOf('size') != -1" :label="$t('批次大小配置')">
                <el-input v-model="form.source.size" type="number" class="width30" />
              </el-form-item>
              <el-form-item v-if="sourceConfig.indexOf('query') != -1" label="QUERY DSL">

                <json-editor

                    v-model:value="form.source.query"
                    styles="width: 100%"
                    :point-out="pointOut"
                    :read="false"
                    title="QUERY"
                    @getValue="getBody"
                />

              </el-form-item>
            </el-form>
          </el-card>
          <el-card class="box-card card" style="width: 45%;">
            <el-tag effect="dark" type="danger">{{ $t('目标索引') }}</el-tag>
            <div class="margin-20" />
            <el-form>
              <el-form-item :label="$t('配置')">
                <el-checkbox-group v-model="destConfig">
                  <el-checkbox v-for="(v,k,index) in form.dest" v-if="k != 'index'" :key="index" :label="k" />
                </el-checkbox-group>
              </el-form-item>

              <el-form-item :label="$t('索引名')">
                <el-select v-model="form.dest.index" :placeholder="$t('请选择索引名')" clearable filterable>
                  <el-option
                      v-for="(item,index) of indexNameList"
                      :key="index"
                      :label="item"
                      :value="item"
                  />
                </el-select>
              </el-form-item>
              <el-form-item v-if="destConfig.indexOf('version_type') != -1" :label="$t('version_type(用于索引操作的版本控制)')">
                <el-select v-model="form.dest.version_type" :placeholder="$t('请选择version_type')" filterable>
                  <el-option label="external" value="external" />
                  <el-option label="internal" value="internal" />
                  <el-option label="external_gt" value="external_gt" />
                  <el-option label="external_gte" value="external_gte" />
                </el-select>
              </el-form-item>
              <el-form-item v-if="destConfig.indexOf('op_type') != -1" :label="$t('op_type(仅创建不存在的索引文档)')">
                <el-select v-model="form.dest.op_type" :placeholder="$t('op_type(仅创建不存在的索引文档)')" filterable>
                  <el-option
                      label="create"
                      value="create"
                  />
                  <el-option
                      label="index"
                      value="index"
                  />
                </el-select>
              </el-form-item>
              <el-form-item v-if="destConfig.indexOf('routing') != -1" :label="$t('routing')">
                <el-input v-model="form.dest.routing" class="width30" clearable />
              </el-form-item>
              <el-form-item v-if="destConfig.indexOf('pipeline') != -1" :label="$t('pipeline')">
                <el-input v-model="form.dest.pipeline" class="width30" clearable />
              </el-form-item>
            </el-form>
          </el-card>
        </template>
      </el-tab-pane>
      <el-tab-pane name="脚本" :lazy="true" :label="$t('脚本')">
        <template #default>
          <el-card class="box-card card">
            <el-tag effect="dark" type="warning">{{ $t('脚本') }}</el-tag>
            <div class="margin-20" />
            <el-form>
              <el-form-item :label="$t('配置')">
                <el-checkbox-group v-model="scriptConfig">
                  <el-checkbox v-for="(v,k,index) in form.script" v-if="k != 'index'" :key="index" :label="k" />
                </el-checkbox-group>
              </el-form-item>
              <el-form-item v-if="scriptConfig.indexOf('source') != -1" :label="$t('script.source 定义脚本语言')">
                <el-input v-model="form.script.source" clearable />
              </el-form-item>
              <el-form-item v-if="scriptConfig.indexOf('lang') != -1" :label="$t('script.lang 定义脚本实现的代码')">
                <el-input v-model="form.script.lang" clearable />
              </el-form-item>
              <el-form-item v-if="scriptConfig.indexOf('params') != -1" :label="$t('定义脚本语言的参数')">
                <json-editor
                    v-model:value="form.script.params"
                    style="height: 300px"
                    styles="width: 100%"
                    :read="false"
                    :title="$t('定义脚本语言的参数')"
                    @getValue="getParams"
                />
              </el-form-item>
            </el-form>
          </el-card>
        </template>
      </el-tab-pane>
      <el-tab-pane name="额外参数"  :lazy="true" :label="$t('额外参数')">
        <template #default>
          <el-card class="box-card card">
            <el-tag effect="dark" type="info">{{ $t('额外参数') }}</el-tag>
            <div class="margin-20" />
            <el-form>
              <el-form-item :label="$t('配置')">
                <el-checkbox-group v-model="extendConfig">
                  <el-checkbox v-for="(v,k,index) in form.extend" v-if="k != 'index'" :key="index" :label="k" />
                </el-checkbox-group>
              </el-form-item>
              <el-form-item v-if="extendConfig.indexOf('conflicts') != -1" label="conflicts(版本冲突时中止)">
                <el-select v-model="form.extend.conflicts" filterable>
                  <el-option label="proceed(遇到冲突时继续)" value="proceed" />
                  <el-option label="abort(遇到冲突时中止)" value="internal" />
                </el-select>
              </el-form-item>
              <el-form-item v-if="extendConfig.indexOf('size') != -1" label="size (每批要编制索引的文档数)">
                <el-input v-model="form.extend.size" type="number" class="width30" />
              </el-form-item>
            </el-form>

          </el-card>
        </template>
      </el-tab-pane>
    </el-tabs>
  </div>
</template>

<script>
import { esBodyKeyWords } from '@/utils/base-data'
import { ListAction } from '@/api/es-map'
import { ReindexAction } from '@/api/es-index'

import JsonEditor from '@/components/JsonEditor/index.vue';
import {sdk} from "@elasticview/plugin-sdk"

import {ElMessage,ElNotification} from "element-plus";

export default {
  name: 'Reindex',
  components: {
    JsonEditor
  },
  setup(){

  },
  data() {
    return {
      activeName:'url参数',
      ctx:null,
      es_connect: 0,
      extendConfig: [],
      typeList: [],
      pointOut: esBodyKeyWords,
      urlParmasConfig: [
          'requests_per_second',
          'timeout','slices','scroll','wait_for_active_shards','refresh','wait_for_completion'
      ],
      sourceConfig: [],
      destConfig: [],
      scriptConfig: [],
      indexNameList: [],
      form: {
        urlParmas: {
          timeout: 30,
          requests_per_second: -1,
          slices: 5,
          scroll: '',
          wait_for_active_shards: '',
          refresh: false,
          wait_for_completion: true
        },
        source: {
          index: [],
          size: 5000,
          type: [],
          query: '{}',
          _source: {
            excludes: [],
            includes: []
          },
          sort: {
            key: '',
            sortType: 'asc'
          },
          remote: {
            link: '',
            host: 'http://127.0.0.1:9200', // 远程es的ip和port列表
            username: 'elastic',
            password: '',
            socket_timeout: '',
            connect_timeout: '' // 超时时间设置
          }
        },
        dest: {
          index: '',
          version_type: '',
          op_type: 'index',
          routing: '=cat',
          pipeline: 'some_ingest_pipeline'
        },
        script: {
          source: "if (ctx._source.foo == 'bar') {ctx._version++; ctx._source.remove('foo')}",
          lang: 'painless',
          params: '{}'
        },
        extend: {
          conflicts: 'proceed',
          size: 0
        }
      },
      mappings: {},
      mappingConfig: [],
      linkOpt: []
    }
  },
  activated(){
    console.log('reindex activated')
  },
  mounted() {
    console.log('reindex mounted')
    this.es_connect = sdk.GetSelectEsConnID()
    this.getLinkList()
    this.getIndexList()
  },

  methods: {
    changeLink(link) {
      for (const v of this.linkOpt) {
        if (v.remark == link) {
          this.form.source.remote.host = v.ip
          this.form.source.remote.username = v.user
          this.form.source.remote.password = v.pwd
        }
      }
    },
    async getLinkList() {
      const res = await sdk.LinkOptAction()
      this.linkOpt = res.data
    },
    async commit() {
      const body = {
        source: {
          index: this.form.source.index
        },
        dest: {
          index: this.form.dest.index
        }
      }

      if (this.scriptConfig.length > 0) {
        body['script'] = {}
      }
      const urlParmas = {}

      for (const extendConfig of this.extendConfig) {
        if (extendConfig == 'size') {
          if (this.form.extend[extendConfig] == 0) {
            continue
          }
        }
        body[extendConfig] = this.form.extend[extendConfig]
      }
      for (const urlParma of this.urlParmasConfig) {
        urlParmas[urlParma] = this.form.urlParmas[urlParma]
      }
      for (const source of this.sourceConfig) {
        switch (source) {
          case 'sort':
            body['source']['sort'] = {}
            body['source']['sort'][this.form.source.sort.key] = this.form.source.sort.sortType
            break
          case 'query':
            try {
              body['source']['query'] = JSON.parse(this.form.source.query)
            } catch (e) {
              console.log(e)
              ElMessage({
                type: 'error',
                message: 'query json解析失败'
              })
              return
            }
            break
          case 'remote':
            body['source'][source] = this.form.source[source]
            delete body['source'][source]['link']
            if (this.form.source.remote.socket_timeout == '' || this.form.source.remote.connect_timeout == '') {
              delete body['source'][source]['socket_timeout']
              delete body['source'][source]['connect_timeout']
            }
          default:
            body['source'][source] = this.form.source[source]
        }
      }
      for (const dest of this.destConfig) {
        body['dest'][dest] = this.form.dest[dest]
      }
      for (const script of this.scriptConfig) {
        if (script == 'params') {
          try {
            body['script']['params'] = JSON.parse(this.form.script.params)
          } catch (e) {
            console.log(e)
            ElMessage({
              type: 'error',
              message: '脚本语言参数 json解析失败'
            })
            return
          }
        } else {
          body['script'][script] = this.form.script[script]
        }
      }
      console.log(body, 'body')

      const input = {}
      input['url_values'] = urlParmas
      input['body'] = body
      input['es_connect'] = this.es_connect

      const { data, code, msg } = await ReindexAction(input)
      if (code == 0) {
        if (data.hasOwnProperty('task')) {
          ElNotification({
            offset: 100,
            position: 'top-left',
            title: 'Success',
            dangerouslyUseHTMLString: true,
            message: `
                        <div>异步任务id: ${data.task}</div>
                      `,
            type: 'success'
          })
          return
        }

        ElNotification({
          offset: 100,
          position: 'top-left',
          title: 'Success',
          dangerouslyUseHTMLString: true,
          message: `
                        <div>整个操作花费的总毫秒数: ${data.took}</div>
                        <div>由重新索引拉回的滚动响应数: ${data.batches}</div>
                        <div style="color: green">成功创建的文档数: ${data.created}</div>
                        <div  style="color: red">成功删除的文档数: ${data.deleted}</div>
                        <div>在重新索引期间执行的任何请求超时: ${data.timed_out}</div>
                         <div>成功处理的文档数: ${data.total}</div>
                        <div>重新索引命中的版本冲突数: ${data.version_conflicts}</div>
                        <div>由于用于重新索引的脚本返回的noop值而被忽略的文档数: ${data.noops}</div>
                         <div>重试的批量操作search数: ${data.retries.bulk}</div>
                        <div>重试的搜索操作数: ${data.retries.search}</div>

                         <div>请求睡眠以符合的毫秒数: ${data.throttled_millis}</div>
                        <div>在重新索引期间每秒有效执行的请求数: ${data.requests_per_second}</div>

                          <div>失败数组: ${JSON.stringify(data.failures, null, '\t')}</div>
                      `,
          type: 'success'
        })
      } else {
        if (msg.indexOf('whitelisted') != -1) {
          ElMessage({
            type: 'error',
            message: msg + '在源es与目标es上都需要进行配置白名单，具体解决办法 :https://my.oschina.net/xiaominmin/blog/1627579'
          })
          return
        } else {
          ElMessage({
            type: 'error',
            message: msg
          })
          return
        }
      }
    },
    refresh() {
      // 将data恢复到初始状态
      this.form = this.$options.data().form
    },
    changeSourceIndex() {
      if (this.sourceConfig.indexOf('type') == -1) {
        return
      }
      this.mappingConfig = []
      const mappingConfig = []
      for (const indexName of this.form.source.index) {
        if (this.mappings.hasOwnProperty(indexName)) {
          const mappings = Object.keys(this.mappings[indexName].mappings)
          if (mappings.length > 0) {
            mappingConfig.push(mappings[0])
          }
        }
      }
      this.mappingConfig = mappingConfig
    },
    getParams(v) {
      console.log(v, typeof v)
      this.form.script.params = v
    },
    getBody(v) {
      this.form.source.query = v
    },
    changeSourceConfig() {
      console.log(111)
    },
    async getIndexList() {
      const input = {}
      input['es_connect'] = sdk.GetSelectEsConnID()
      const { data, code, msg } = await ListAction(input)
      if (code == 0) {
        this.indexNameList = Object.keys(data.list)
        this.mappings = data
      } else {
        ElMessage({
          type: 'error',
          message: msg
        })
        return
      }
    }
  }
}
</script>

<style scoped>
.width30 {
  width: 300px
}

.sort-key {
  width: 30%
}

.margin-20 {
  margin: 20px;
}

.margin-left30 {
  margin-left: 30px;
}

.card {
  width: 95%;
  float: left;

}
</style>
