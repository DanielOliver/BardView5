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
import LayoutSidebar from './components/LayoutSidebar'
import { Dnd5eSettingEdit } from './routes/dnd5e/settings/edit.route'
import Profile from './routes/user/profile'

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
                  <Route path="/register" element={<LayoutWrapper title="Register"><RegisterRoute/></LayoutWrapper>}/>
                  <Route path="/login" element={<LayoutWrapper title="Login"><LoginRoute/></LayoutWrapper>}/>
                  <Route path="/" element={<ProtectedRoute isAuthenticated={state.isAuthenticated}/>}>
                    <Route path="/" element={<LayoutSidebar title="Home"><HomeRoute/></LayoutSidebar>}/>
                    <Route path="/profile" element={<LayoutSidebar title="Profile"><Profile/></LayoutSidebar>}/>
                    <Route path="/dnd5e/settings"
                           element={<LayoutSidebar title="Settings D&D 5e"><Dnd5eSettingList/></LayoutSidebar>}/>
                    <Route path="/dnd5e/settings/create"
                           element={<LayoutSidebar title="New: Setting D&D 5e"><Dnd5eSettingCreate/></LayoutSidebar>}/>
                    <Route path="/dnd5e/settings/:dnd5eSettingId/edit"
                           element={<LayoutWrapper title="Edit: Setting D&D 5e"
                                                   isAuthenticated={state.isAuthenticated}><Dnd5eSettingEdit/></LayoutWrapper>}/>
                  </Route>
                  <Route path="/dnd5e/settings/:dnd5eSettingId"
                         element={<LayoutWrapper title="Setting D&D 5e"
                                                 isAuthenticated={state.isAuthenticated}>
                           <Dnd5eSettingView
                                   isAuthenticated={state.isAuthenticated}/>
                         </LayoutWrapper>}/>
                  <Route path="/*" element={<h1>Unknown</h1>}/>
                </Routes>
              </BrowserRouter>
            </div>
          </AuthContext.Provider>
  )
}

export default App
