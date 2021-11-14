import { SelfServiceRegistrationFlow } from '@ory/kratos-client'
import { Controller, useForm } from 'react-hook-form'
import { Form, Message } from 'semantic-ui-react'
import React from 'react'

export function RegisterForm ({
  registrationData,
  onSubmit
}: { registrationData: SelfServiceRegistrationFlow, onSubmit: any }) {
  const {
    register,
    handleSubmit,
    control
  } = useForm()

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
                                name
                              }
                            }) => (
                              <Form.Input required={attributes.required}
                                          readOnly={disabled}
                                          hidden={hidden}
                                          fluid
                                          placeholder={node.meta?.label?.text}
                                          label={node.meta?.label?.text}
                                          onBlur={onBlur}
                                          onChange={onChange}
                                          name={name}
                                          type={attributes.type}
                                          error={errorMessage}
                              />)}
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
}
