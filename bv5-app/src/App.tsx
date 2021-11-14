import React from 'react'
import logo from './logo.svg'
import './App.css'
import { Link } from 'react-router-dom'

// import 'whatwg-fetch';

function App () {
  return (
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
          </div>
  )
}

export default App
