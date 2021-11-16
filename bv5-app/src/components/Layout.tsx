import { Icon, Menu } from 'semantic-ui-react'
import { Link } from 'react-router-dom'
import React from 'react'
import { AuthContext } from '../context/Auth.context'

export function Layout ({ logout }: {
  logout: () => void
}) {
  const { state } = React.useContext(AuthContext)

  return <Menu stackable>
    <Menu.Header as="h1">
      <Link to="/">Bardview5</Link>
    </Menu.Header>
    {state.isAuthenticated
      ? (<Menu.Menu position="right">
                      <Menu.Item> <Link onClick={logout} to='/'><Icon name="sign-out"/> Logout</Link> </Menu.Item>
                    </Menu.Menu>
        )
      : (<Menu.Menu position="right">
              <Menu.Item><Icon name="sign-in"/> Login</Menu.Item>
              <Menu.Item><Link to="/register"><Icon name="signup"/>Register</Link> </Menu.Item>
            </Menu.Menu>)}
  </Menu>
}
