export interface Metadata {
  id: number
  createdAt: string
  updateAt: string
}

export interface Host {
  metadata: Metadata
  ip: string
  node_type: string
  node_role: string
  ssh_user: string
  ssh_port: string
  ssh_password: string
  login_type: string
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
  public_network: string
  cluster_network: string
  replicas: number
  devices: Devices[]
}

export interface Hosts {
  hosts: Host[]
  devs: StorageInspect
}
