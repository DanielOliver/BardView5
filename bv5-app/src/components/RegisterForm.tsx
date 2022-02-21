import { UiContainer } from '@ory/kratos-client'
import { Controller, useForm } from 'react-hook-form'
import React from 'react'
import { Button, Form } from 'react-bootstrap'
import { Link } from 'react-router-dom'

export function RegisterForm ({
  ui,
  onSubmit,
  formType
}: { ui: UiContainer, onSubmit: any, formType: 'Login' | 'Register' }) {
  const {
    register,
    handleSubmit,
    control
  } = useForm()

  let alternative = <Link to="/register"><i className="bi-box-arrow-in-right px-1"></i>Returning user?</Link>
  if (formType === 'Register') {
    alternative = <Link to="/login"><i className="bi-box-arrow-in-right px-1"></i>Existing user?</Link>
  }

  return (<div>
            <h1>{formType}</h1>

            <Form onSubmit={handleSubmit(onSubmit)} validated={(ui.messages?.length ?? 0) === 0}>
              {ui.nodes.filter(node => node.type === 'input').map(node => {
                if ('name' in node.attributes) {
                  const hidden = node.attributes.type === 'hidden'
                  const disabled = node.attributes.disabled
                  if (node.attributes.type === 'submit') {
                    return <div key={node.attributes.name}>
                              <input hidden
                                     type={node.attributes.type}
                                     value={node.attributes.value}
                                     {...register(node.attributes.name)}></input>
                            </div>
                  }
                  if ('required' in node.attributes) {
                    const attributes = node.attributes
                    const errorMessage = node.messages && node.messages.length > 0 ? node.messages.map(x => x.text).concat(' ') : undefined
                    return <Form.Group key={node.attributes.name}>
                              <Form.Label>{node.meta?.label?.text}</Form.Label>
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
                                                  <Form.Control required={attributes.required}
                                                                readOnly={disabled}
                                                                hidden={hidden}
                                                                placeholder={node.meta?.label?.text}
                                                                onBlur={onBlur}
                                                                onChange={onChange}
                                                                name={name}
                                                                type={attributes.type}
                                                                isInvalid={(errorMessage?.length ?? 0) > 0 || (ui.messages?.length ?? 0) > 0}
                                                  />)}
                              />
                              <Form.Control.Feedback type="invalid">
                                {errorMessage}
                              </Form.Control.Feedback>
                            </Form.Group>
                  }
                }
                return <></>
              }
              )
              }
              <Form.Group>
                <Form.Control isInvalid={(ui.messages?.length ?? 0) > 0} hidden/>
                {ui.messages?.map(message => (
                        <Form.Control.Feedback key={message.id} type="invalid">{message.text}
                        </Form.Control.Feedback>)
                )
                }
              </Form.Group>
              <Button type="submit">{formType}</Button>
            </Form>

            <br/>
            <br/>
            {alternative}
          </div>
  )
}
