import React, { useEffect, useState } from 'react'
import { Outlet } from 'react-router-dom'
import { AuthContext, AuthInitialState, AuthReducer } from './context/Auth.context'
import { getSelfServiceLogout, getSession } from './services/auth'
import { Layout } from './components/Layout'

function App () {
  const [state, dispatch] = React.useReducer(AuthReducer, AuthInitialState)
  const [checking, setChecking] = useState<'CHECKED' | 'CHECKING_SESSION' | 'UNCHECKED'>('UNCHECKED')

  const submitLogout = () => {
    getSelfServiceLogout().then(value => {
      dispatch({
        type: 'LOGOUT',
        isRegistrationComplete: false
      })
      if (value) {
        window.location.href = value.data.logout_url
      }
    })
  }

  useEffect(() => {
    if (checking === 'UNCHECKED') {
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
           <Layout logout={submitLogout}/>
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
