package index_service

import (
	"context"
	"github.com/1340691923/eve-plugin-sdk-go/ev_api/dto"
	"github.com/1340691923/eve-plugin-sdk-go/ev_api/vo"
	"github.com/goccy/go-json"
	"github.com/pkg/errors"

	"github.com/1340691923/eve-plugin-sdk-go/ev_api/pkg"
	proto2 "github.com/1340691923/eve-plugin-sdk-go/ev_api/proto"
	"github.com/tidwall/gjson"
	"time"
)

type IndexService struct{}

func NewIndexService() *IndexService {
	return &IndexService{}
}

func (this *IndexService) EsIndexCreate(ctx context.Context, esClient pkg.ClientInterface, esIndexInfo *dto.EsIndexInfo) (err error) {
	if esIndexInfo.IndexName == "" {
		return errors.New("索引名不能为空")
	}
	if esIndexInfo.Types == "update" {
		res, err := esClient.EsIndicesPutSettingsRequest(ctx,
			proto2.IndicesPutSettingsRequest{
				Index: []string{esIndexInfo.IndexName},
			}, esIndexInfo.Settings,
		)
		if err != nil {
			return err
		}
		if res.StatusErr() != nil {
			err = res.StatusErr()
			return err
		}
	} else {
		res, err := esClient.EsCreateIndex(ctx, proto2.IndicesCreateRequest{
			Index: esIndexInfo.IndexName,
		},
			map[string]interface{}{
				"settings": esIndexInfo.Settings,
			})
		if err != nil {
			return err
		}
		if res.StatusErr() != nil {
			err = res.StatusErr()
			return err
		}

	}
	return nil
}

func (this *IndexService) EsIndexDelete(ctx context.Context, esClient pkg.ClientInterface, esIndexInfo *dto.EsIndexInfo) (err error) {
	if esIndexInfo.IndexName == "" {
		return errors.New("索引名不能为空")
	}
	resp, err := esClient.EsDeleteIndex(ctx, proto2.IndicesDeleteRequest{
		Index: []string{esIndexInfo.IndexName},
	})
	if err != nil {
		return
	}
	if resp.StatusErr() != nil {
		err = resp.StatusErr()
		return err
	}
	return
}

func (this *IndexService) EsIndexGetSettings(ctx context.Context, esClient pkg.ClientInterface, esIndexInfo *dto.EsIndexInfo) (settings map[string]interface{}, err error) {
	if esIndexInfo.IndexName == "" {
		err = errors.New("索引名不能为空")
		return
	}

	res, err := esClient.EsIndicesGetSettingsRequest(ctx, proto2.IndicesGetSettingsRequest{
		Index: []string{esIndexInfo.IndexName},
	})
	if err != nil {
		return
	}

	if res.StatusErr() != nil {
		err = res.StatusErr()
		return
	}

	settings = map[string]interface{}{}
	gjson.GetBytes(res.ResByte(), esIndexInfo.IndexName).Get("settings").ForEach(func(key, value gjson.Result) bool {
		settings[key.String()] = value.Value()
		return true
	})

	return
}

func (this *IndexService) EsIndexGetSettingsInfo(ctx context.Context, esClient pkg.ClientInterface, esIndexInfo *dto.EsIndexInfo) (settings map[string]interface{}, err error) {
	if esIndexInfo.IndexName == "" {
		err = errors.New("索引名不能为空")
		return
	}

	res, err := esClient.EsIndicesGetSettingsRequest(ctx, proto2.IndicesGetSettingsRequest{
		Index: []string{esIndexInfo.IndexName},
	})
	if err != nil {
		return
	}
	if res.StatusErr() != nil {
		err = res.StatusErr()
		return
	}

	settings = map[string]interface{}{}
	err = json.Unmarshal(res.ResByte(), &settings)
	if err != nil {
		return
	}
	return
}

func (this *IndexService) EsIndexReindex(ctx context.Context, esClient pkg.ClientInterface, esReIndexInfo *dto.EsReIndexInfo) (res map[string]interface{}, err error) {

	reindexRequest := proto2.ReindexRequest{}

	urlValues := esReIndexInfo.UrlValues
	if urlValues.WaitForActiveShards != "" {
		reindexRequest.WaitForActiveShards = urlValues.WaitForActiveShards
	}
	if urlValues.Slices != 0 {
		reindexRequest.Slices = urlValues.Slices
	}
	if urlValues.Refresh != nil {
		reindexRequest.Refresh = urlValues.Refresh
	}
	if urlValues.Timeout != 0 {
		reindexRequest.Timeout = time.Duration(int64(urlValues.Timeout)) * time.Second
	}
	if urlValues.RequestsPerSecond != 0 {
		requestsPerSecond := urlValues.RequestsPerSecond
		reindexRequest.RequestsPerSecond = &requestsPerSecond
	}
	if urlValues.WaitForCompletion != nil {
		reindexRequest.WaitForCompletion = urlValues.WaitForCompletion
	}

	reindexRes, err := esClient.EsReindex(ctx, reindexRequest, esReIndexInfo.Body)
	if err != nil {
		return
	}
	if reindexRes.StatusErr() != nil {
		err = reindexRes.StatusErr()
		return
	}

	res = map[string]interface{}{}
	err = json.Unmarshal(reindexRes.ResByte(), &res)
	if err != nil {
		return
	}
	return
}

func (this *IndexService) EsIndexNames(ctx context.Context, esClient pkg.ClientInterface) (indexNames []string, err error) {
	catIndicesResponse, err := esClient.EsGetIndices(ctx, proto2.CatIndicesRequest{
		Format: "json",
	})
	if err != nil {
		return
	}
	if catIndicesResponse.StatusErr() != nil {
		err = catIndicesResponse.StatusErr()
		return
	}
	var list []proto2.CatIndex
	err = json.Unmarshal(catIndicesResponse.ResByte(), &list)
	if err != nil {
		return
	}
	for _, v := range list {
		indexNames = append(indexNames, v.Index)
	}
	return
}

func (this *IndexService) EsIndexCount(ctx context.Context, esClient pkg.ClientInterface) (indexNameLen int, err error) {
	catIndicesResponse, err := esClient.EsGetIndices(ctx, proto2.CatIndicesRequest{
		Format: "json",
	})
	if err != nil {
		return
	}
	if catIndicesResponse.StatusErr() != nil {
		err = catIndicesResponse.StatusErr()
		return
	}
	var list []proto2.CatIndex
	err = json.Unmarshal(catIndicesResponse.ResByte(), &list)
	if err != nil {
		return
	}
	indexNameLen = len(list)
	return
}

func (this *IndexService) EsIndexStats(ctx context.Context, esClient pkg.ClientInterface, indexName string) (res []vo.Status, err error) {
	catIndicesResponse, err := esClient.EsGetIndices(ctx, proto2.CatIndicesRequest{
		Index:  []string{indexName},
		H:      []string{"status"},
		Format: "json",
	})
	if err != nil {
		return
	}
	if catIndicesResponse.StatusErr() != nil {
		err = catIndicesResponse.StatusErr()
		return
	}
	err = json.Unmarshal(catIndicesResponse.ResByte(), &res)
	return
}

func (this *IndexService) EsIndexCatStatus(ctx context.Context, esClient pkg.ClientInterface, indexName string) (res []vo.Status, err error) {
	catIndicesResponse, err := esClient.EsGetIndices(ctx, proto2.CatIndicesRequest{
		Index:  []string{indexName},
		H:      []string{"status"},
		Format: "json",
	})
	if err != nil {
		return
	}
	if catIndicesResponse.StatusErr() != nil {
		err = catIndicesResponse.StatusErr()
		return
	}

	err = json.Unmarshal(catIndicesResponse.ResByte(), &res)
	return
}

func (this *IndexService) Refresh(ctx context.Context, esClient pkg.ClientInterface, indexNames []string) (err error) {
	if len(indexNames) == 0 {
		indexNames = []string{"*"}
	}
	resp, err := esClient.EsRefresh(ctx, indexNames)
	if err != nil {
		return
	}
	if resp.StatusErr() != nil {
		err = resp.StatusErr()
		return
	}
	return
}

func (this *IndexService) CacheClear(ctx context.Context, esClient pkg.ClientInterface, indexNames []string) (err error) {
	if len(indexNames) == 0 {
		indexNames = []string{"*"}
	}
	resp, err := esClient.EsIndicesClearCache(ctx, indexNames)
	if err != nil {
		return
	}
	if resp.StatusErr() != nil {
		err = resp.StatusErr()
		return
	}
	return
}

func (this *IndexService) Close(ctx context.Context, esClient pkg.ClientInterface, indexNames []string) (err error) {
	if len(indexNames) == 0 {
		indexNames = []string{"*"}
	}
	resp, err := esClient.EsIndicesClose(ctx, indexNames)
	if err != nil {
		return
	}
	if resp.StatusErr() != nil {
		err = resp.StatusErr()
		return
	}
	return
}

func (this *IndexService) Empty(ctx context.Context, esClient pkg.ClientInterface, indexNames []string) (err error) {
	if len(indexNames) == 0 {
		indexNames = []string{"*"}
	}
	resp, err := esClient.EsDeleteByQuery(ctx, indexNames, nil, map[string]interface{}{
		"query": map[string]interface{}{
			"match_all": map[string]interface{}{},
		},
	})
	if err != nil {
		return
	}
	if resp.StatusErr() != nil {
		err = resp.StatusErr()
		return
	}
	return
}

func (this *IndexService) Flush(ctx context.Context, esClient pkg.ClientInterface, indexNames []string) (err error) {
	if len(indexNames) == 0 {
		indexNames = []string{"*"}
	}
	resp, err := esClient.EsFlush(ctx, indexNames)
	if err != nil {
		return
	}
	if resp.StatusErr() != nil {
		err = resp.StatusErr()
		return
	}
	return
}

func (this *IndexService) IndicesForcemerge(ctx context.Context, esClient pkg.ClientInterface, indexNames []string) (err error) {
	if len(indexNames) == 0 {
		indexNames = []string{"*"}
	}
	maxNumSegments := 1
	resp, err := esClient.EsIndicesForcemerge(ctx, indexNames, &maxNumSegments)
	if err != nil {
		return
	}
	if resp.StatusErr() != nil {
		err = resp.StatusErr()
		return
	}
	return
}

func (this *IndexService) Open(ctx context.Context, esClient pkg.ClientInterface, indexNames []string) (err error) {
	if len(indexNames) == 0 {
		indexNames = []string{"*"}
	}
	resp, err := esClient.EsOpen(ctx, indexNames)
	if err != nil {
		return
	}
	if resp.StatusErr() != nil {
		err = resp.StatusErr()
		return
	}
	return
}

func (this *IndexService) UpdateMapping(ctx context.Context, esClient pkg.ClientInterface, updateMapping *dto.UpdateMapping) (res json.RawMessage, err error) {
	resp, err := esClient.EsPutMapping(ctx, proto2.IndicesPutMappingRequest{
		Index:        []string{updateMapping.IndexName},
		DocumentType: updateMapping.TypeName,
	}, updateMapping.Properties)
	if err != nil {
		return
	}
	if resp.StatusErr() != nil {
		err = resp.StatusErr()
		return
	}
	res = resp.JsonRawMessage()

	return
}

func (this *IndexService) EsMappingList(ctx context.Context, esClient pkg.ClientInterface, esConnect *dto.EsMapGetProperties) (res json.RawMessage, err error) {
	indexNames := []string{}
	if esConnect.IndexName != "" {
		indexNames = []string{esConnect.IndexName}
	}
	resp, err := esClient.EsGetMapping(ctx, indexNames)
	if err != nil {
		return
	}
	if resp.StatusErr() != nil {
		err = resp.StatusErr()
		return
	}

	res = resp.JsonRawMessage()

	return
}
