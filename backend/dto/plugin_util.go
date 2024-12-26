package dto

import (
	"github.com/1340691923/eve-plugin-sdk-go/ev_api/proto"
	"mime/multipart"
	"net/http"
	"net/url"
)

type LoadDebugPlugin struct {
	ID   string `json:"id"`
	Addr string `json:"addr"`
	Pid  int    `json:"pid"`
}

type StopDebugPlugin struct {
	ID string `json:"id"`
}

type EsConnectData struct {
	UserID    int `json:"user_id"`
	EsConnect int `json:"es_connect"`
}

type PluginRunDsl struct {
	EsConnectData *EsConnectData `json:"es_connect_data"`
	Params        url.Values     `json:"-"`
	HttpMethod    string         `json:"http_method"`
	Path          string         `json:"path"`
	Dsl           string         `json:"dsl"`
}

type PluginRunDsl2 struct {
	Params     url.Values `json:"-"`
	HttpMethod string     `json:"http_method"`
	Path       string     `json:"path"`
	Dsl        string     `json:"dsl"`
}

type CatNodeReqData struct {
	H []string `json:"h"`
}

type CatNodesReq struct {
	EsConnectData  EsConnectData  `json:"es_connect_data"`
	CatNodeReqData CatNodeReqData `json:"cat_node_req_data"`
}

type ClusterStatsReqData struct {
	Human bool `json:"human"`
}

type ClusterStatsReq struct {
	EsConnectData       EsConnectData       `json:"es_connect_data"`
	ClusterStatsReqData ClusterStatsReqData `json:"cluster_stats_req_data"`
}

type PerformRequest struct {
	EsConnectData EsConnectData `json:"es_connect_data"`
	Request       *Request      `json:"request"`
}

type Request struct {
	Method        string
	URL           *url.URL
	Header        http.Header
	Form          url.Values
	PostForm      url.Values
	MultipartForm *multipart.Form
	JsonBody      string
}

type DslHistoryListReq struct {
	EsConnectData EsConnectData `json:"es_connect_data"`
	IndexName     string        `json:"indexName"` // 索引名
	Date          []string      `json:"date"`      //开始时间与结束时间（格式：”年-月-日 时:分:秒“ ）
	Page          int           `json:"page"`      //拉取数据当前页
	Limit         int           `json:"limit"`     //拉取条数
}

type IndicesSegmentsRequestData struct {
	Human bool `json:"human"`
}

type IndicesSegmentsRequest struct {
	EsConnectData              EsConnectData              `json:"es_connect_data"`
	IndicesSegmentsRequestData IndicesSegmentsRequestData `json:"indices_segments_request_data"`
}

type PingReq struct {
	EsConnectData EsConnectData `json:"es_connect_data"`
}

type RefreshReqData struct {
	IndexNames []string `json:"index_names"`
}

type RefreshReq struct {
	EsConnectData  EsConnectData  `json:"es_connect_data"`
	RefreshReqData RefreshReqData `json:"refresh_req_data"`
}

type OpenReqData struct {
	IndexNames []string `json:"index_names"`
}

type OpenReq struct {
	EsConnectData EsConnectData `json:"es_connect_data"`
	OpenReqData   OpenReqData   `json:"open_req_data"`
}

type FlushReqData struct {
	IndexNames []string `json:"index_names"`
}

type FlushReq struct {
	EsConnectData EsConnectData `json:"es_connect_data"`
	FlushReqData  FlushReqData  `json:"flush_req"`
}

type IndicesClearCacheReqData struct {
	IndexNames []string `json:"index_names"`
}

type IndicesClearCacheReq struct {
	EsConnectData            EsConnectData            `json:"es_connect_data"`
	IndicesClearCacheReqData IndicesClearCacheReqData `json:"indices_clear_cache_req_data"`
}

type IndicesCloseReqData struct {
	IndexNames []string `json:"index_names"`
}

type IndicesCloseReq struct {
	EsConnectData       EsConnectData       `json:"es_connect_data"`
	IndicesCloseReqData IndicesCloseReqData `json:"indices_close_req_data"`
}

type IndicesForcemergeReqData struct {
	IndexNames     []string `json:"index_names"`
	MaxNumSegments *int     `json:"max_num_segments"`
}

type IndicesForcemergeReq struct {
	EsConnectData            EsConnectData            `json:"es_connect_data"`
	IndicesForcemergeReqData IndicesForcemergeReqData `json:"indices_forcemerge_req_data"`
}

type DeleteByQueryReqData struct {
	IndexNames []string
	Documents  []string
	Body       interface{}
}

type DeleteByQueryReq struct {
	EsConnectData        EsConnectData        `json:"es_connect_data"`
	DeleteByQueryReqData DeleteByQueryReqData `json:"delete_by_query_req_data"`
}

type SnapshotCreateReqData struct {
	Repository        string
	Snapshot          string
	WaitForCompletion *bool
	ReqJson           proto.Json
}

type SnapshotCreateReq struct {
	EsConnectData         EsConnectData         `json:"es_connect_data"`
	SnapshotCreateReqData SnapshotCreateReqData `json:"snapshot_create_req_data"`
}

type SnapshotDeleteReqData struct {
	Repository string
	Snapshot   string
}

type SnapshotDeleteReq struct {
	EsConnectData         EsConnectData         `json:"es_connect_data"`
	SnapshotDeleteReqData SnapshotDeleteReqData `json:"snapshot_delete_req_data"`
}

type RestoreSnapshotReqData struct {
	Repository        string
	Snapshot          string
	WaitForCompletion *bool
	ReqJson           proto.Json
}

type RestoreSnapshotReq struct {
	EsConnectData          EsConnectData          `json:"es_connect_data"`
	RestoreSnapshotReqData RestoreSnapshotReqData `json:"restore_snapshot_req_data"`
}

type SnapshotStatusReqData struct {
	Repository        string
	Snapshot          []string
	IgnoreUnavailable *bool
}

type SnapshotStatusReq struct {
	EsConnectData         EsConnectData         `json:"es_connect_data"`
	SnapshotStatusReqData SnapshotStatusReqData `json:"snapshot_status_req_data"`
}
type SnapshotGetRepositoryReqData struct {
	Repository []string
}
type SnapshotGetRepositoryReq struct {
	EsConnectData                EsConnectData                `json:"es_connect_data"`
	SnapshotGetRepositoryReqData SnapshotGetRepositoryReqData `json:"snapshot_get_repository_req_data"`
}

type SnapshotCreateRepositoryReqData struct {
	Repository string
	ReqJson    proto.Json
}

type SnapshotCreateRepositoryReq struct {
	EsConnectData                   EsConnectData                   `json:"es_connect_data"`
	SnapshotCreateRepositoryReqData SnapshotCreateRepositoryReqData `json:"snapshot_create_repository_req_data"`
}

type SnapshotDeleteRepositoryReqData struct {
	Repository []string
}

type SnapshotDeleteRepositoryReq struct {
	EsConnectData                   EsConnectData                   `json:"es_connect_data"`
	SnapshotDeleteRepositoryReqData SnapshotDeleteRepositoryReqData `json:"snapshot_delete_repository_req_data"`
}
type GetIndicesReqData struct {
	CatIndicesRequest proto.CatIndicesRequest
}

type GetIndicesReq struct {
	EsConnectData     EsConnectData     `json:"es_connect_data"`
	GetIndicesReqData GetIndicesReqData `json:"get_indices_req_data"`
}

type CatHealthReqData struct {
	CatRequest proto.CatHealthRequest
}

type CatHealthReq struct {
	EsConnectData    EsConnectData    `json:"es_connect_data"`
	CatHealthReqData CatHealthReqData `json:"cat_health_req_data"`
}

type CatShardsReqData struct {
	CatRequest proto.CatShardsRequest
}

type CatShardsReq struct {
	EsConnectData    EsConnectData    `json:"es_connect_data"`
	CatShardsReqData CatShardsReqData `json:"cat_shards_req_data"`
}

type CatCountReqData struct {
	CatRequest proto.CatCountRequest
}

type CatCountReq struct {
	EsConnectData   EsConnectData   `json:"es_connect_data"`
	CatCountReqData CatCountReqData `json:"cat_count_req_data"`
}

type CatAllocationRequestData struct {
	CatRequest proto.CatAllocationRequest
}

type CatAllocationRequest struct {
	EsConnectData            EsConnectData            `json:"es_connect_data"`
	CatAllocationRequestData CatAllocationRequestData `json:"cat_allocation_request_data"`
}

type CatAliasesReqData struct {
	CatRequest proto.CatAliasesRequest
}

type CatAliasesReq struct {
	EsConnectData     EsConnectData     `json:"es_connect_data"`
	CatAliasesReqData CatAliasesReqData `json:"cat_aliases_req_data"`
}

type DeleteReqData struct {
	DeleteRequest proto.DeleteRequest
}

type DeleteReq struct {
	EsConnectData EsConnectData `json:"es_connect_data"`
	DeleteReqData DeleteReqData `json:"delete_req_data"`
}

type UpdateReqData struct {
	UpdateRequest proto.UpdateRequest
	Body          interface{}
}

type UpdateReq struct {
	EsConnectData EsConnectData `json:"es_connect_data"`
	UpdateReqData UpdateReqData `json:"update_req_data"`
}

type CreateReqData struct {
	CreateRequest proto.CreateRequest
	Body          interface{}
}

type CreateReq struct {
	EsConnectData EsConnectData `json:"es_connect_data"`
	CreateReqData CreateReqData `json:"create_req_data"`
}

type SearchReqData struct {
	SearchRequest proto.SearchRequest
	Query         interface{}
}

type SearchReq struct {
	EsConnectData EsConnectData `json:"es_connect_data"`
	SearchReqData SearchReqData `json:"search_req_data"`
}

type IndicesPutSettingsRequestData struct {
	IndexSettingsRequest proto.IndicesPutSettingsRequest
	Body                 interface{}
}

type IndicesPutSettingsRequest struct {
	EsConnectData                 EsConnectData                 `json:"es_connect_data"`
	IndicesPutSettingsRequestData IndicesPutSettingsRequestData `json:"indices_put_settings_request_data"`
}

type CreateIndexReqData struct {
	IndexCreateRequest proto.IndicesCreateRequest
	Body               interface{}
}

type CreateIndexReq struct {
	EsConnectData      EsConnectData      `json:"es_connect_data"`
	CreateIndexReqData CreateIndexReqData `json:"create_index_req_data"`
}

type DeleteIndexReqData struct {
	IndicesDeleteRequest proto.IndicesDeleteRequest
}

type DeleteIndexReq struct {
	EsConnectData      EsConnectData      `json:"es_connect_data"`
	DeleteIndexReqData DeleteIndexReqData `json:"delete_index_req_data"`
}

type ReindexReqData struct {
	ReindexRequest proto.ReindexRequest
	Body           interface{}
}

type ReindexReq struct {
	EsConnectData  EsConnectData  `json:"es_connect_data"`
	ReindexReqData ReindexReqData `json:"reindex_req_data"`
}

type IndicesGetSettingsRequestReqData struct {
	IndicesGetSettingsRequest proto.IndicesGetSettingsRequest
}

type IndicesGetSettingsRequestReq struct {
	EsConnectData                    EsConnectData                    `json:"es_connect_data"`
	IndicesGetSettingsRequestReqData IndicesGetSettingsRequestReqData `json:"indices_get_settings_request_req_data"`
}

type PutMappingReqData struct {
	IndicesPutMappingRequest proto.IndicesPutMappingRequest
	Body                     interface{}
}

type PutMappingReq struct {
	EsConnectData     EsConnectData     `json:"es_connect_data"`
	PutMappingReqData PutMappingReqData `json:"put_mapping_req_data"`
}

type GetMappingReqData struct {
	IndexNames []string
}

type GetMappingReq struct {
	EsConnectData     EsConnectData     `json:"es_connect_data"`
	GetMappingReqData GetMappingReqData `json:"get_mapping_req_data"`
}

type GetAliasesReqData struct {
	IndexNames []string
}

type GetAliasesReq struct {
	EsConnectData     EsConnectData     `json:"es_connect_data"`
	GetAliasesReqData GetAliasesReqData `json:"get_aliases_req_data"`
}

type AddAliasesReqData struct {
	IndexName []string
	AliasName string
}

type AddAliasesReq struct {
	EsConnectData     EsConnectData     `json:"es_connect_data"`
	AddAliasesReqData AddAliasesReqData `json:"add_aliases_req_data"`
}

type RemoveAliasesReqData struct {
	IndexName []string
	AliasName []string
}

type RemoveAliasesReq struct {
	EsConnectData        EsConnectData        `json:"es_connect_data"`
	RemoveAliasesReqData RemoveAliasesReqData `json:"remove_aliases_req_data"`
}

type MoveToAnotherIndexAliasesReqData struct {
	Body proto.AliasAction
}

type MoveToAnotherIndexAliasesReq struct {
	EsConnectData                    EsConnectData                    `json:"es_connect_data"`
	MoveToAnotherIndexAliasesReqData MoveToAnotherIndexAliasesReqData `json:"move_to_another_index_aliases_req_data"`
}

type TaskListReq struct {
	EsConnectData EsConnectData `json:"es_connect_data"`
}

type TasksCancelReqData struct {
	TaskId string
}

type TasksCancelReq struct {
	EsConnectData      EsConnectData      `json:"es_connect_data"`
	TasksCancelReqData TasksCancelReqData `json:"tasks_cancel_req_data"`
}
