let BASE_URL = ''
const TIME_OUT = 30000

if (process.env.NODE_ENV === 'development') {
  BASE_URL = 'http://192.168.251.76:9527/api'
} else if (process.env.NODE_ENV === 'production') {
  BASE_URL = '/api'
}

export { BASE_URL, TIME_OUT }
