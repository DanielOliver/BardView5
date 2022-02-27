import React from 'react'
import { Layout } from './Layout'
import { Container, Row } from 'react-bootstrap'
import LayoutSidebar from './LayoutSidebar'
import { Bv5RouteProps } from './Common'

const LayoutWrapper: React.FC<{ title?: string } & Bv5RouteProps> = ({
  children,
  title,
  isAuthenticated
}) => {
  if (isAuthenticated) {
    return <LayoutSidebar title={title}>
      {children}
    </LayoutSidebar>
  }

  return <Container fluid>
    <Row> <Layout title={title}/></Row>
    <Row>
      {children}
    </Row>
  </Container>
}

export default LayoutWrapper
