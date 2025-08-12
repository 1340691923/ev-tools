# Elasticsearch 客户端 API 调用文档

## 索引管理 (Index Management)

### 创建索引
```go
res, err := esClient.EsCreateIndex(ctx, proto2.IndicesCreateRequest{
    Index: "index_name",
    Body:  indexMapping, // 索引映射配置
})
```

### 删除索引
```go
resp, err := esClient.EsDeleteIndex(ctx, proto2.IndicesDeleteRequest{
    Index: []string{"index_name"},
})
```

### 获取索引设置
```go
res, err := esClient.EsIndicesGetSettingsRequest(ctx, proto2.IndicesGetSettingsRequest{
    Index: []string{"index_name"},
})
```

### 更新索引设置
```go
res, err := esClient.EsIndicesPutSettingsRequest(ctx, proto2.IndicesPutSettingsRequest{
    Index: []string{"index_name"},
    Body:  settingsBody,
})
```

### 获取索引映射
```go
resp, err := esClient.EsGetMapping(ctx, []string{"index_name"})
```

### 更新索引映射
```go
resp, err := esClient.EsPutMapping(ctx, proto2.IndicesPutMappingRequest{
    Index: []string{"index_name"},
    Body:  mappingBody,
})
```

### 重建索引
```go
reindexRes, err := esClient.EsReindex(ctx, reindexRequest, reindexBody)
```

### 索引操作
```go
// 刷新索引
resp, err := esClient.EsRefresh(ctx, indexNames)

// 清空缓存
resp, err := esClient.EsIndicesClearCache(ctx, indexNames)

// 关闭索引
resp, err := esClient.EsIndicesClose(ctx, indexNames)

// 打开索引
resp, err := esClient.EsOpen(ctx, indexNames)

// 强制合并
resp, err := esClient.EsIndicesForcemerge(ctx, indexNames, &maxNumSegments)

// 刷新到磁盘
resp, err := esClient.EsFlush(ctx, indexNames)

// 按查询删除
resp, err := esClient.EsDeleteByQuery(ctx, indexNames, nil, deleteQuery)
```

## 别名管理 (Alias Management)

### 添加别名到索引
```go
resp, err := esClient.EsAddAliases(ctx, []string{"index_name"}, "alias_name")
```

### 批量添加别名
```go
resp, err := esClient.EsAddAliases(ctx, []string{"index_name"}, aliasName)
```

### 移除别名
```go
resp, err := esClient.EsRemoveAliases(
    ctx,
    []string{"index_name"},
    []string{"alias_name"},
)
```

### 获取索引别名
```go
aliasRes, err := esClient.EsGetAliases(ctx, []string{"index_name"})
```

### 移动别名到另一个索引
```go
_, err = esClient.EsMoveToAnotherIndexAliases(
    ctx,
    proto.AliasAction{Actions: []proto.AliasAddAction{
        {
            Add: proto.AliasAdd{
                Indices: newIndexList,
                Alias:   aliasName,
            },
        },
    }})
```

## 文档操作 (Document Operations)

### 删除文档
```go
res, err := esClient.EsDelete(ctx, proto2.DeleteRequest{
    Index:        "index_name",
    DocumentType: "doc_type",
    DocumentID:   "document_id",
})
```

### 更新文档
```go
res, err := esClient.EsUpdate(ctx, proto2.UpdateRequest{
    Index:        "index_name",
    DocumentType: "doc_type",
    DocumentID:   "document_id",
}, updateJSON)
```

### 创建文档
```go
resp, err := esClient.EsCreate(ctx, proto2.CreateRequest{
    Index:        "index_name",
    DocumentType: "doc_type",
    DocumentID:   "document_id",
}, documentJSON)
```

### 搜索文档
```go
resp, err := esClient.EsSearch(
    ctx,
    proto2.SearchRequest{Index: []string{"index_name"}},
    searchBody,
)
```

### 批量删除文档
```go
// 构建批量删除请求
var buf bytes.Buffer
for _, id := range ids {
    meta := fmt.Sprintf(`{ "delete" : { "_index" : "%s", "_id" : "%s" } }`, index, id)
    buf.WriteString(meta + "\n")
}

req, err := http.NewRequest(http.MethodPost, "/_bulk", strings.NewReader(buf.String()))
req.Header.Set("Content-Type", "application/x-ndjson")

res, err := esClient.EsPerformRequest(ctx, req)
```

## 集群状态监控 (Cluster Monitoring)

### 集群健康状态
```go
res, err = esClient.EsCatHealth(ctx, proto2.CatHealthRequest{
    Format: "json",
    Human:  true,
})
```

### 分片信息
```go
res, err = esClient.EsCatShards(ctx, proto2.CatShardsRequest{
    Format: "json",
    Human:  true,
})
```

### 文档数量统计
```go
res, err = esClient.EsCatCount(ctx, proto2.CatCountRequest{
    Format: "json",
    Human:  true,
})
```

### 节点分配信息
```go
res, err = esClient.EsCatAllocationRequest(ctx, proto2.CatAllocationRequest{
    Format: "json",
    Human:  true,
})
```

### 别名信息
```go
res, err = esClient.EsCatAliases(ctx, proto2.CatAliasesRequest{
    Format: "json",
    Human:  true,
})
```

### 节点信息
```go
res, err = esClient.EsCatNodes(ctx, strings.Split("ip,name,heap.percent,heap.current,heap.max,ram.percent,ram.current,ram.max,node.role,master,cpu,load_1m,load_5m,load_15m,disk.used_percent,disk.used,disk.total", ","))
```

### 索引信息
```go
catIndicesResponse, err := esClient.EsGetIndices(ctx, proto2.CatIndicesRequest{
    Format: "json",
    Human:  true,
    Index:  []string{"*,-.*"}, // 排除系统索引
})
```

### 分片段信息
```go
res, err = esClient.EsIndicesSegmentsRequest(ctx, true)
```

### 集群统计
```go
res, err = esClient.EsClusterStats(ctx, true)
```

## 快照管理 (Snapshot Management)

### 获取快照仓库
```go
res, err := esClient.EsSnapshotGetRepository(ctx, repositoryNames)
```

### 创建快照仓库
```go
res, err := esClient.EsSnapshotCreateRepository(ctx, repositoryName, repositoryBody)
```

### 删除快照仓库
```go
res, err := esClient.EsSnapshotDeleteRepository(ctx, []string{repositoryName})
```

### 创建快照
```go
esRes, err := esClient.EsSnapshotCreate(
    ctx,
    repositoryName,
    snapshotName,
    snapshotBody,
)
```

### 获取快照状态
```go
esRes, err := esClient.EsSnapshotStatus(
    ctx,
    repositoryName,
    snapshotName,
)
```

### 删除快照
```go
esRes, err := esClient.EsSnapshotDelete(ctx, repositoryName, snapshotName)
```

### 恢复快照
```go
esRes, err := esClient.EsRestoreSnapshot(
    ctx,
    repositoryName,
    snapshotName,
    restoreBody,
)
```

## 任务管理 (Task Management)

### 获取任务列表
```go
resp, err := esClient.EsTaskList(ctx)
```

### 取消任务
```go
res, err := esClient.EsTasksCancel(ctx, taskID)
```

## 通用请求 (Generic Requests)

### 执行自定义请求
```go
res, err := esClient.EsPerformRequest(ctx, httpRequest)
```

### 获取 ES 版本
```go
version, err := esClient.EsVersion()
```

## 错误处理

所有 API 调用都应该检查错误：

```go
if err != nil {
    return err
}

if res.StatusErr() != nil {
    return res.StatusErr()
}
```

## 注意事项

1. 所有请求都需要传入 `context.Context`
2. 文档类型 (DocumentType) 在 ES 7.0+ 中已废弃，但仍需传入
3. 批量操作时注意请求体格式
4. 监控 API 支持 `Format: "json"` 和 `Human: true` 参数
5. 索引操作支持批量处理多个索引名 