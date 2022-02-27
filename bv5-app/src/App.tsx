import React, { useEffect, useState } from 'react'
import { BrowserRouter, Route, Routes } from 'react-router-dom'
import { AuthContext, AuthInitialState, AuthReducer } from './context/Auth.context'
import { getSession } from './services/auth'
import { RouteHome } from './routes/home.route'
import { RouteRegister } from './routes/register.route'
import { RouteLogin } from './routes/login.route'
import { ProtectedRoute } from './components/ProtectedRoute'
import { RouteDnd5eSettingView } from './routes/dnd5e/settings/view.route'
import { RouteDnd5eSettingCreate } from './routes/dnd5e/settings/create.route'
import { RouteDnd5eSettingList } from './routes/dnd5e/settings/list.route'
import { RouteDnd5eSettingEdit } from './routes/dnd5e/settings/edit.route'
import { RouteProfile } from './routes/user/profile'
import { RouteDnd5eMonsterList } from './routes/dnd5e/monsters/list.route'

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
                  <Route path="/register" element={<RouteRegister isAuthenticated={state.isAuthenticated}/>}/>
                  <Route path="/login" element={<RouteLogin isAuthenticated={state.isAuthenticated}/>}/>
                  <Route path="/" element={<ProtectedRoute isAuthenticated={state.isAuthenticated}/>}>
                    <Route path="/" element={<RouteHome/>}/>
                    <Route path="/profile" element={<RouteProfile/>}/>
                    <Route path="/dnd5e/settings"
                           element={<RouteDnd5eSettingList/>}/>
                    <Route path="/dnd5e/settings/create"
                           element={<RouteDnd5eSettingCreate/>}/>
                    <Route path="/dnd5e/settings/:dnd5eSettingId/edit"
                           element={<RouteDnd5eSettingEdit/>}/>
                  </Route>
                  <Route path="/dnd5e/settings/:dnd5eSettingId/monsters"
                         element={<RouteDnd5eMonsterList isAuthenticated={state.isAuthenticated}/>}/>
                  <Route path="/dnd5e/settings/:dnd5eSettingId"
                         element={<RouteDnd5eSettingView isAuthenticated={state.isAuthenticated}/>}/>
                  <Route path="/*" element={<h1>Unknown</h1>}/>
                </Routes>
              </BrowserRouter>
            </div>
          </AuthContext.Provider>
  )
}

export default App
