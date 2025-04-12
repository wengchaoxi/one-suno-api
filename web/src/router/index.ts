import { createRouter, createWebHistory } from 'vue-router'
import type { RouteRecordRaw } from 'vue-router'
import HomeView from '@/views/HomePage.vue'
import LoginView from '@/views/LoginPage.vue'
import authService from '@/api/auth'

const routes: Array<RouteRecordRaw> = [
  {
    path: '/',
    name: 'home',
    component: HomeView,
    meta: { requiresAuth: true },
  },
  {
    path: '/login',
    name: 'login',
    component: LoginView,
  },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

// 导航守卫 - 检查用户是否已登录
router.beforeEach((to) => {
  // 如果路由需要验证
  if (to.matched.some((record) => record.meta.requiresAuth)) {
    // 检查用户是否已登录
    if (!authService.isLoggedIn() && to.name !== 'login') {
      return { name: 'login' }
    } else {
      // 已登录则允许访问
      return true
    }
  } else {
    // 不需要验证的路由直接访问
    return true
  }
})

export default router
