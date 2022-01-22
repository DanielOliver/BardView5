import React, { useEffect, useState } from 'react'
import { BrowserRouter, Route, Routes } from 'react-router-dom'
import { AuthContext, AuthInitialState, AuthReducer } from './context/Auth.context'
import { getSession } from './services/auth'
import HomeRoute from './routes/home.route'
import Dnd5eWorldView from './routes/dnd5e/worlds/view.route'
import RegisterRoute from './routes/register.route'
import LoginRoute from './routes/login.route'
import LayoutWrapper from './components/LayoutWrapper'
import ProtectedRoute from './components/ProtectedRoute'

function App () {
  const [state, dispatch] = React.useReducer(AuthReducer, AuthInitialState)
  const [checking, setChecking] = useState<'CHECKED' | 'CHECKING_SESSION' | 'UNCHECKED'>('UNCHECKED')

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
            <div className="App">
              <BrowserRouter>
                <Routes>
                  <Route path="/register" element={<LayoutWrapper><RegisterRoute/></LayoutWrapper>}/>
                  <Route path="/login" element={<LayoutWrapper><LoginRoute/></LayoutWrapper>}/>
                  <Route path="/" element={<ProtectedRoute isAuthenticated={state.isAuthenticated}/>}>
                    <Route path="/" element={<LayoutWrapper><HomeRoute/> </LayoutWrapper>}/>
                    <Route path="/dnd5e/worlds/:dnd5eWorldId"
                           element={<LayoutWrapper><Dnd5eWorldView/></LayoutWrapper>}/>
                    <Route path="/*" element={<h1>Unknown</h1>}/>
                  </Route>
              </Routes>
            </BrowserRouter>
          </div>
</AuthContext.Provider>
  )
}

export default App
