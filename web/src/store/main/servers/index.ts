import { Module } from 'vuex'
import { IRootState } from '../../types'
import { Hosts } from './types'
// import { Devices } from './types'
import { StorageInspect } from './types'
import {
  hostListRequest,
  deviceListRequest,
  hostInitRequest
} from '@/service/host/hosts'

const ServerModule: Module<Hosts, IRootState> = {
  namespaced: true,
  state() {
    return {
      hosts: [
        {
          metadata: { id: 0, createdAt: '', updateAt: '' },
          ip: '',
          node_type: '',
          node_role: '',
          ssh_user: '',
          ssh_port: '',
          ssh_password: '',
          login_type: ''
        }
      ],
      devs: {
        public_network: '',
        cluster_network: '',
        replicas: 1,
        devices: [
          {
            ip: '',
            device: [
              {
                name: '',
                size: '',
                type: '',
                enabled: true,
                cached: false
              }
            ]
          }
        ]
      }
    }
  },
  mutations: {
    changeHosts(state, value: any) {
      state.hosts = value
    },
    changeDevice(state, value: StorageInspect) {
      state.devs = value
    }
  },
  actions: {
    serverList({ commit }) {
      hostListRequest().then((res) => {
        if (res.hosts) {
          commit('changeHosts', res.hosts)
          return
        }
      })
    },
    async deviceList({ commit }) {
      deviceListRequest().then((res) => {
        commit('changeDevice', res)
      })
    },
    initHostStorage({ commit }, payload: any) {
      hostInitRequest(payload).then((res) => {
        if (res.successes == 'ok') {
          return
        }
      })
    }
  },
  getters: {
    getServers(state) {
      return state.hosts
    },
    getDevices(state) {
      return state.devs
    }
  }
}

export default ServerModule
