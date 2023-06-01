export interface IHost {
  ip: string
  node_type: string
  node_role: any
  ssh_user: string
  ssh_port: string
  is_no_pass?: boolean
}
export interface IServer {
  ip: any
  node_role: any
  ssh_user: string
  ssh_password?: string
  ssh_port: string
  is_no_pass?: boolean
  public_network: string
  cluster_network?: string
}

export interface Device {
  name: string
  size: string
  type: string
  enabled: boolean
  cached: boolean
}

export interface Devices {
  ip: string
  device: Device[]
}

export interface StorageInspect {
  replicas: string
  devs: Devices[]
}

export interface rule {
  name: string
  description: string
  func: string
  status: string
  message: string
}

export interface rules {
  rules: rule[]
  code?: number
  message?: string
}

export interface commonResponse {
  code?: number
  message?: string
  successes?: string
}

export interface Status {
  status: number
}
