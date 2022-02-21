import { Link } from 'react-router-dom'
import React from 'react'
import { Container, Nav, Navbar, NavDropdown } from 'react-bootstrap'
import { AuthContext } from '../context/Auth.context'
import { bv5V1GetMe } from '../services/bardview5'
import { UserGet } from '../bv5-server'
import { useQuery } from 'react-query'
import { getSelfServiceLogout } from '../services/auth'

const Layout: React.FC<{ title?: string }> = ({
  title
}) => {
  const {
    state,
    dispatch
  } = React.useContext(AuthContext)

  const submitLogout = () => {
    getSelfServiceLogout().then(value => {
      dispatch({
        type: 'LOGOUT',
        isRegistrationComplete: false
      })
      if (value) {
        window.location.href = value.data.logout_url
      }
    })
  }

  const {
    data
  } = useQuery<UserGet>('me', async () => {
    return (await bv5V1GetMe()).data
  })

  return <Navbar bg="light" expand="md">
    <Container fluid>
      <Navbar.Brand><Link style={{ textDecoration: 'none' }} className="link-primary"
                          to="/">BardView5</Link></Navbar.Brand>
      {title && <Navbar.Text>{title}</Navbar.Text>}
      <Navbar.Toggle aria-controls="basic-navbar-nav"/>
      <Navbar.Collapse id="basic-navbar-nav">
        <Nav className="me-auto">
        </Nav>
        {state.isAuthenticated
          ? (
                        <Nav>
                          <NavDropdown title={data?.name ?? 'me'}>
                            <NavDropdown.Item>Profile</NavDropdown.Item>
                            <NavDropdown.Divider/>
                            <NavDropdown.Item onClick={submitLogout}>Logout</NavDropdown.Item>
                          </NavDropdown>
                        </Nav>)
          : (
                        <Nav>
                          <Nav.Item className="px-1"><Link to="/login"><i className="bi-box-arrow-in-right px-1"></i>Login</Link></Nav.Item>
                          <Nav.Item className="px-1"><Link to="/register"><i
                                  className="bi-box-arrow-in-right px-1"></i>Register</Link></Nav.Item>
                        </Nav>
            )
        }
      </Navbar.Collapse>
    </Container>
  </Navbar>
}

export {
  Layout
}
