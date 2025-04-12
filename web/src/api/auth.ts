import api from './api'

interface HttpResponse<T> {
  code: number
  msg: string
  data?: T
}

type Token = {
  token: string
}
type LoginResponse = HttpResponse<Token>

const authService = {
  async login(username: string, password: string): Promise<LoginResponse> {
    const response = await api.post(
      `/users/token`,
      {
        password: password,
        username: username,
      },
      {
        headers: {
          'Content-Type': 'application/json',
        },
      },
    )
    return response.data
  },

  // 检查用户是否已登录
  isLoggedIn(): boolean {
    return !!localStorage.getItem('token')
  },

  // 注销
  logout(): void {
    localStorage.removeItem('token')
  },
}

export default authService
