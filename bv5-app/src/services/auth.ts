/* eslint-disable no-useless-constructor */
import axios, { AxiosResponse } from 'axios'
import { JsonError, SelfServiceRegistrationFlow, Session } from '@ory/kratos-client'
import { SuccessfulSelfServiceRegistrationWithoutBrowser } from '@ory/kratos-client/dist/api'

const isSelfServiceRegistrationFlow = (variableToCheck: any): variableToCheck is SelfServiceRegistrationFlow => (variableToCheck as SelfServiceRegistrationFlow).ui !== undefined
const isJsonError = (variableToCheck: any): variableToCheck is JsonError => (variableToCheck as JsonError).error !== undefined
const isSuccessfulSelfServiceRegistrationWithoutBrowser = (variableToCheck: any): variableToCheck is SelfServiceRegistrationFlow => (variableToCheck as SuccessfulSelfServiceRegistrationWithoutBrowser).session !== undefined

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

async function startSelfServiceRegister (): Promise<ApiResponse<SelfServiceRegistrationFlow> | ApiResponse<JsonError>> {
  const response = await axios.get('/self-service/registration/browser', {
    validateStatus: undefined
  })
  return convertAxiosResponse(response)
}

async function submitSelfServiceRegister (data: any, flowUrl: string): Promise<ApiResponse<SelfServiceRegistrationFlow> | ApiResponse<SuccessfulSelfServiceRegistrationWithoutBrowser>> {
  const response = await axios.post(flowUrl, data, {
    validateStatus: undefined
  })
  return convertAxiosResponse(response)
}

async function getSession (): Promise<ApiResponse<Session> | null> {
  const response = await axios.get('/sessions/whoami', {
    validateStatus: status => {
      return (status >= 200 && status < 300) || status === 401
    }
  })
  if (response.status === 401) {
    return null
  }
  return convertAxiosResponse(response)
}

export {
  startSelfServiceRegister,
  submitSelfServiceRegister,
  getSession,
  ApiResponse,

  isJsonError,
  isSuccessfulSelfServiceRegistrationWithoutBrowser,
  isSelfServiceRegistrationFlow
}
export type {
  ApiResponseCategory
}
