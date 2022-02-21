import React from 'react'
import { useQuery } from 'react-query'
import { UserGet } from '../../bv5-server'
import { bv5V1GetMe } from '../../services/bardview5'
import { Col, Container, Row, Spinner } from 'react-bootstrap'
import { AccessBadge } from '../../components/AccessBadge'
import { formatDistance } from 'date-fns'

const Profile: React.FC<{}> = () => {
  const {
    data,
    isLoading,
    isError
  } = useQuery<UserGet>('me', async () => {
    return (await bv5V1GetMe()).data
  })

  if (isLoading) {
    return <Spinner
            as="span"
            animation="grow"
            size="sm"
            role="status"
            aria-hidden="true"
    />
  }

  if (data === undefined || isError) {
    return <p>Failed to load</p>
  }

  const daysAgo = formatDistance(new Date(data.created), new Date(), { addSuffix: true })

  return <Container fluid="lg">
    <Row className="mb-3">
      <Col><h6>Name</h6></Col>
      <Col><p>{data.name}</p></Col>
    </Row>

    <Row className="mb-3">
      <Col><h6>Email</h6></Col>
      <Col><p>{data.email}</p></Col>
    </Row>

    <Row className="mb-3">
      <Col><h6>Access</h6></Col>
      <Col><AccessBadge accessType={data.commonAccess}/></Col>
    </Row>

    <Row className="mb-3">
      <Col><h6>Created</h6></Col>
      <Col><p>{daysAgo}</p></Col>
    </Row>

    <Row className="mb-3">
      <Col><h6>User Id</h6></Col>
      <Col><p>{data.userId}</p></Col>
    </Row>

    <Row className="mb-3">
      <Col><h6>Tags</h6></Col>
      <Col>
        { data.userTags.length === 0
          ? <p><i>None</i></p>
          : <ul>
                  { data.userTags.map(x => <li key={x}>{x}</li>) }
                </ul>
        }
      </Col>
    </Row>

    <Row className="mb-3">
      <Col><h6>System Tags</h6></Col>
      <Col>
        { data.systemTags.length === 0
          ? <p><i>None</i></p>
          : <ul>
                  { data.systemTags.map(x => <li key={x}>{x}</li>) }
                </ul>
        }
      </Col>
    </Row>

  </Container>
  // <Form as={Container} fluid="lg">
  //   <Form.Group as={Row} className="mb-3" controlId="formPlaintextEmail">
  //     <Form.Label column sm="2">
  //       Name
  //     </Form.Label>
  //     <Col sm="10">
  //       <Form.Control plaintext readOnly defaultValue={data.name}/>
  //     </Col>
  //   </Form.Group>
  //
  //   <Form.Group as={Row} className="mb-3" controlId="formPlaintextEmail">
  //     <Form.Label column sm="2">
  //       Email
  //     </Form.Label>
  //     <Col sm="10">
  //       <Form.Control plaintext readOnly defaultValue={data.email}/>
  //     </Col>
  //   </Form.Group>
  // </Form>
}

export default Profile
