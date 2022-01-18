import { Link } from 'react-router-dom'
import React from 'react'
import { Container, Nav, Navbar, NavDropdown } from 'react-bootstrap'
import { AuthContext } from '../context/Auth.context'
import { bv5V1GetMe } from '../services/bardview5'
import { UserGet } from '../bv5-server'
import { useQuery } from 'react-query'

export function Layout ({ logout }: {
  logout: () => void
}) {
  const { state } = React.useContext(AuthContext)

  const {
    data
  } = useQuery<UserGet>('me', async () => {
    return (await bv5V1GetMe()).data
  })

  return <Navbar bg="light" expand="md">
    <Container fluid="lg">
      <Navbar.Brand><Link style={{ textDecoration: 'none' }} className="link-primary" to="/">BardView5</Link></Navbar.Brand>
      <Navbar.Toggle aria-controls="basic-navbar-nav"/>
      <Navbar.Collapse id="basic-navbar-nav">
        <Nav className="me-auto">
          {/* <Nav.Link><Link style={{ textDecoration: 'none' }} to="/">Home</Link></Nav.Link> */}
        </Nav>

        {state.isAuthenticated
          ? (
                <Nav>
                  <NavDropdown title={data?.name ?? 'me'}>
                    <NavDropdown.Item>Profile</NavDropdown.Item>
                    <NavDropdown.Divider />
                    <NavDropdown.Item onClick={logout}>Logout</NavDropdown.Item>
                  </NavDropdown>
                </Nav>)
          : (
                <Nav>
                  <Nav.Link><Link to="/login"><i className="bi-box-arrow-in-right"></i>Login</Link></Nav.Link>
                  <Nav.Link><Link to="/register"><i className="bi-box-arrow-in-right"></i>Register</Link></Nav.Link>
                </Nav>
            )
        }
      </Navbar.Collapse>
    </Container>
  </Navbar>
}
