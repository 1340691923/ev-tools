import {request} from "../plugin_sdk/sdk";

const api = '/api/es_task/'

export function ListAction(data) {
  return request({
    url: api + 'ListAction',
    method: 'post',
    data
  })
}

export function CancelAction(data) {
  return request({
    url: api + 'CancelAction',
    method: 'post',
    data
  })
}


