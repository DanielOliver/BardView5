import React, { useEffect, useState } from 'react'
import './App.css'
import { Link, Outlet } from 'react-router-dom'
import { AuthContext, AuthInitialState, AuthReducer } from './context/Auth.context'
import { getSession } from './services/auth'
import { Icon, Menu } from 'semantic-ui-react'

// import 'whatwg-fetch';

function App () {
  const [state, dispatch] = React.useReducer(AuthReducer, AuthInitialState)
  const [checking, setChecking] = useState<'CHECKED' | 'CHECKING_SESSION' | 'CHECKED_STORAGE' | 'UNCHECKED' | 'CHECKING_STORAGE'>('UNCHECKED')
  useEffect(() => {
    if (!state.isAuthenticated && state.checked === 'UNCHECKED' && checking === 'UNCHECKED') {
      setChecking('CHECKING_STORAGE')
      dispatch({
        type: 'CHECK',
        isRegistrationComplete: false
      })
    }
    if (!state.isAuthenticated && state.checked === 'CHECKED' && checking === 'CHECKING_STORAGE') {
      setChecking('CHECKED_STORAGE')
    }
    if (!state.isAuthenticated && checking === 'CHECKED_STORAGE') {
      setChecking('CHECKING_SESSION')
      getSession().then(
        value => {
          if (value) {
            dispatch({
              type: 'LOGIN',
              isRegistrationComplete: false
            })
          } else {
            dispatch({
              type: 'LOGOUT',
              isRegistrationComplete: false
            })
          }
          setChecking('CHECKED')
        }
      )
    }
  }, [state, checking])

  return (

          <AuthContext.Provider
                  value={{
                    state,
                    dispatch
                  }}
          >
            <Menu stackable>
              <Menu.Header as="h1">
                <Link to="/">Bardview5</Link>
              </Menu.Header>
              {!state.isAuthenticated && <>
                <Menu.Item><Icon name='sign-in' /> Login</Menu.Item>
                <Menu.Item><Link to="/register"><Icon name='signup'/>Register</Link> </Menu.Item>
              </>}
            </Menu>

            <div className="App">
              <Outlet/>
              {state.isAuthenticated && <p>
                WASSSSSSUUUPPP!!!
              </p>}
            </div>
          </AuthContext.Provider>
  )
}

export default App
