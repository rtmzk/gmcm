export interface ILoginState {
  token: string
  username: string
}
export interface UserInfo {
  metadata: UserMeta
  uuid: string
  username: string
  password: string
  nickname: string
  headerImg: string
  authorityId: string
}

export interface UserMeta {
  id: number
  createAt: string
  updateAt: string
}
