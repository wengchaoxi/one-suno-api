<template>
  <div class="login-container">
    <h2>登录</h2>
    <form @submit.prevent="handleLogin">
      <div class="form-group">
        <label for="username">用户名</label>
        <input 
          type="text" 
          id="username" 
          v-model="loginForm.username" 
          required 
        />
      </div>
      
      <div class="form-group">
        <label for="password">密码</label>
        <input 
          type="password" 
          id="password" 
          v-model="loginForm.password" 
          required 
        />
      </div>
      
      <div v-if="errorMessage" class="error-message">
        {{ errorMessage }}
      </div>
      
      <button type="submit" :disabled="isLoading">
        {{ isLoading ? '登录中...' : '登录' }}
      </button>
    </form>
  </div>
</template>

<script lang="ts">
import { defineComponent, reactive, ref } from 'vue'
import { useRouter } from 'vue-router'
import authService from '@/api/auth'

export default defineComponent({
  name: 'LoginView',
  setup() {
    const router = useRouter()
    const isLoading = ref(false)
    const errorMessage = ref('')
    
    const loginForm = reactive({
      username: '',
      password: ''
    })
    
    const handleLogin = async () => {
      try {
        isLoading.value = true
        errorMessage.value = ''
        
        // 调用登录 API
        const response = await authService.login(loginForm.username, loginForm.password)
        
        // 存储 token
        localStorage.setItem('token', response.token)
        
        // 跳转到首页
        router.push('/')
      } catch (error) {
        errorMessage.value = '登录失败，请检查用户名和密码'
        console.error('Login error:', error)
      } finally {
        isLoading.value = false
      }
    }
    
    return {
      loginForm,
      isLoading,
      errorMessage,
      handleLogin
    }
  }
})
</script>

<style scoped>
.login-container {
  max-width: 400px;
  margin: 0 auto;
  padding: 20px;
}

.form-group {
  margin-bottom: 15px;
}

label {
  display: block;
  margin-bottom: 5px;
}

input {
  width: 100%;
  padding: 8px;
  border: 1px solid #ddd;
  border-radius: 4px;
}

button {
  width: 100%;
  padding: 10px;
  background-color: #4CAF50;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

button:disabled {
  background-color: #cccccc;
}

.error-message {
  color: red;
  margin: 10px 0;
}
</style>
