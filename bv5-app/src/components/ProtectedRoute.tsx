import { Link, Outlet, RouteProps } from 'react-router-dom'
import React from 'react'
import { Layout } from './Layout'
import { Container } from 'react-bootstrap'

export type ProtectedRouteProps = {
  isAuthenticated: boolean;
  title?: string;
} & RouteProps;

// Originally https://stackoverflow.com/a/47754325
function ProtectedRoute ({
  isAuthenticated,
  title,
  ...routeProps
}: ProtectedRouteProps) {
  if (isAuthenticated) {
    return <Outlet/>
  } else {
    return <Layout title={title}>
      <Container fluid="lg">
        <p>A self-hosted Fifth Edition Dungeons & Dragons Campaign Management Tool.</p>
        <br/>
        <p>Go ahead and login. Adventure awaits!</p>
        <br/>
        <p><Link to="/login"><i className="bi-box-arrow-in-right px-1"></i>Login</Link> or <Link to="/register"><i
                className="bi-box-arrow-in-right px-1"></i>Register</Link></p>
      </Container>
    </Layout>
  }
};

export {
  ProtectedRoute
}
