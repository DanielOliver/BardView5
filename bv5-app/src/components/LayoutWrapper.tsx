import React from 'react'
import { Layout } from './Layout'
import { Container, Row } from 'react-bootstrap'
import LayoutSidebar from './LayoutSidebar'

const LayoutWrapper: React.FC<{isAuthenticated?: boolean}> = ({
  children,
  isAuthenticated = false
}) => {
  if (isAuthenticated) {
    return <LayoutSidebar>
      {children}
    </LayoutSidebar>
  }

  return <Container fluid>
      <Row> <Layout/></Row>
      <Row>
        {children}
      </Row>
    </Container>
}

export default LayoutWrapper
