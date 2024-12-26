package dto

import "github.com/1340691923/eve-plugin-sdk-go/ev_api/dto/common"

type EsAliasInfo struct {
	EsConnect        int         `json:"es_connect"`
	Settings         common.Json `json:"settings"`
	IndexName        string      `json:"index_name"`
	AliasName        string      `json:"alias_name"`
	NewAliasNameList []string    `json:"new_alias_name_list"`
	NewIndexList     []string    `json:"new_index_list"`
	Types            int         `json:"types"`
}
