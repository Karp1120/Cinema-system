import { createRouter, createWebHistory } from 'vue-router'

import Films from '../views/Films.vue'
import Cinemas from '../views/Cinemas.vue'
import Sessions from '../views/Sessions.vue'
import Reports from '../views/Reports.vue'
import Login from '../views/Login.vue'

const routes = [
  { path: '/', redirect: '/films' },
  { path: '/films', component: Films },
  { path: '/cinemas', component: Cinemas },
  { path: '/sessions', component: Sessions },
  { path: '/reports', component: Reports },
  { path: '/login', component: Login }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

// Глобальный перехват маршрутов для защиты
router.beforeEach((to, from, next) => {
  const publicPages = ['/login']
  const authRequired = !publicPages.includes(to.path)
  const loggedIn = !!localStorage.getItem('user')

  if (authRequired && !loggedIn) {
    return next('/login')
  }

  next()
})

export default router
