package es_doc_service

import (
	"context"
	"ev-plugin/backend/dto"
	"github.com/goccy/go-json"

	"github.com/1340691923/eve-plugin-sdk-go/ev_api/pkg"
	proto2 "github.com/1340691923/eve-plugin-sdk-go/ev_api/proto"
)

type EsDocService struct{}

func NewEsDocService() *EsDocService {
	return &EsDocService{}
}

func (this *EsDocService) DeleteRowByIDAction(ctx context.Context, esClient pkg.ClientInterface, esDocDeleteRowByID *dto.EsDocDeleteRowByID) (err error) {
	res, err := esClient.EsDelete(ctx, proto2.DeleteRequest{
		Index:        esDocDeleteRowByID.IndexName,
		DocumentType: esDocDeleteRowByID.Type,
		DocumentID:   esDocDeleteRowByID.ID,
	})
	if err != nil {
		return
	}
	if res.StatusErr() != nil {
		return res.StatusErr()
	}
	return
}

func (this *EsDocService) EsDocUpdateByID(ctx context.Context, esClient pkg.ClientInterface, esDocUpdateByID *dto.EsDocUpdateByID) (err error) {

	res, err := esClient.EsUpdate(ctx, proto2.UpdateRequest{
		Index:        esDocUpdateByID.Index,
		DocumentType: esDocUpdateByID.Type,
		DocumentID:   esDocUpdateByID.ID,
	}, esDocUpdateByID.JSON)
	if err != nil {
		return
	}
	if res.StatusErr() != nil {
		return res.StatusErr()
	}
	return
}

func (this *EsDocService) EsDocInsert(ctx context.Context, esClient pkg.ClientInterface, esDocUpdateByID *dto.EsDocUpdateByID) (res json.RawMessage, err error) {

	resp, err := esClient.EsCreate(ctx, proto2.CreateRequest{
		Index:        esDocUpdateByID.Index,
		DocumentType: esDocUpdateByID.Type,
		DocumentID:   esDocUpdateByID.ID,
	}, esDocUpdateByID.JSON)
	if err != nil {
		return nil, err
	}
	if resp.StatusErr() != nil {
		return nil, resp.StatusErr()
	}

	return resp.JsonRawMessage(), nil
}
