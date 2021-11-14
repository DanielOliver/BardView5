import React, { useEffect, useState } from 'react'
import logo from './logo.svg'
import './App.css'
import { Link, Outlet } from 'react-router-dom'
import { AuthContext, AuthInitialState, AuthReducer } from './context/Auth.context'
import { getSession } from './services/auth'

// import 'whatwg-fetch';

function App () {
  const [state, dispatch] = React.useReducer(AuthReducer, AuthInitialState)
  const [checking, setChecking] = useState<'CHECKED' | 'UNCHECKED' | 'CHECKING'>('UNCHECKED')
  useEffect(() => {
    if (!state.isAuthenticated && state.checked === 'UNCHECKED' && checking === 'UNCHECKED') {
      setChecking('CHECKING')
      getSession().then(
        value => {
          if (value) {
            dispatch({
              type: 'LOGIN',
              isRegistrationComplete: false,
              checked: 'CHECKED'
            })
          } else {
            dispatch({
              type: 'LOGOUT',
              isRegistrationComplete: false,
              checked: 'CHECKED'
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
            <div className="App">
              <header className="App-header">
                <img src={logo} className="App-logo" alt="logo"/>
                <h1>Bardview5</h1>
                <nav
                        style={{
                          borderBottom: 'solid 1px',
                          paddingBottom: '1rem'
                        }}
                >
                  <Link to="/register">Register</Link>
                </nav>
              </header>
              <Outlet/>
              {state.isAuthenticated && <p>
                WASSSSSSUUUPPP!!!
              </p>}
            </div>
          </AuthContext.Provider>
  )
}

export default App
