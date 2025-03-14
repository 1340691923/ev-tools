import {request} from "@elasticview/plugin-sdk";

const api = '/api/es_index/'

export function CreateAction(data) {
  return request({
    url: api + 'CreateAction',
    method: 'post',
    data
  })
}

export function GetSettingsAction(data) {
  return request({
    url: api + 'GetSettingsAction',
    method: 'post',
    data
  })
}

export function IndexNamesAction(data) {
  return request({
    url: api + 'IndexNamesAction',
    method: 'post',
    data
  })
}

export function ReindexAction(data) {
  return request({
    url: api + 'ReindexAction',
    method: 'post',
    data
  })
}

export function GetAliasAction(data) {
  return request({
    url: api + 'GetAliasAction',
    method: 'post',
    data
  })
}

export function DeleteAction(data) {
  return request({
    url: api + 'DeleteAction',
    method: 'post',
    data
  })
}

export function MoveAliasToIndex(data) {
  return request({
    url: api + 'MoveAliasToIndex',
    method: 'post',
    data
  })
}

export function AddAliasToIndex(data) {
  return request({
    url: api + 'AddAliasToIndex',
    method: 'post',
    data
  })
}
export function BatchAddAliasToIndex(data) {
  return request({
    url: api + 'BatchAddAliasToIndex',
    method: 'post',
    data
  })
}

export function RemoveAlias(data) {
  return request({
    url: api + 'RemoveAlias',
    method: 'post',
    data
  })
}

export function GetSettingsInfoAction(data) {
  return request({
    url: api + 'GetSettingsInfoAction',
    method: 'post',
    data
  })
}

export function StatsAction(data) {
  return request({
    url: api + 'StatsAction',
    method: 'post',
    data
  })
}

export function CatStatusAction(data) {
  return request({
    url: api + 'CatStatusAction',
    method: 'post',
    data
  })
}

export function IndexsCountAction(data) {
  return request({
    url: api + 'IndexsCountAction',
    method: 'post',
    data
  })
}

export function Forcemerge(data) {
  return request({
    url: api + `Forcemerge`,
    method: 'post',
    data
  })
}

export function Empty(data) {
  return request({
    url: api + `Empty`,
    method: 'post',
    data
  })
}

export function Flush(data) {
  return request({
    url: api + `Flush`,
    method: 'post',
    data
  })
}

export function CacheClear(data) {
  return request({
    url: api + `CacheClear`,
    method: 'post',
    data
  })
}

export function Open(data) {
  return request({
    url: api + `Open`,
    method: 'post',
    data
  })
}

export function Close(data) {
  return request({
    url: api + `Close`,
    method: 'post',
    data
  })
}

export function Refresh(data) {
  return request({
    url: api + `Refresh`,
    method: 'post',
    data
  })
}


