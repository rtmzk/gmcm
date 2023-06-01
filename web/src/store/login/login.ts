import { Module } from 'vuex'
import { IRootState } from '../types'
import { ILoginState } from './types'
import {
  accountLoginRequest,
  accountUserInfoByToken
} from '@/service/login/login'

import localCache from '@/utils/cache'
import router from '@/router'

const loginModule: Module<ILoginState, IRootState> = {
  namespaced: true,
  state() {
    return {
      token: '',
      username: ''
    }
  },
  mutations: {
    changeToken(state, token: string) {
      state.token = token
    },
    changeUserInfo(state, username: string) {
      state.username = username
    }
  },
  actions: {
    async accountLoginAction({ commit }, payload: any) {
      // 获取token,缓存到浏览器本地
      const loginResult = await accountLoginRequest(payload)
      const token = loginResult.token
      if (!token) {
        return
      }
      commit('changeToken', token)
      localCache.setCache('token', token)

      // 获取用户信息
      const userResult = await accountUserInfoByToken()
      const username = userResult.username
      commit('changeUserInfo', username)
      localCache.setCache('username', username)
      localCache.setCache('userinfo', userResult)
      router.push('/main/index')
      // 获取规则列表
      this.dispatch('addServer/getCheckRules')
    },

    // vuex: 重新获取用户相关信息
    loadLocalLogin({ commit }) {
      const token = localCache.getCache('token')
      if (token) {
        commit('changeToken', token)
      }

      const user = localCache.getCache('username')
      if (user) {
        commit('changeUserInfo', user)
      }
    }
  },
  getters: {}
}

export default loginModule
