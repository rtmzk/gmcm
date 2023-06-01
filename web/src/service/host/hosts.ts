import defaultRequest from '../index'
import { StorageInspect, IHost, rules, commonResponse, Status } from './types'

enum hostApi {
  hostList = '/host/list',
  connectionCheck = '/host/check',
  hostAdd = '/host/add',
  hostDelete = '/host/delete/',
  deviceList = '/host/devices',
  initAPI = '/host/init',
  logAPI = '/host/init/log/',
  envcRules = '/host/envc/prepare',
  envcAction = '/host/envc',
  getInitStatus = '/host/init/status',
  setInitStatus = '/host/init/status/update'
}

export function connectionCheckRequest(data: IHost) {
  return defaultRequest.post<commonResponse>({
    url: hostApi.connectionCheck,
    data: data
  })
}

export function hostAddRequest(data: IHost) {
  return defaultRequest.post<any>({
    url: hostApi.hostAdd,
    data: data
  })
}

export function hostListRequest() {
  return defaultRequest.get({
    url: hostApi.hostList
  })
}

export function hostDeleteRequest(id: number) {
  return defaultRequest.delete<any>({
    url: hostApi.hostDelete + id
  })
}

export function deviceListRequest() {
  return defaultRequest.get({
    url: hostApi.deviceList
  })
}

export function hostInitRequest(data: StorageInspect) {
  return defaultRequest.post<any>({
    url: hostApi.initAPI,
    data: data
  })
}

export function hostInitLogFetchRequest(offset: number) {
  return defaultRequest.get({
    url: hostApi.logAPI + offset
  })
}

export function hostEnvcRequest() {
  return defaultRequest.get<rules>({
    url: hostApi.envcRules
  })
}

export function hostEnvcActionRequest(data: IHost) {
  return defaultRequest.post<rules>({
    url: hostApi.envcAction,
    data: data
  })
}

export function getInitStatus() {
  return defaultRequest.get({
    url: hostApi.getInitStatus
  })
}

export function setInitStatus(data: Status) {
  return defaultRequest.post<commonResponse>({
    url: hostApi.setInitStatus,
    data: data
  })
}
