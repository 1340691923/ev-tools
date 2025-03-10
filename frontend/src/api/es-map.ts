import {request} from "@elasticview/plugin-sdk";

const api = '/api/es_map/'

export function ListAction(data) {
  return request({
    url: api + 'ListAction',
    method: 'post',
    data
  })
}

export function UpdateMappingAction(data) {
  return request({
    url: api + 'UpdateMappingAction',
    method: 'post',
    data
  })
}

