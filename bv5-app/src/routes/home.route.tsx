import { useQuery } from 'react-query'
import React from 'react'
import { bv5V1GetDnd5eSettingsMine } from '../services/bardview5'
import { Col, Container, ListGroup, Row } from 'react-bootstrap'
import { Dnd5eSettingGet } from '../bv5-server'
import { formatDistance } from 'date-fns'
import { AccessBadge } from '../components/AccessBadge'
import { Link } from 'react-router-dom'

const Dnd5eSettingList: React.FC<{ dnd5eSettings: Dnd5eSettingGet[] }> = ({ dnd5eSettings }) => {
  return <ListGroup>
    {dnd5eSettings.map((setting) => {
      const daysAgo = formatDistance(new Date(setting.created), new Date(), { addSuffix: true })
      return (
              <ListGroup.Item key={setting.dnd5eSettingId} as={Link}
                              to={`/dnd5e/settings/${setting.dnd5eSettingId}`}>
                <div className="d-flex justify-content-between flex-column flex-md-row">
                  <h5 className="mb-1">{setting.name}</h5>
                  <div className="mb-1 justify-content-between d-flex">
                    <small className="me-1">{daysAgo}</small>
                    <AccessBadge accessType={setting.commonAccess}/>
                  </div>
                </div>
              </ListGroup.Item>
      )
    })
    }
  </ListGroup>
}

function HomeRoute () {
  const {
    data: dnd5eSettings
  } = useQuery<Dnd5eSettingGet[]>('dnd5e-setting-mine', async () => {
    const { data } = await bv5V1GetDnd5eSettingsMine()
    return data
  })

  return <Container fluid="lg">
    <Row><Col><h1 className="m-1">D&D 5e</h1></Col></Row>
    <Row>
      <Col md={6}>
        <h2>Settings</h2>
        {dnd5eSettings && dnd5eSettings?.length > 0
          ? <Dnd5eSettingList dnd5eSettings={dnd5eSettings}/>
          : <p>Create a new setting!</p>
        }
      </Col>
      <Col md={6}>
      </Col>
    </Row>
  </Container>
}

export default HomeRoute
