package es_service

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/1340691923/eve-plugin-sdk-go/ev_api/pkg"
	"net/http"
)

type EsService struct {
}

func NewEsService() *EsService {
	return &EsService{}
}

func (this *EsService) RecoverCanWrite(ctx context.Context, esClient pkg.EsI) (err error) {

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

	res, err := esClient.PerformRequest(ctx, req)

	if err != nil {
		return
	}
	if res.StatusErr() != nil {
		err = res.StatusErr()
		return
	}

	return
}
