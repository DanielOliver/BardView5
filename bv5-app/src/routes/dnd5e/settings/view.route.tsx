import React from 'react'
import { useNavigate, useParams } from 'react-router-dom'
import { useQuery } from 'react-query'
import { Dnd5eSettingGet } from '../../../bv5-server'
import { bv5V1GetDnd5eSetting } from '../../../services/bardview5'
import { AxiosResponse } from 'axios'
import { Button, Col, Container, Row, Spinner } from 'react-bootstrap'
import ReactMarkdown from 'react-markdown'
import { AccessBadge } from '../../../components/AccessBadge'
import { formatDistance } from 'date-fns'
import remarkGfm from 'remark-gfm'

const Dnd5eSettingView: React.FC<{isAuthenticated: boolean}> = ({ isAuthenticated = false }) => {
  const params = useParams()
  const navigate = useNavigate()
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

  const daysAgo = formatDistance(new Date(data.created), new Date(), { addSuffix: true })

  return <Container fluid="lg">
    <Row>
      <p><i>D&D 5e Setting</i></p>
    </Row>
    <Row className="m-1">
      <h1>{data.name}</h1>
    </Row>

    <Row>
      <Col md={true}>
        <small><i>Created {daysAgo}</i></small>
      </Col>
      <Col md={true}>
        <AccessBadge accessType={data.commonAccess ?? ''}/>
      </Col>
      {isAuthenticated &&
              <Col md={true}>
                <Button variant="primary" onClick={() => {
                  navigate(`/dnd5e/settings/${dnd5eSettingId}/edit`)
                }}>Edit</Button>
              </Col>
      }
    </Row>
    {data?.module &&
            <Row>
              <small><i>Module:</i> {data.module}</small>
            </Row>
    }

    <Row className="m-1">
      <hr/>
      <ReactMarkdown remarkPlugins={[remarkGfm]}>
        {data.description}
      </ReactMarkdown>
    </Row>
  </Container>
}

export {
  Dnd5eSettingView
}
