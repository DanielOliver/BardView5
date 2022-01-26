import React from 'react'
import { Badge } from 'react-bootstrap'

export function AccessBadge ({
  accessType
}: {
  accessType: string | undefined
}) {
  switch (accessType) {
    case 'private':
      return <Badge bg="info">Private</Badge>
    case 'anyuser':
      return <Badge bg="info">Any User</Badge>
    case 'public':
      return <Badge bg="info">Public</Badge>
    default:
      return <Badge bg="danger">Unknown</Badge>
  }
}
