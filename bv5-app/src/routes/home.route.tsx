import { useQuery } from 'react-query'
import React from 'react'
import { bv5V1GetDnd5eWorldsMine } from '../services/bardview5'
import { Col, Container, ListGroup, Row } from 'react-bootstrap'
import { Dnd5eWorldGet } from '../bv5-server'
import { formatDistance } from 'date-fns'
import { AccessBadge } from '../components/AccessBadge'
import { Link } from 'react-router-dom'

function HomeRoute () {
  const {
    data: dnd5eWorlds
  } = useQuery<Dnd5eWorldGet[]>('my-worlds', async () => {
    const { data } = await bv5V1GetDnd5eWorldsMine()
    return data
  })

  return <Container fluid="lg">
      <Row><Col><h1 className="m-1">D&D 5e</h1></Col></Row>
      <Row>
        <Col md={6}>
          <h2>Worlds</h2>
          <ListGroup>
            {dnd5eWorlds?.map((world) => {
              const daysAgo = formatDistance(new Date(world.created), new Date(), { addSuffix: true })
              const shortDescription = world.description.substring(0, 200) + (world.description.length > 200 ? ' ...' : '')

              return (
                      <ListGroup.Item key={world.dnd5eWorldId} as={Link} to={`/dnd5e/worlds/${world.dnd5eWorldId}`}>
                        {/* <Link to={`/dnd5e/worlds/${world.dnd5eWorldId}`} > */}
                        <div className="d-flex justify-content-between flex-column flex-md-row">
                          <h5 className="mb-1">{world.name}</h5>
                          <div className="mb-1 justify-content-between d-flex">
                            <small className="me-1">{daysAgo}</small>
                            <AccessBadge accessType={world.commonAccess}/>
                          </div>
                        </div>
                        <p className="mb-1"><i>{shortDescription}</i></p>
                        {/* </Link> */}
                      </ListGroup.Item>
              )
            })
            }
          </ListGroup>
        </Col>
        <Col md={6}>
        </Col>
      </Row>
    </Container>
}

export default HomeRoute
