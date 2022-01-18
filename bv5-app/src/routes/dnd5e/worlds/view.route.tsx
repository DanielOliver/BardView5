import React from 'react'
import { useParams } from 'react-router-dom'
import { useQuery } from 'react-query'
import { Dnd5eWorldGet } from '../../../bv5-server'
import { bv5V1GetDnd5eWorld } from '../../../services/bardview5'
import { AxiosResponse } from 'axios'
import { Container } from 'react-bootstrap'

function Dnd5eWorldView () {
  const params = useParams()
  const dnd5eWorldId: string = params.dnd5eWorldId ?? '0'
  const {
    data,
    error
  } = useQuery<Dnd5eWorldGet[], AxiosResponse>(`dnd5e-world-${dnd5eWorldId}`, async () => {
    const { data } = await bv5V1GetDnd5eWorld(dnd5eWorldId)
    return data
  }, {
    retry: false
  })

  if (error) {
    return <Container fluid="lg">
      <p>Did not find world {dnd5eWorldId}</p>
      <pre>
        {JSON.stringify(error, null, 2)}
      </pre>
    </Container>
  }

  return <Container fluid="lg">
    <p>World data</p>
    <pre>
    {JSON.stringify(data, null, 2)}
  </pre>
  </Container>
}

export default Dnd5eWorldView
