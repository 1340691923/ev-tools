package cat_service

import (
	"context"
	"github.com/1340691923/eve-plugin-sdk-go/ev_api/pkg"
	proto2 "github.com/1340691923/eve-plugin-sdk-go/ev_api/proto"
	"strings"
)

type CatService struct{}

func NewCatService() *CatService {
	return &CatService{}
}

func (this *CatService) CatHealth(ctx context.Context, esClient pkg.ClientInterface) (res *proto2.Response, err error) {
	res, err = esClient.EsCatHealth(ctx, proto2.CatHealthRequest{
		Format: "json",
		Human:  true,
	})
	return
}

func (this *CatService) CatShards(ctx context.Context, esClient pkg.ClientInterface) (res *proto2.Response, err error) {
	res, err = esClient.EsCatShards(ctx, proto2.CatShardsRequest{
		Format: "json",
		Human:  true,
	})
	return
}

func (this *CatService) CatCount(ctx context.Context, esClient pkg.ClientInterface) (res *proto2.Response, err error) {
	res, err = esClient.EsCatCount(ctx, proto2.CatCountRequest{
		Format: "json",
		Human:  true,
	})
	return
}

func (this *CatService) CatAllocation(ctx context.Context, esClient pkg.ClientInterface) (res *proto2.Response, err error) {
	res, err = esClient.EsCatAllocationRequest(ctx, proto2.CatAllocationRequest{
		Format: "json",
		Human:  true,
	})
	return
}

func (this *CatService) CatAliases(ctx context.Context, esClient pkg.ClientInterface) (res *proto2.Response, err error) {
	res, err = esClient.EsCatAliases(ctx, proto2.CatAliasesRequest{
		Format: "json",
		Human:  true,
	})
	return
}

func (this *CatService) CatNodes(ctx context.Context, esClient pkg.ClientInterface) (res *proto2.Response, err error) {
	res, err = esClient.EsCatNodes(ctx, strings.Split("ip,name,heap.percent,heap.current,heap.max,ram.percent,ram.current,ram.max,node.role,master,cpu,load_1m,load_5m,load_15m,disk.used_percent,disk.used,disk.total", ","))
	return
}

func (this *CatService) CatIndices(ctx context.Context, esClient pkg.ClientInterface,
	sort []string, indexBytesFormat string, isBanSysIndex bool) (res *proto2.Response, err error) {

	req := proto2.CatIndicesRequest{}
	req.S = sort
	req.Human = true
	req.Format = "json"
	if indexBytesFormat != "" {
		req.Bytes = indexBytesFormat
	}
	if isBanSysIndex {
		req.Index = append(req.Index, "*,-.*")
	}

	res, err = esClient.EsGetIndices(ctx, req)
	return
}

func (this *CatService) IndicesSegmentsRequest(ctx context.Context, esClient pkg.ClientInterface) (res *proto2.Response, err error) {
	res, err = esClient.EsIndicesSegmentsRequest(ctx, true)
	return
}

func (this *CatService) ClusterStats(ctx context.Context, esClient pkg.ClientInterface) (res *proto2.Response, err error) {
	res, err = esClient.EsClusterStats(ctx, true)
	return
}
