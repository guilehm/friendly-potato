import axios, { AxiosInstance, AxiosResponse } from 'axios'
import { Cookies } from 'react-cookie'

const API_URL = process.env.API_URL || 'http://localhost:8080'

type LoginResponse = {
  id: string
  token: string
  refresh_token: string
}

class ApiService {
  baseUrl: string
  client: AxiosInstance

  constructor(baseUrl: string = API_URL) {
    this.baseUrl = `${baseUrl}`
    this.client = this.getClient()
  }

  getClient(): AxiosInstance {
    const cookies = new Cookies(['access'])

    const client = axios.create()
    client.interceptors.request.use(
      config => {
        const accessToken = cookies.get('access')
        if (accessToken) {
          config.headers['Authorization'] = accessToken
        }
        return config
      },
      error => {
        Promise.reject(error)
      }
    )
    return client
  }

  getGameList(): Promise<AxiosResponse> {
    return this.client.get(`${this.baseUrl}/games/`)
  }

  createUser(email: string, password: string): Promise<AxiosResponse> {
    return this.client.post(`${this.baseUrl}/users/`, { email, password })
  }

  login(email: string, password: string): Promise<AxiosResponse> {
    return this.client.post(`${this.baseUrl}/users/login/`, { email, password })
  }

}


export default ApiService
export type { LoginResponse }
