import { useQuery } from 'react-query'
import { JsonError, SelfServiceRegistrationFlow } from '@ory/kratos-client'
import {
  ApiResponse, isJsonError,
  isSelfServiceRegistrationFlow, isSuccessfulSelfServiceRegistrationWithoutBrowser,
  startSelfServiceRegister,
  submitSelfServiceRegister
} from '../services/auth'
import React, { useEffect, useState } from 'react'
import { Controller, useForm } from 'react-hook-form'
import { AuthContext } from '../context/Auth.context'
import { Form, Message } from 'semantic-ui-react'

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

  const {
    register,
    handleSubmit,
    control
  } = useForm()

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
      return (
              <Form onSubmit={handleSubmit(onSubmit)} error={(registrationData.ui.messages?.length ?? 0) > 0}>
                {registrationData.ui.nodes.filter(node => node.type === 'input').map(node => {
                  if ('name' in node.attributes) {
                    const hidden = node.attributes.type === 'hidden'
                    const disabled = node.attributes.disabled
                    if (node.attributes.type === 'submit') {
                      return <div key={node.attributes.name}>
                                <Form.Field hidden>
                                  <input hidden
                                         type={node.attributes.type}
                                         value={node.attributes.value}
                                         {...register(node.attributes.name)}></input>
                                </Form.Field>
                              </div>
                    }
                    if ('required' in node.attributes) {
                      const attributes = node.attributes
                      const errorMessage = node.messages && node.messages.length > 0 ? node.messages.map(x => x.text).concat(' ') : undefined
                      return <div key={node.attributes.name}>
                                <Controller name={node.attributes.name}
                                            defaultValue={node.attributes.value}
                                            control={control}
                                            render={({
                                              field: {
                                                onChange,
                                                onBlur,
                                                name,
                                                ref
                                              }
                                            }) => (
                                                    <Form.Field>
                                                      <label>{node.meta?.label?.text}</label>
                                                      <Form.Input required={attributes.required}
                                                                  readOnly={disabled}
                                                                  hidden={hidden}
                                                                  fluid
                                                                  onBlur={onBlur}
                                                                  onChange={onChange}
                                                                  name={name}
                                                                  type={attributes.type}
                                                                  error={errorMessage}
                                                              // ref={ref}
                                                      />
                                                    </Form.Field>)}
                                />
                              </div>
                    }
                  }
                  return <></>
                }
                )
                }
                {registrationData.ui.messages?.map(message => (
                        <Message error key={message.id} content={message.text}/>))
                }
                <Form.Button>Register</Form.Button>
              </Form>
      )
    default:
      return <pre>
    {JSON.stringify(registrationData, null, 2)}
  </pre>
  }
}

export default RegisterRoute
