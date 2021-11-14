import { useQuery } from 'react-query'
import { JsonError, SelfServiceRegistrationFlow } from '@ory/kratos-client'
import {
  ApiResponse,
  isJsonError,
  isSelfServiceRegistrationFlow,
  isSuccessfulSelfServiceRegistrationWithoutBrowser,
  startSelfServiceRegister,
  submitSelfServiceRegister
} from '../services/auth'
import React, { useEffect, useState } from 'react'
import { AuthContext } from '../context/Auth.context'
import { RegisterForm } from '../components/RegisterForm'

function RegisterRoute () {
  const {
    dispatch,
    state
  } = React.useContext(AuthContext)

  const [apiResponseData, setApiResponseData] = useState<ApiResponse<any> | null>(null)
  const [registrationData, setRegistrationData] = useState<SelfServiceRegistrationFlow | JsonError | null>(null)

  const {
    isLoading,
    error,
    data: apicall
  } = useQuery<ApiResponse<SelfServiceRegistrationFlow> | ApiResponse<JsonError>>('login?', async () => {
    return await startSelfServiceRegister()
  }, {
    refetchOnWindowFocus: false
  })

  useEffect(() => {
    if (!isLoading && !error && apicall) {
      setRegistrationData(apicall.data)
      setApiResponseData(apicall)

      if (isJsonError(apicall.data)) {
        console.log('JsonError!', apicall.data)
        if (apicall.data.error.id === 'session_already_available') {
          console.log('Already logged in!')
          dispatch({
            type: 'LOGIN',
            isRegistrationComplete: false
          })
        }
      }
    }
  }, [apicall])

  if (state.isAuthenticated) {
    return <div>
      <p>You are already registered you silly goose!</p>
    </div>
  }

  if (error) return <p>Registration is currently unavailable.</p>

  if (isLoading || !apicall || !apiResponseData || !registrationData) return <p>Loading...</p>

  const onSubmit = async (data: any) => {
    console.log(data)
    if (isSelfServiceRegistrationFlow(registrationData)) {
      const result = await submitSelfServiceRegister(data, registrationData.ui.action)
      if (isSuccessfulSelfServiceRegistrationWithoutBrowser(result.data)) {
        dispatch({
          type: 'LOGIN',
          isRegistrationComplete: false
        })
      }
      if (isSelfServiceRegistrationFlow(result.data)) {
        setRegistrationData(result.data)
        setApiResponseData(result)
      }
    }
  }

  switch (apiResponseData?.category) {
    case 'Not Found':
      return <p>Not Found</p>
    case 'Ok':
    case 'Bad Request':
      if ('error' in registrationData) {
        return <pre>
          {JSON.stringify(registrationData, null, 2)}
          </pre>
      }
      return <RegisterForm registrationData={registrationData} onSubmit={onSubmit}/>
    default:
      return <pre>
    {JSON.stringify(registrationData, null, 2)}
  </pre>
  }
}

export default RegisterRoute
