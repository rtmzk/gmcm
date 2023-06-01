import { Module } from 'vuex'
import { IRootState } from '../../types'
import { storageOpenStatus, Status } from './types'
import { getInitStatus, setInitStatus } from '@/service/host/hosts'
// import LocalCache from '@/utils/cache'

const initStorageModule: Module<storageOpenStatus, IRootState> = {
  namespaced: true,
  state() {
    return {
      openOr: false,
      openLogPannel: false,
      btnStatus: false,
      initStatus: {
        status: 0
      }
    }
  },
  mutations: {
    changeOpenStatus(state, value: boolean) {
      state.openOr = value
    },
    changeOpenLogPannelStatus(state, value: boolean) {
      state.openLogPannel = value
    },
    changeBtnStatus(state, value: boolean) {
      state.btnStatus = value
    }
  },
  actions: {
    changeOpenStatus({ commit }, payload: any) {
      commit('changeOpenStatus', payload)
    },
    changeOpenLogPannelStatus({ commit }, payload) {
      commit('changeOpenLogPannelStatus', payload)
    },
    setInitStatusReloadBtn({ commit }, payload: Status) {
      setInitStatus(payload).then((res) => {
        if (res.successes) {
          commit('changeBtnStatus', true)
        }
      })
    },
    loadBtnStatus({ commit }) {
      getInitStatus().then((res) => {
        if (res.status == 0) {
          commit('changeBtnStatus', false)
        }
        if (res.status == 1) {
          commit('changeBtnStatus', true)
        }
      })
    }
  },
  getters: {
    getOpenStatus(state) {
      return state.openOr
    },
    getOpenLogPannelStatus(state) {
      return state.openLogPannel
    },
    getBtnstatus(state) {
      return state.btnStatus
    }
  }
}

export default initStorageModule
