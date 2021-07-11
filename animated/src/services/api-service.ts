import axios, { AxiosInstance, AxiosResponse } from "axios"
import { Cookies } from "react-cookie"
import { ACCESS_TOKEN_LIFTIME, REFRESH_TOKEN_LIFTIME } from "../settings"

const API_URL = process.env.API_URL || "http://localhost:8080"

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
    const cookies = new Cookies(["access"])
    const client = axios.create()
    client.interceptors.request.use(
      config => {
        const accessToken = cookies.get("access")
        if (accessToken) {
          config.headers["Authorization"] = accessToken
        }
        return config
      },
      error => {
        Promise.reject(error)
      }
    )

    client.interceptors.response.use(
      response => response,
      async error => {
        const cookies = new Cookies(["access", "refresh"])
        const accessToken = cookies.get("access")
        const refreshToken = cookies.get("refresh")
        const originalRequest = error.config
        if (refreshToken && error.response.status === 401 && !originalRequest._retry) {
          originalRequest._retry = true
          const response = await this.refreshTokens(accessToken, refreshToken)
          axios.defaults.headers.common["Authorization"] = response.data.token
          cookies.set("access", response.data.token, {
            path: "/",
            maxAge: ACCESS_TOKEN_LIFTIME,
            sameSite: true,
          })
          cookies.set("refresh", response.data.refresh_token, {
            path: "/",
            maxAge: REFRESH_TOKEN_LIFTIME,
            sameSite: true,
          })
          return client(originalRequest)
        }
        return Promise.reject(error)
      })
    return client
  }

  refreshTokens(accessToken: string, refreshToken: string): Promise<AxiosResponse> {
    return this.client.post(`${this.baseUrl}/users/refresh/`, {
      access_token: accessToken,
      refresh_token: refreshToken,
    })
  }

  getGameList(): Promise<AxiosResponse> {
    return this.client.get(`${this.baseUrl}/games/`)
  }

  getGameDetail(slug: string): Promise<AxiosResponse> {
    return this.client.get(`${this.baseUrl}/games/${slug}`)
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
