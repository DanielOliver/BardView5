import React, { useState } from 'react'
import { Layout } from './Layout'
import { Button, Col, Container, Offcanvas, Row } from 'react-bootstrap'
import { NavLink } from 'react-router-dom'

const Menu = ({ onNav }: {
  onNav: () => void
}) => {
  return <ul className="list-unstyled">
    <li>
      <NavLink to="/dnd5e/settings" onClick={onNav}><strong>D&D 5e Settings</strong></NavLink>
      <ul className="list-unstyled ms-3">
        <li><NavLink to="/dnd5e/settings" onClick={onNav}>Search</NavLink></li>
        <li><NavLink to="/dnd5e/settings/create" onClick={onNav}>Create</NavLink></li>
      </ul>
    </li>
  </ul>
}

const LayoutWrapper: React.FC<{}> = ({
  children
}) => {
  const [show, setShow] = useState(false)

  const handleClose = () => setShow(false)
  const handleShow = () => setShow(true)

  return (<Container fluid>
            <Row> <Layout/></Row>
            <Row>
              <Col lg={2} className="d-none d-lg-block">
                <Menu onNav={() => {}}/>
              </Col>

              <Col xs={12} className="d-lg-none d-flex flex-row-reverse">
                <Button onClick={handleShow}
                        variant="primary"
                >
                  <i className={ show ? 'bi-arrows-expand' : 'bi-arrows-collapse'}></i>
                </Button>

                <Offcanvas show={show} onHide={handleClose}>
                  <Offcanvas.Header closeButton>
                    <Offcanvas.Title>Go to</Offcanvas.Title>
                  </Offcanvas.Header>
                  <Offcanvas.Body>
                    <Menu onNav={handleClose}/>
                  </Offcanvas.Body>
                </Offcanvas>
              </Col>

              <Col>
                {children}
              </Col>
            </Row>
          </Container>
  )
}

export default LayoutWrapper
