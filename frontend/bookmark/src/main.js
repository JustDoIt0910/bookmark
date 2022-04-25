import Vue from 'vue'
import App from './App.vue'
import vuetify from './plugins/vuetify'
import axios from 'axios'
import router from './router'

Vue.config.productionTip = false

axios.defaults.baseURL = 'http://localhost:3000/api/v1'
axios.interceptors.request.use(config => {
  config.headers.Authorization = `Bearer ${window.localStorage.getItem('token')}`
  return config
})

Vue.prototype.$http = axios

new Vue({
  vuetify,
  router,
  render: h => h(App)
}).$mount('#app')
