/* eslint-disable no-useless-constructor */
import axios, { AxiosResponse } from 'axios'
import { JsonError, SelfServiceRegistrationFlow } from '@ory/kratos-client'

type ApiResponseCategory = 'Not Found' | 'Unauthorized' | 'Unknown' | 'Ok' | 'Bad Request'

class ApiResponse<T> {
  constructor (public readonly status: number, public readonly category: ApiResponseCategory, public readonly data:T) {
  }
}

function convertAxiosResponse<T> (response: AxiosResponse<T>): ApiResponse<T> {
  let category: ApiResponseCategory = 'Unknown'
  if (response.status === 400) {
    category = 'Bad Request'
  }
  if (response.status === 401) {
    category = 'Unauthorized'
  }
  if (response.status === 404) {
    category = 'Not Found'
  }
  if (response.status === 200) {
    category = 'Ok'
  }
  return new ApiResponse(response.status, category, response.data)
}

async function startSelfServiceRegister (): Promise<ApiResponse<SelfServiceRegistrationFlow | JsonError>> {
  const response: AxiosResponse<SelfServiceRegistrationFlow> = await axios.get('/self-service/registration/browser', {
    validateStatus: undefined
  })
  return convertAxiosResponse(response)
}

async function submitSelfServiceRegister (data: any, flowUrl: string): Promise<ApiResponse<SelfServiceRegistrationFlow>> {
  const response: AxiosResponse<SelfServiceRegistrationFlow> = await axios.post(flowUrl, data, {
    validateStatus: undefined
  })
  return convertAxiosResponse(response)
}

export {
  startSelfServiceRegister,
  submitSelfServiceRegister,
  ApiResponse
}
export type {
  ApiResponseCategory
}
