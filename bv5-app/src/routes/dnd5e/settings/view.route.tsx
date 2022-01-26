import React from 'react'
import { useParams } from 'react-router-dom'
import { useQuery } from 'react-query'
import { Dnd5eSettingGet } from '../../../bv5-server'
import { bv5V1GetDnd5eSetting } from '../../../services/bardview5'
import { AxiosResponse } from 'axios'
import { Container } from 'react-bootstrap'

function Dnd5eSettingView () {
  const params = useParams()
  const dnd5eSettingId: string = params.dnd5eSettingId ?? '0'
  const {
    data,
    error
  } = useQuery<Dnd5eSettingGet[], AxiosResponse>(`dnd5e-setting-${dnd5eSettingId}`, async () => {
    const { data } = await bv5V1GetDnd5eSetting(dnd5eSettingId)
    return data
  }, {
    retry: false
  })

  if (error) {
    return <Container fluid="lg">
      <p>Did not find setting {dnd5eSettingId}</p>
      <pre>
        {JSON.stringify(error, null, 2)}
      </pre>
    </Container>
  }

  return <Container fluid="lg">
    <p>Setting data</p>
    <pre>
    {JSON.stringify(data, null, 2)}
  </pre>
  </Container>
}

export {
  Dnd5eSettingView
}
