import React from 'react'

interface AuthState {
  isAuthenticated: boolean,
  isRegistrationComplete: boolean,
  checked: 'UNCHECKED' | 'CHECKED'
}

const AuthInitialState: AuthState = {
  isAuthenticated: localStorage.login === 'true',
  isRegistrationComplete: localStorage.registrationComplete === 'false',
  checked: 'UNCHECKED'
}

interface AuthReducerAction {
  type: 'LOGIN' | 'LOGOUT' | 'CHECK',
  isRegistrationComplete: boolean
  // checked: | 'UNCHECKED' | 'CHECKED'
}

interface ContextDispatch<T, TReducer> {
  state: T,
  dispatch: (reduce: TReducer) => void
}

const AuthContext = React.createContext<ContextDispatch<AuthState, AuthReducerAction>>({
  state: AuthInitialState,
  dispatch: reduce => {}
})

function AuthReducer (state: AuthState, action: AuthReducerAction): AuthState {
  switch (action.type) {
    case 'CHECK':
      return {
        ...state,
        isAuthenticated: localStorage.login === 'true',
        isRegistrationComplete: localStorage.registrationComplete === 'false',
        checked: 'CHECKED'
      }
    case 'LOGIN':
      localStorage.registrationComplete = String(action.isRegistrationComplete)
      localStorage.login = 'true'
      return {
        ...state,
        isAuthenticated: true,
        isRegistrationComplete: action.isRegistrationComplete,
        checked: 'CHECKED'
      }
    case 'LOGOUT':
      localStorage.clear()
      return {
        ...state,
        isAuthenticated: false,
        isRegistrationComplete: false,
        checked: 'CHECKED'
      }
    default:
      return state
  }
}

export {
  AuthReducer,
  AuthContext,
  AuthInitialState
}
