import { useQuery, useQueryClient } from 'react-query'
import { JsonError, SelfServiceLoginFlow } from '@ory/kratos-client'
import {
  ApiResponse,
  isJsonError, isSelfServiceLoginFlow, isSuccessfulSelfServiceLoginWithoutBrowser,
  startSelfServiceLogin, submitSelfServiceLogin
} from '../services/auth'
import React, { useEffect, useState } from 'react'
import { AuthContext } from '../context/Auth.context'
import { RegisterForm } from '../components/RegisterForm'
import { useNavigate } from 'react-router-dom'
import { Container } from 'react-bootstrap'
import { Bv5RouteProps } from '../components/Common'
import LayoutWrapper from '../components/LayoutWrapper'

function LoginFormWrapper () {
  const navigate = useNavigate()
  const {
    dispatch,
    state
  } = React.useContext(AuthContext)

  const [apiResponseData, setApiResponseData] = useState<ApiResponse<any> | null>(null)
  const [loginData, setLoginData] = useState<SelfServiceLoginFlow | JsonError | null>(null)

  const {
    isLoading,
    error,
    data: apicall
  } = useQuery<ApiResponse<SelfServiceLoginFlow> | ApiResponse<JsonError>>('login?', async () => {
    return await startSelfServiceLogin()
  }, {
    refetchOnWindowFocus: false
  })
  const queryClient = useQueryClient()

  useEffect(() => {
    if (!isLoading && !error && apicall) {
      setLoginData(apicall.data)
      setApiResponseData(apicall)

      if (isJsonError(apicall.data)) {
        if (apicall.data.error.id === 'session_already_available') {
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

  if (isLoading || !apicall || !apiResponseData || !loginData) return <p>Loading...</p>

  const onSubmit = async (data: any) => {
    if (isSelfServiceLoginFlow(loginData)) {
      const result = await submitSelfServiceLogin(data, loginData.ui.action)
      if (isSuccessfulSelfServiceLoginWithoutBrowser(result.data)) {
        dispatch({
          type: 'LOGIN',
          isRegistrationComplete: false
        })
        navigate('/', { replace: true })
      }
      if (isSelfServiceLoginFlow(result.data)) {
        setLoginData(result.data)
        setApiResponseData(result)
      }
      await queryClient.refetchQueries(['me'])
    }
  }

  switch (apiResponseData?.category) {
    case 'Not Found':
      return <p>Not Found</p>
    case 'Ok':
    case 'Bad Request':
      if ('error' in loginData) {
        return <pre>
          {JSON.stringify(loginData, null, 2)}
          </pre>
      }
      return <Container>
        <RegisterForm formType="Login" ui={loginData.ui} onSubmit={onSubmit}/>
      </Container>
    default:
      return <pre>
    {JSON.stringify(loginData, null, 2)}
  </pre>
  }
}

function RouteLogin ({ isAuthenticated }: Bv5RouteProps) {
  return <LayoutWrapper isAuthenticated={isAuthenticated} title="Login"><LoginFormWrapper/></LayoutWrapper>
}

export {
  RouteLogin
}
