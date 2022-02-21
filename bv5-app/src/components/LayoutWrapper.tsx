import React from 'react'
import { Layout } from './Layout'
import { Container, Row } from 'react-bootstrap'
import LayoutSidebar from './LayoutSidebar'

const LayoutWrapper: React.FC<{isAuthenticated?: boolean, title?: string}> = ({
  children,
  title,
  isAuthenticated = false
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
