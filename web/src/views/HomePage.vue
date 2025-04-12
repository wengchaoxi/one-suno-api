<template>
  <div>
    <h1>管理后台</h1>
    <nav>
      <router-link to="/providers">Providers 管理</router-link> |
      <router-link to="/api-keys">API Keys 管理</router-link> |
      <router-link to="/logout" @click="logout">退出登录</router-link>
    </nav>
    <router-view />
  </div>
</template>

<script setup lang="ts">
import authService from '@/api/auth'
import router from '@/router'

const logout = () => {
  authService.logout()
  router.push({ path: '/login' })
}
</script>
