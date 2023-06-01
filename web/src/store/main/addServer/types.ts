import { IServer } from '@/service/host/types'

export interface OpenStatus {
  openOr: boolean
  HostInfo: IServer
  rules: rule[]
  checkStatus: boolean
}

export interface rule {
  name: string
  description: string
  func: string
  status: string
  message: string
}
