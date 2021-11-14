import { useQuery } from 'react-query'
import { JsonError, SelfServiceRegistrationFlow } from '@ory/kratos-client'
import { ApiResponse, startSelfServiceRegister, submitSelfServiceRegister } from '../services/auth'
import React, { useEffect, useState } from 'react'
import { useForm } from 'react-hook-form'
import { AuthContext } from '../context/Auth.context'
import { SuccessfulSelfServiceRegistrationWithoutBrowser } from '@ory/kratos-client/dist/api'

const isSelfServiceRegistrationFlow = (variableToCheck: any): variableToCheck is SelfServiceRegistrationFlow => (variableToCheck as SelfServiceRegistrationFlow).ui !== undefined
const isJsonError = (variableToCheck: any): variableToCheck is JsonError => (variableToCheck as JsonError).error !== undefined
const isSuccessfulSelfServiceRegistrationWithoutBrowser = (variableToCheck: any): variableToCheck is SelfServiceRegistrationFlow => (variableToCheck as SuccessfulSelfServiceRegistrationWithoutBrowser).session !== undefined

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
    handleSubmit
    // watch,
    // formState: { errors }
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

  if (state.isAuthenticated) return <p>You are already registered you silly goose!</p>

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

    // axios.post<SelfServiceRegistrationFlow>(apicall.data.ui.action, data, {
    //   validateStatus: function (status) {
    //     return (status >= 200 && status < 300) || status === 400
    //   }
    // }).then(x => {
    //   console.log(JSON.stringify(x.data))
    //   set(x.data)
    // })
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
              /* "handleSubmit" will validate your inputs before invoking "onSubmit" */
              <form onSubmit={handleSubmit(onSubmit)}>
                {registrationData.ui.nodes.filter(node => node.type === 'input').map(node => {
                  if ('name' in node.attributes) {
                    const hidden = node.attributes.type === 'hidden'
                    const disabled = node.attributes.disabled
                    if (node.attributes.type === 'submit') {
                      return <button
                                      hidden={hidden}
                                      type={node.attributes.type}
                                      value={node.attributes.value}
                                      {...register(node.attributes.name)}>
                                {node?.meta?.label?.text}
                              </button>
                    }
                    return <div key={node.attributes.name}>
                              <label>{node?.meta?.label?.text}
                                <input required={node.attributes.required}
                                       readOnly={disabled}
                                       hidden={hidden}
                                       type={node.attributes.type}
                                       defaultValue={node.attributes.value}
                                       {...register(node.attributes.name)}/>
                              </label>
                              {node.messages && node.messages.length > 0 && node.messages.map(message => {
                                return <p key={message.id}>
                                  {message.text}
                                </p>
                              })

                              }
                              <br/>

                            </div>
                  }
                  return <></>
                }
                )
                }
                {registrationData.ui.messages?.map(message => <>
                  <p>{message.text}</p>
                  <br/>
                </>)

                }
              </form>
      )
    default:
      return <pre>
    {JSON.stringify(registrationData, null, 2)}
  </pre>
  }
}

export default RegisterRoute
