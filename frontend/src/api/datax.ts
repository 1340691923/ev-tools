import {request} from "@elasticview/plugin-sdk";

let api = "/api/datax/"

export function LinkInfoList(data) {

  return request({
    url: api + 'LinkInfoList',
    method: 'post',
    data
  })
}

export function InsertLink(data) {
  return request({
    url: api + 'InsertLink',
    method: 'post',
    data
  })
}

export function DelLinkById(data) {
  return request({
    url: api + 'DelLinkById',
    method: 'post',
    data
  })
}

export function TestLink(data) {
  return request({
    url: api + 'TestLink',
    method: 'post',
    data
  })
}

export function LinkSelectOpt(data) {
  return request({
    url: api + 'LinkSelectOpt',
    method: 'post',
    data
  })
}

export function GetTables(data) {
  return request({
    url: api + 'Tables',
    method: 'post',
    data
  })
}

export function GetTableColumns(data) {
  return request({
    url: api + 'GetTableColumns',
    method: 'post',
    data
  })
}

export function Transfer(data) {
  return request({
    url: api + 'Transfer',
    method: 'post',
    data
  })
}

export function TransferLogList(data) {
  return request({
    url: api + 'TransferLogList',
    method: 'post',
    data
  })
}

export function CancelTaskById(data) {
  return request({
    url: api + 'CancelTaskById',
    method: 'post',
    data
  })
}

export function DeleteTaskById(data) {
  return request({
    url: api + 'DeleteTaskById',
    method: 'post',
    data
  })
}
