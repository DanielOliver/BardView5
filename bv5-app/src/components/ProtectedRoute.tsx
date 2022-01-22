import { Outlet, RouteProps } from 'react-router-dom'
import React from 'react'
import LayoutWrapper from './LayoutWrapper'
import { Container } from 'react-bootstrap'

export type ProtectedRouteProps = {
  isAuthenticated: boolean;
} & RouteProps;

// Originally https://stackoverflow.com/a/47754325
export default function ProtectedRoute ({
  isAuthenticated,
  ...routeProps
}: ProtectedRouteProps) {
  if (isAuthenticated) {
    return <Outlet/>
  } else {
    return <LayoutWrapper>
      <Container fluid="lg">
        <p>A self-hosted Fifth Edition Dungeons & Dragons Campaign Management Tool.</p>
        <br/>
        <p>Go ahead and login. Adventure awaits!</p>
      </Container>
    </LayoutWrapper>
  }
};
