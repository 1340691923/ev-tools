import {request} from "@elasticview/plugin-sdk";

import jsonlint from  "jsonlint-mod"

const api = '/api/es/'

export function PingAction(data) {
  return request({
    url: api + `PingAction`,
    method: 'post',
    data
  })
}

export function CatAction(data) {
  return request({
    url: api + `CatAction`,
    method: 'post',
    data
  })
}

export function RunDslAction(data) {

  let action = "RunDsl"+data["method"]+"Action"

  return request({
    url: api + action,
    method: 'post',
    transformResponse : [
      data => {
        return jsonlint.parse(data)
      }
    ],
    data
  })
}

export function SqlToDslAction(data) {
  return request({
    url: api + `SqlToDslAction`,
    method: 'post',
     data
  })
}

export function OptimizeAction(data) {
  return request({
    url: api + `OptimizeAction`,
    method: 'post',
    transformResponse : [
      data => {
        return jsonlint.parse(data)
      }
    ],
    data
  })
}


export function RecoverCanWrite(data) {
  return request({
    url: api + `RecoverCanWrite`,
    method: 'post',
    data
  })
}
