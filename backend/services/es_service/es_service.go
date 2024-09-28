package es_service

import (
	"bytes"
	"context"
	"github.com/1340691923/eve-plugin-sdk-go/ev_api/pkg"
	"github.com/goccy/go-json"
	"net/http"
)

type EsService struct {
}

func NewEsService() *EsService {
	return &EsService{}
}

func (this *EsService) RecoverCanWrite(ctx context.Context, esClient pkg.ClientInterface) (err error) {

	body := map[string]interface{}{
		"index": map[string]interface{}{
			"blocks": map[string]interface{}{
				"read_only_allow_delete": "false",
			},
		},
	}

	js, _ := json.Marshal(body)

	req, err := http.NewRequest(http.MethodPut, "/_settings", bytes.NewBuffer(js))

	if err != nil {
		return
	}

	res, err := esClient.EsPerformRequest(ctx, req)

	if err != nil {
		return
	}
	if res.StatusErr() != nil {
		err = res.StatusErr()
		return
	}

	return
}
