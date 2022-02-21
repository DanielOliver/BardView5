// @ts-nocheck
import React from 'react'
import { useQuery } from 'react-query'
import { Dnd5eSettingGet } from '../../../bv5-server'
import { bv5V1GetDnd5eSettings } from '../../../services/bardview5'
import { AxiosResponse } from 'axios'
import { Container, Row, Spinner, Table } from 'react-bootstrap'
import { CellProps, Column, useGlobalFilter, useSortBy, useTable } from 'react-table'
import { Link } from 'react-router-dom'
import { GlobalFilter } from '../../../components/Table'

function SettingTable ({
  data
}: {
  data: Dnd5eSettingGet[]
}) {
  const columns: Column<Dnd5eSettingGet>[] = React.useMemo(() =>
    [
      {
        Header: 'Name',
        accessor: 'name',
        Cell: (cell: CellProps<Dnd5eSettingGet>) => (
                              <Link to={`/dnd5e/settings/${cell.row.original.dnd5eSettingId}`}>{cell.value}</Link>)
      },
      {
        Header: 'Access',
        accessor: 'commonAccess'
      },
      {
        Header: 'Module',
        accessor: 'module'
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
  } = useTable<Dnd5eSettingGet>({
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

function Dnd5eSettingList () {
  const {
    data,
    error,
    isLoading
  } = useQuery<Dnd5eSettingGet[], AxiosResponse>('dnd5e-settings', async () => {
    const { data } = await bv5V1GetDnd5eSettings()
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

  return <SettingTable data={data}/>
}

export {
  Dnd5eSettingList
}
