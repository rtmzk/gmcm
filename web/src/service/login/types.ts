export interface IAccount {
  username: string
  password: string
}

export interface IToken {
  expire?: string
  token?: string
  message?: string
}

export interface IUser<T = any> {
  user: T
}
