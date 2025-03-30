import api from "./api";

interface LoginResponse {
  token: string;
  user: {
    id: string;
    username: string;
  }
}

const authService = {
  async login(username: string, password: string): Promise<LoginResponse> {
    const response = await api.get(`/users/${username}`, {
        params: { password }
      })
    return response.data
  },
  
  // 检查用户是否已登录
  isLoggedIn(): boolean {
    return !!localStorage.getItem('token')
  },

  // 注销
  logout(): void {
    localStorage.removeItem('token')
  }
}

export default authService
