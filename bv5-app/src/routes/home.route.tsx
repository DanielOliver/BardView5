import { useQuery } from 'react-query'
import React from 'react'
import { bv5V1GetDnd5eWorldsMine } from '../services/bardview5'
import { Col, Container, ListGroup, Row } from 'react-bootstrap'
import { Dnd5eWorldGet } from '../bv5-server'
import { formatDistance } from 'date-fns'

function HomeRoute () {
  const {
    data: dnd5eWorlds
  } = useQuery<Dnd5eWorldGet[]>('my-worlds', async () => {
    return (await bv5V1GetDnd5eWorldsMine()).data
  })

  return <Container>
    <Row><h1>D&D 5e</h1></Row>
    <Row>
      <Col md={6}>
        <h2>Worlds</h2>
        <ListGroup>
          {dnd5eWorlds?.map((world) => {
            const daysAgo = formatDistance(new Date(world.created), new Date(), { addSuffix: true })
            const shortDescription = world.description.substring(0, 200) + (world.description.length > 200 ? ' ...' : '')

            return (
                    <ListGroup.Item key={world.dnd5eWorldId}>
                      <div className="d-flex w-100 justify-content-between">
                        <h5 className="mb-1">{world.name}</h5>
                        <small>{daysAgo}</small>
                      </div>
                      <p className="mb-1"><i>{shortDescription}</i></p>
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
