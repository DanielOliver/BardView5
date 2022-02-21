import axios, { AxiosResponse } from 'axios'
import { Dnd5eMonsterGet, Dnd5eSetting, Dnd5eSettingGet, Dnd5eSettingPostOk, UserGet } from '../bv5-server'

async function bv5V1GetMe (): Promise<AxiosResponse<UserGet>> {
  return await axios.get<UserGet>('/api/v1/users/me')
}

async function bv5V1GetDnd5eSettingsMine (): Promise<AxiosResponse<Dnd5eSettingGet[]>> {
  return await axios.get<Dnd5eSettingGet[]>('/api/v1/dnd5e/settings/assigned')
}

async function bv5V1GetDnd5eSettings (): Promise<AxiosResponse<Dnd5eSettingGet[]>> {
  return await axios.get<Dnd5eSettingGet[]>('/api/v1/dnd5e/settings')
}

async function bv5V1GetDnd5eSetting (dnd5eSettingId: string): Promise<AxiosResponse<Dnd5eSettingGet>> {
  return await axios.get<Dnd5eSettingGet>(`/api/v1/dnd5e/settings/${dnd5eSettingId}`)
}

async function bv5V1GetDnd5eMonstersBySetting (dnd5eSettingId: string): Promise<AxiosResponse<Dnd5eMonsterGet[]>> {
  return await axios.get<Dnd5eMonsterGet[]>(`/api/v1/dnd5e/settings/${dnd5eSettingId}/monsters`)
}

async function bv5V1CreateDnd5eSetting (setting: Dnd5eSetting): Promise<AxiosResponse<Dnd5eSettingPostOk>> {
  return await axios.post<Dnd5eSettingPostOk>('/api/v1/dnd5e/settings/', setting)
}

async function bv5V1EditDnd5eSetting (dnd5eSettingId: string, setting: Dnd5eSetting): Promise<AxiosResponse<Dnd5eSettingPostOk>> {
  return await axios.post<Dnd5eSettingPostOk>(`/api/v1/dnd5e/settings/${dnd5eSettingId}`, setting)
}

export {
  bv5V1GetMe,

  bv5V1GetDnd5eSettingsMine,
  bv5V1GetDnd5eSetting,
  bv5V1CreateDnd5eSetting,
  bv5V1EditDnd5eSetting,
  bv5V1GetDnd5eSettings,

  bv5V1GetDnd5eMonstersBySetting
}
