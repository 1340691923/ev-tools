import {request} from "../plugin_sdk/sdk";

const api = '/api/es_doc/'

export function DeleteRowByIDAction(data) {
  return request({
    url: api + 'DeleteRowByIDAction',
    method: 'post',
    data
  })
}

export function UpdateByIDAction(data) {
  return request({
    url: api + 'UpdateByIDAction',
    method: 'post',
    data
  })
}

export function InsertAction(data) {
  return request({
    url: api + 'InsertAction',
    method: 'post',
    data
  })
}




