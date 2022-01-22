import React from 'react'
import { Layout } from './Layout'

const LayoutWrapper: React.FC<{ }> = ({
  children
}) => (
        <>
          <Layout/>
          {children}
        </>
)

export default LayoutWrapper
