import React from 'react'

interface AuthState {
  isAuthenticated: boolean,
  isRegistrationComplete: boolean,
  checked: 'UNCHECKED' | 'CHECKED'
}

const AuthInitialState: AuthState = {
  isAuthenticated: localStorage.login === 'true',
  isRegistrationComplete: false,
  checked: 'UNCHECKED'
}

interface AuthReducerAction {
  type: 'LOGIN' | 'LOGOUT',
  isRegistrationComplete: boolean,
  checked: | 'UNCHECKED' | 'CHECKED'
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
    case 'LOGIN':
      localStorage.setItem('login', 'true')
      return {
        ...state,
        isAuthenticated: true,
        isRegistrationComplete: action.isRegistrationComplete,
        checked: action.checked
      }
    case 'LOGOUT':
      localStorage.clear()
      return {
        ...state,
        isAuthenticated: false,
        isRegistrationComplete: false,
        checked: action.checked
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
