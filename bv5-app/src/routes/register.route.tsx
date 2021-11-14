import { useQuery } from 'react-query'
import { JsonError, SelfServiceRegistrationFlow } from '@ory/kratos-client'
import { ApiResponse, startSelfServiceRegister, submitSelfServiceRegister } from '../services/auth'
import React, { useEffect, useState } from 'react'
import { useForm } from 'react-hook-form'

function RegisterRoute () {
  const [registrationData, setRegistrationData] = useState<ApiResponse<SelfServiceRegistrationFlow | JsonError> | null>(null)

  const {
    isLoading,
    error,
    data: apicall
  } = useQuery<ApiResponse<SelfServiceRegistrationFlow | JsonError>>('login?', async () => {
    return await startSelfServiceRegister()
  })

  const {
    register,
    handleSubmit
    // watch,
    // formState: { errors }
  } = useForm()

  useEffect(() => {
    if (!isLoading && !error && apicall) {
      setRegistrationData(apicall)
    }
  }, [apicall])

  if (error) return <p>Registration is currently unavailable.</p>

  // if (error) return <pre>An error has occurred: {JSON.stringify(error, null, 2)}</pre>

  if (isLoading || !apicall || !registrationData) return <p>Loading...</p>

  const onSubmit = async (data: any) => {
    console.log(data)
    if ('ui' in registrationData.data) {
      const result = await submitSelfServiceRegister(data, registrationData.data.ui.action)
      setRegistrationData(result)
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

  switch (registrationData.category) {
    case 'Not Found':
      return <p>Not Found</p>
    case 'Ok':
    case 'Bad Request':
      if ('error' in registrationData.data) {
        return <pre>
          {JSON.stringify(registrationData, null, 2)}
          </pre>
      }
      return (
              /* "handleSubmit" will validate your inputs before invoking "onSubmit" */
              <form onSubmit={handleSubmit(onSubmit)}>
                {registrationData.data.ui.nodes.filter(node => node.type === 'input').map(node => {
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
                                 value={node.attributes.value}
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
              </form>
      )
    default:
      return <pre>
    {JSON.stringify(registrationData, null, 2)}
  </pre>
  }
}

export default RegisterRoute
