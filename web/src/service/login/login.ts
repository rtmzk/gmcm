import defaultRequest from '../index'
import DefaultRequest from '../request'
import { BASE_URL, TIME_OUT } from '../request/config'
import { IAccount, IToken } from './types'

const loginRequest = new DefaultRequest({
  baseURL: BASE_URL,
  timeout: TIME_OUT
})

enum LoginApi {
  AccountLogin = '/login',
  AccountLogout = '/logout',
  AccountUserInfo = '/users/current'
}

export function accountLoginRequest(account: IAccount) {
  return loginRequest.post<IToken>({
    url: LoginApi.AccountLogin,
    data: account
  })
}

export function accountUserInfoByToken() {
  return defaultRequest.get({
    url: LoginApi.AccountUserInfo
  })
}

export function accountLogoutRequest() {
  return defaultRequest.post({
    url: LoginApi.AccountLogout
  })
}
