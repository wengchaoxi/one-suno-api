import { createRouter, createWebHistory } from 'vue-router'
import type { RouteRecordRaw } from 'vue-router'
import HomeView from '../views/Home.vue'
import LoginView from '../views/Login.vue'
import authService from '@/api/auth'

const API_URL = 'http://localhost:8080/v1'

const routes: Array<RouteRecordRaw> = [
  {
    path: '/',
    name: 'home',
    component: HomeView,
    meta: { requiresAuth: true }
  },
  {
    path: '/login',
    name: 'login',
    component: LoginView
  }
]

const router = createRouter({
  history: createWebHistory(API_URL),
  routes
})

// 导航守卫 - 检查用户是否已登录
router.beforeEach((to, from, next) => {
  // 如果路由需要验证
  if (to.matched.some(record => record.meta.requiresAuth)) {
    // 检查用户是否已登录
    if (!authService.isLoggedIn()) {
      // 未登录则重定向到登录页
      next({
        path: '/login',
        query: { redirect: to.fullPath }
      })
    } else {
      // 已登录则允许访问
      next()
    }
  } else {
    // 不需要验证的路由直接访问
    next()
  }
})

export default router
