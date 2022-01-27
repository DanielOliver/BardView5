import React, { useEffect, useState } from 'react'
import { BrowserRouter, Route, Routes } from 'react-router-dom'
import { AuthContext, AuthInitialState, AuthReducer } from './context/Auth.context'
import { getSession } from './services/auth'
import HomeRoute from './routes/home.route'
import RegisterRoute from './routes/register.route'
import LoginRoute from './routes/login.route'
import LayoutWrapper from './components/LayoutWrapper'
import ProtectedRoute from './components/ProtectedRoute'
import { Dnd5eSettingView } from './routes/dnd5e/settings/view.route'
import { Dnd5eSettingCreate } from './routes/dnd5e/settings/create.route'
import { Dnd5eSettingList } from './routes/dnd5e/settings/list.route'

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
                    <Route path="/dnd5e/settings/"
                           element={<LayoutWrapper><Dnd5eSettingList/></LayoutWrapper>}/>
                    <Route path="/dnd5e/settings/create"
                           element={<LayoutWrapper><Dnd5eSettingCreate/></LayoutWrapper>}/>
                  </Route>
                  <Route path="/dnd5e/settings/:dnd5eSettingId"
                         element={<LayoutWrapper><Dnd5eSettingView/></LayoutWrapper>}/>
                  <Route path="/*" element={<h1>Unknown</h1>}/>
              </Routes>
            </BrowserRouter>
          </div>
</AuthContext.Provider>
  )
}

export default App
