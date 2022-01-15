import { Dropdown, Icon, Menu } from 'semantic-ui-react'
import { Link } from 'react-router-dom'
import React from 'react'
import { AuthContext } from '../context/Auth.context'
import { useQuery } from 'react-query'
import { UserGet } from '../bv5-server'
import { bv5V1GetMe } from '../services/bardview5'

export function Layout ({ logout }: {
  logout: () => void
}) {
  const { state } = React.useContext(AuthContext)

  const {
    data
  } = useQuery<UserGet>('me', async () => {
    return (await bv5V1GetMe()).data
  })

  return <Menu stackable>
    <Menu.Header as="h1">
      <Link to="/">Bardview5</Link>
    </Menu.Header>
    {state.isAuthenticated
      ? (<Menu.Menu position="right">
                      <Dropdown item icon="user" text={data?.name}>
                        <Dropdown.Menu>
                          <Dropdown.Item icon="edit" text="Profile"/>
                          <Dropdown.Item icon="sign-out" text="Logout" onClick={logout}/>
                        </Dropdown.Menu>
                      </Dropdown>
                    </Menu.Menu>
        )
      : (<Menu.Menu position="right">
              <Menu.Item><Link to="/login"> <Icon name="sign-in"/> Login</Link></Menu.Item>
              <Menu.Item><Link to="/register"><Icon name="signup"/>Register</Link> </Menu.Item>
            </Menu.Menu>)}
  </Menu>
}
