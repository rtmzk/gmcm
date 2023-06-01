import { createStore, Store, useStore as useVueStore } from 'vuex'
import { IRootState, IStoreType } from './types'

import login from './login/login'
import addServer from './main/addServer'
import initStorage from './main/initStorage'
import server from './main/servers'

const store = createStore<IRootState>({
  mutations: {},
  getters: {},
  actions: {},
  modules: {
    login,
    addServer,
    initStorage,
    server
  }
})

export function setupStore() {
  store.dispatch('login/loadLocalLogin')
  store.dispatch('addServer/loadLocalCheckRule')
  store.dispatch('initStorage/loadBtnStatus')
}

export function useStore(): Store<IStoreType> {
  return useVueStore()
}

export default store
