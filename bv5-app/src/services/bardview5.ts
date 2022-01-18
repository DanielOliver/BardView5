import axios, { AxiosResponse } from 'axios'
import { Dnd5eWorldGet, UserGet } from '../bv5-server'

async function bv5V1GetMe (): Promise<AxiosResponse<UserGet>> {
  return await axios.get<UserGet>('/api/v1/users/me')
}

async function bv5V1GetDnd5eWorldsMine (): Promise<AxiosResponse<Dnd5eWorldGet[]>> {
  return await axios.get<Dnd5eWorldGet[]>('/api/v1/dnd5e/worlds/assigned')
}

async function bv5V1GetDnd5eWorld (dnd5eWorldId: string): Promise<AxiosResponse<Dnd5eWorldGet[]>> {
  return await axios.get<Dnd5eWorldGet[]>(`/api/v1/dnd5e/worlds/${dnd5eWorldId}`)
}

export {
  bv5V1GetMe,
  bv5V1GetDnd5eWorldsMine,
  bv5V1GetDnd5eWorld
}
