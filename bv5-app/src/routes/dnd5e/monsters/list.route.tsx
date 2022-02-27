// @ts-nocheck
import React from 'react'
import { useQuery } from 'react-query'
import { Dnd5eMonsterGet } from '../../../bv5-server'
import { bv5V1GetDnd5eMonstersBySetting } from '../../../services/bardview5'
import { AxiosResponse } from 'axios'
import { Breadcrumb, Container, Row, Spinner, Table } from 'react-bootstrap'
import { Column, useGlobalFilter, useSortBy, useTable } from 'react-table'
import { GlobalFilter } from '../../../components/Table'
import { Link, useParams } from 'react-router-dom'
import { Bv5RouteProps } from '../../../components/Common'
import LayoutWrapper from '../../../components/LayoutWrapper'

function MonsterTable ({
  data
}: {
  data: Dnd5eMonsterGet[]
}) {
  const columns: Column<Dnd5eMonsterGet>[] = React.useMemo(() =>
    [
      {
        Header: 'Name',
        accessor: 'name'
      },
      {
        Header: 'Type',
        accessor: 'monsterType'
      },
      {
        Header: 'AC',
        accessor: 'armorClass'
      },
      {
        Header: 'HP',
        accessor: 'hitPoints'
      },
      {
        Header: 'Size',
        accessor: 'sizeCategory'
      },
      {
        Header: 'Alignment',
        accessor: 'alignment'
      }
    ],
  []
  )

  const {
    getTableProps,
    getTableBodyProps,
    headerGroups,
    rows,
    prepareRow,
    state,
    visibleColumns,
    preGlobalFilteredRows,
    setGlobalFilter
  } = useTable<Dnd5eMonsterGet>({
    columns,
    data
  },
  useGlobalFilter,
  useSortBy)

  return <Container fluid="lg">
    <Row>
      <Table {...getTableProps()}>
        <thead>
        <tr>
          <th
                  colSpan={visibleColumns.length}
                  style={{
                    textAlign: 'left'
                  }}
          >
            <GlobalFilter
                    preGlobalFilteredRows={preGlobalFilteredRows}
                    globalFilter={state.globalFilter}
                    setGlobalFilter={setGlobalFilter}
            />
          </th>
        </tr>
        {headerGroups.map(headerGroups => (
                // eslint-disable-next-line react/jsx-key
                <tr {...headerGroups.getHeaderGroupProps()}>
                  {
                    headerGroups.headers.map(column => (
                            // eslint-disable-next-line react/jsx-key
                            <th {...column.getHeaderProps(column.getSortByToggleProps())}>
                              {column.render('Header')}
                              <span>
                    {column.isSorted
                      ? column.isSortedDesc
                        ? ' 🔽'
                        : ' 🔼'
                      : ''}
                  </span>
                            </th>
                    ))}
                </tr>
        ))}
        </thead>
        <tbody {...getTableBodyProps()}>
        {
          rows.map(row => {
            prepareRow(row)
            return (
                    // eslint-disable-next-line react/jsx-key
                    <tr {...row.getRowProps()}>
                      {row.cells.map(cell => {
                        // eslint-disable-next-line react/jsx-key
                        return (<td {...cell.getCellProps()}>
                                  {cell.render('Cell')}
                                </td>
                        )
                      })}
                    </tr>
            )
          })}

        </tbody>
      </Table>
    </Row>
  </Container>
}

function Dnd5eMonsterList () {
  const params = useParams()
  const dnd5eSettingId: string = params.dnd5eSettingId ?? '0'

  const {
    data,
    error,
    isLoading
  } = useQuery<Dnd5eMonsterGet[], AxiosResponse>(`dnd5e-monsters-${dnd5eSettingId}`, async () => {
    const { data } = await bv5V1GetDnd5eMonstersBySetting(dnd5eSettingId)
    return data
  }, {
    retry: false
  })

  if (error) {
    return <Container fluid="lg">
      <pre>
        {JSON.stringify(error, null, 2)}
      </pre>
    </Container>
  }

  if (isLoading || data === undefined) {
    return <Container fluid="lg">
      <Spinner animation="border" role="status">
        <span className="visually-hidden">Loading...</span>
      </Spinner>
    </Container>
  }

  return <div>
    <Breadcrumb>
      <Breadcrumb.Item>
        <Link to={`/dnd5e/settings/${dnd5eSettingId}`}>Setting</Link>
      </Breadcrumb.Item>
      <Breadcrumb.Item active>Monsters</Breadcrumb.Item>
    </Breadcrumb>

    <MonsterTable data={data}/>
  </div>
}

const RouteDnd5eMonsterList = ({
  isAuthenticated
}: Bv5RouteProps) => {
  return <LayoutWrapper title="Monsters D&D 5e"
                        isAuthenticated={isAuthenticated}>
    <Dnd5eMonsterList/>
  </LayoutWrapper>
}

export {
  RouteDnd5eMonsterList
}
