import Vue from 'vue'
import vuetify from '@/plugins/vuetify' // path to vuetify export
import App from '@/App.vue'
import VueRouter from 'vue-router'
import Index from '@/pages/index'

Vue.use(VueRouter)

Vue.config.productionTip = false
Vue.prototype.renderData = window.$renderData

let router = new VueRouter({
  routes: [
    { path: '/', name: 'index', component: Index }
  ]
});

new Vue({
  vuetify,
  App,
  router,
  render: h => h(App)
}).$mount('#app')
