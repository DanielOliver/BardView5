import React, { useEffect } from 'react'
import { useNavigate, useParams } from 'react-router-dom'
import { useMutation, useQuery } from 'react-query'
import { Dnd5eSettingGet, Dnd5eSettingPostOk } from '../../../bv5-server'
import { bv5V1EditDnd5eSetting, bv5V1GetDnd5eSetting } from '../../../services/bardview5'
import { AxiosResponse } from 'axios'
import { Button, Col, Container, Form, Row, Spinner, Tab, Tabs } from 'react-bootstrap'
import ReactMarkdown from 'react-markdown'
import remarkGfm from 'remark-gfm'
import { Dnd5eSettingCreateSchema, Dnd5eSettingCreateType } from './common'
import { useForm } from 'react-hook-form'
import { zodResolver } from '@hookform/resolvers/zod'
import { CommonAccessText, CommonAccessValues } from '../common'

const Dnd5eSettingForm: React.FC<{ dnd5eSettingId: string, data: Dnd5eSettingGet }> = ({
  dnd5eSettingId,
  data
}) => {
  const mutation = useMutation<Dnd5eSettingPostOk, unknown, Dnd5eSettingCreateType>(async (setting) => {
    const { data } = await bv5V1EditDnd5eSetting(dnd5eSettingId, {
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
    watch,
    formState: { errors }
  } = useForm<Dnd5eSettingCreateType>({
    resolver: zodResolver(Dnd5eSettingCreateSchema),
    defaultValues: {
      active: data.active,
      description: data.description,
      module: data.module,
      userTags: data.userTags,
      systemTags: data.systemTags,
      // @ts-ignore
      commonAccess: data.commonAccess,
      name: data.name
    }
  })

  useEffect(() => {
    if (mutation.isSuccess && mutation.data?.dnd5eSettingId !== undefined) {
      navigate(`/dnd5e/settings/${dnd5eSettingId}`)
    }
  }, [mutation])

  const watchDescription = watch('description')

  return <Container fluid="lg">
    <Row>
      <p><i>D&D 5e Setting</i></p>
    </Row>
    <Row>
      <h1>Edit</h1>
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
        <Form.Group as={Col} className="mb-3">
          <Form.Label>Description</Form.Label>

          <Tabs defaultActiveKey="edit">
            <Tab eventKey="edit" title="Edit">
              <Form.Control as="textarea" required placeholder="Description" {...register('description')}
                            isInvalid={errors.description?.message !== undefined}/>
            </Tab>
            <Tab eventKey="preview" title="Preview">
              <ReactMarkdown remarkPlugins={[remarkGfm]}>
                {watchDescription}
              </ReactMarkdown>
              <hr/>
            </Tab>
          </Tabs>
          <Form.Control.Feedback type="invalid">
            {errors.description?.message}
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
              : <>Save</>
            }
          </Button>
        </Col>
      </Row>
    </Form>
  </Container>
}

const Dnd5eSettingEdit = () => {
  const params = useParams()
  const dnd5eSettingId: string = params.dnd5eSettingId ?? '0'
  const {
    data,
    error,
    isLoading
  } = useQuery<Dnd5eSettingGet, AxiosResponse>(`dnd5e-setting-${dnd5eSettingId}`, async () => {
    const { data } = await bv5V1GetDnd5eSetting(dnd5eSettingId)
    return data
  }, {
    retry: false
  })

  if (error) {
    return <Container fluid="lg">
      <p>Did not find setting {dnd5eSettingId}</p>
      <pre>
        {JSON.stringify(error, null, 2)}
      </pre>
    </Container>
  }

  if (isLoading || data === undefined) {
    return <Container fluid="lg">
      <Spinner animation="border" role="status">
        <span className="visually-hidden">Loading...</span>
      </Spinner>
    </Container>
  }

  return <Dnd5eSettingForm dnd5eSettingId={dnd5eSettingId} data={data}/>
}

export {
  Dnd5eSettingEdit
}
