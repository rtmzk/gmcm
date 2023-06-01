// 统一出口
import { BASE_URL, TIME_OUT } from './request/config'
import DefaultRequest from './request/index'
import localCache from '@/utils/cache'

const defaultRequest = new DefaultRequest({
  baseURL: BASE_URL,
  timeout: TIME_OUT,
  interceptors: {
    requestInterceptor: (config) => {
      const token = localCache.getCache('token')
      if (token) {
        config.headers.Authorization = `Bearer ${token}`
      }
      return config
    }
  }
})

export default defaultRequest
