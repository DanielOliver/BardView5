import React, { useEffect } from 'react'
import { Button, Col, Container, Form, Row, Spinner } from 'react-bootstrap'
import { z } from 'zod'
import { zodResolver } from '@hookform/resolvers/zod'
import { useForm } from 'react-hook-form'
import { useMutation } from 'react-query'
import { bv5V1CreateDnd5eSetting } from '../../../services/bardview5'
import { Dnd5eSettingPostOk } from '../../../bv5-server'
import { useNavigate } from 'react-router-dom'

const CommonAccessValues = ['private', 'anyuser', 'public'] as const
const CommonAccessText = ['Private', 'Any User', 'Public'] as const

const SettingCreateSchema = z.object({
  name: z.string().min(1).max(512),
  description: z.string().min(1).max(1024),
  module: z.string().optional(),
  commonAccess: z.enum(CommonAccessValues).default('private'),
  userTags: z.string().array().default([]),
  systemTags: z.string().array().default([]),
  active: z.boolean().default(true)
})

type SettingCreate = z.infer<typeof SettingCreateSchema>

export function Dnd5eSettingCreate () {
  const mutation = useMutation<Dnd5eSettingPostOk, unknown, SettingCreate>(async (setting: SettingCreate) => {
    const { data } = await bv5V1CreateDnd5eSetting({
      name: setting.name,
      description: setting.description,
      module: setting.module,
      userTags: setting.userTags,
      systemTags: setting.systemTags,
      commonAccess: setting.commonAccess,
      active: setting.active
    })
    return data
  })

  const navigate = useNavigate()

  const {
    register,
    handleSubmit,
    formState: { errors }
  } = useForm<SettingCreate>({
    resolver: zodResolver(SettingCreateSchema)
  })

  useEffect(() => {
    if (mutation.isSuccess && mutation.data?.dnd5eSettingId !== undefined) {
      navigate(`/dnd5e/settings/${mutation.data.dnd5eSettingId}`)
    }
  }, [mutation])

  return <Container fluid="lg">
    <Row>
      <h1>D&D 5e Setting: Create</h1>
    </Row>
    <Form onSubmit={handleSubmit((d) => mutation.mutate(d))}>

      <Row>
        <Form.Group as={Col} className="mb-3">
          <Form.Label>Name</Form.Label>
          <Form.Control type="text" required placeholder="Name" {...register('name')}
                        isInvalid={errors.name?.message !== undefined}/>
          <Form.Control.Feedback type="invalid">
            {errors.name?.message}
          </Form.Control.Feedback>
        </Form.Group>
      </Row>

      <Row>
        <Form.Group as={Col} className="mb-3">
          <Form.Label>Description</Form.Label>
          <Form.Control as="textarea" required placeholder="Description" {...register('description')}
                        isInvalid={errors.description?.message !== undefined}/>
          <Form.Control.Feedback type="invalid">
            {errors.description?.message}
          </Form.Control.Feedback>
        </Form.Group>
      </Row>

      <Row>
        <Form.Group as={Col} md={true} className="mb-3">
          <Form.Label>Module</Form.Label>
          <Form.Control type="text" placeholder="Module (optional)" {...register('module')}
                        isInvalid={errors.module?.message !== undefined}/>
          <Form.Control.Feedback type="invalid">
            {errors.module?.message}
          </Form.Control.Feedback>
        </Form.Group>

        <Form.Group as={Col} md={true} className="mb-3">
          <Form.Label>Access</Form.Label>
          <Form.Select aria-label="Common Access" {...register('commonAccess')} >

            {CommonAccessValues.map((item, index) => (
                    <option value={item} key={index}>{CommonAccessText[index]}</option>
            ))
            }
          </Form.Select>

          <Form.Control.Feedback type="invalid">
            {errors.module?.message}
          </Form.Control.Feedback>
        </Form.Group>
      </Row>

      <Row>
        <Col>
          <Button variant="primary" type="submit"
                  disabled={mutation.isLoading || mutation.isError || mutation.isSuccess}>
            {mutation.isLoading
              ? <>
                      <Spinner
                              as="span"
                              animation="grow"
                              size="sm"
                              role="status"
                              aria-hidden="true"
                      />
                      Saving...
                    </>
              : <>Submit</>
            }
          </Button>
        </Col>
      </Row>
    </Form>
  </Container>
}

export default {
  Dnd5eSettingCreate
}
