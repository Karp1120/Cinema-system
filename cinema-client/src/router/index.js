import { createRouter, createWebHistory } from 'vue-router'

import Films from '../views/Films.vue'
import Cinemas from '../views/Cinemas.vue'
import Sessions from '../views/Sessions.vue'
import Reports from '../views/Reports.vue'

const routes = [
  { path: '/', redirect: '/films' },
  { path: '/films', component: Films },
  { path: '/cinemas', component: Cinemas },
  { path: '/sessions', component: Sessions },
  { path: '/reports', component: Reports },
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router
