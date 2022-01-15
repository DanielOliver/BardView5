import axios, { AxiosResponse } from 'axios'
import { UserGet } from '../bv5-server'

async function bv5V1GetMe (): Promise<AxiosResponse<UserGet>> {
  return await axios.get<UserGet>('/api/v1/users/me')
}

export {
  bv5V1GetMe
}
