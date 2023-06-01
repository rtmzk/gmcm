import { Module } from 'vuex'
import { IRootState } from '../../types'
import { OpenStatus, rule } from './types'
import { IServer } from '@/service/host/types'
import LocalCache from '@/utils/cache'
import {
  hostAddRequest,
  hostEnvcActionRequest,
  hostEnvcRequest
} from '@/service/host/hosts'

const OpenModule: Module<OpenStatus, IRootState> = {
  namespaced: true,
  state() {
    return {
      openOr: false,
      HostInfo: {
        ip: [''],
        node_role: '',
        ssh_user: '',
        ssh_port: '',
        ssh_password: '',
        is_no_pass: false,
        public_network: '',
        cluster_network: ''
      },
      rules: [
        {
          name: '',
          description: '',
          func: '',
          status: '',
          message: ''
        }
      ],
      checkStatus: false
    }
  },
  mutations: {
    changeOpenStatus(state, value: boolean) {
      state.openOr = value
    },
    changeHostInfo(state, value: IServer) {
      state.HostInfo = value
    },
    changeRuleInfo(state, value: rule[]) {
      state.rules = value
    },
    changeCheckStatus(state, value: boolean) {
      state.checkStatus = value
    }
  },
  actions: {
    changeOpenStatus({ commit }, payload: any) {
      commit('changeOpenStatus', payload)
    },
    addHostRequest({ commit }, payload: any) {
      hostAddRequest(payload).then((res) => {
        if (res.successes) {
          return
        }
      })
    },
    async getCheckRules({ commit }) {
      const envRuleResult = await hostEnvcRequest()
      if (envRuleResult.rules) {
        commit('changeRuleInfo', envRuleResult.rules)
        localStorage.setItem('rules', JSON.stringify(envRuleResult.rules))
        return
      }
    },

    hostEnvc({ commit }, payload) {
      let failedCount = 0
      hostEnvcActionRequest(payload).then((res) => {
        if (res.rules) {
          commit('changeRuleInfo', res.rules)
          res.rules.forEach((e) => {
            if (e.status == 'FAILED') {
              failedCount = failedCount + 1
            }
          })
          if (failedCount == 3) {
            commit('changeCheckStatus', true)
          }
        }
      })
    },

    loadLocalCheckRule({ commit }) {
      const rules = LocalCache.getCache('rules')
      if (rules) {
        commit('changeRuleInfo', rules)
      }
    }
  },
  getters: {
    getOpenStatus(state) {
      return state.openOr
    },
    getHostInfo(state) {
      return state.HostInfo
    },
    getRuleList(state) {
      return state.rules
    },
    getCheckStatus(state) {
      return state.checkStatus
    }
  }
}

export default OpenModule
