import React from 'react'
import { Button, Col, Container, Form, Row, Spinner } from 'react-bootstrap'
import { z } from 'zod'
import { zodResolver } from '@hookform/resolvers/zod'
import { useForm } from 'react-hook-form'
import { useMutation } from 'react-query'
import { bv5V1CreateDnd5eWorld } from '../../../services/bardview5'
import { Dnd5eWorldPostOk } from '../../../bv5-server'

const CommonAccessValues = ['private', 'anyuser', 'public'] as const
const CommonAccessText = ['Private', 'Any User', 'Public'] as const

const WorldCreateSchema = z.object({
  name: z.string().min(1).max(512),
  description: z.string().min(1).max(1024),
  module: z.string().optional(),
  commonAccess: z.enum(CommonAccessValues).default('private'),
  userTags: z.string().array().default([]),
  systemTags: z.string().array().default([]),
  active: z.boolean().default(true)
})

type WorldCreate = z.infer<typeof WorldCreateSchema>

function Dnd5eWorldCreate () {
  const mutation = useMutation<Dnd5eWorldPostOk, unknown, WorldCreate>(async (world: WorldCreate) => {
    const { data } = await bv5V1CreateDnd5eWorld({
      name: world.name,
      description: world.description,
      module: world.module,
      userTags: world.userTags,
      systemTags: world.systemTags,
      commonAccess: world.commonAccess,
      active: world.active
    })
    return data
  })

  const {
    register,
    handleSubmit,
    formState: { errors }
  } = useForm<WorldCreate>({
    resolver: zodResolver(WorldCreateSchema)
  })

  return <Container fluid="lg">
    <Row>
      <h1>D&D 5e World: Create</h1>
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

export default Dnd5eWorldCreate
