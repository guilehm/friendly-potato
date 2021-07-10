import axios, { AxiosInstance, AxiosResponse } from 'axios'

const API_URL = process.env.API_URL || 'http://localhost:8080'

class ApiService {
  baseUrl: string
  client: AxiosInstance

  constructor(baseUrl: string = API_URL) {
    this.baseUrl = `${baseUrl}`
    this.client = axios.create()
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
