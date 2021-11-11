import React from 'react';
import logo from './logo.svg';
import './App.css';
import {useQuery} from "react-query";
import {startSelfServiceRegister} from "./services/auth";
import {SelfServiceRegistrationFlow} from "@ory/kratos-client";

// import 'whatwg-fetch';


function App() {
  const {isLoading, error, data} = useQuery<SelfServiceRegistrationFlow>('login?', async () => {
    return await startSelfServiceRegister();
  })

  if (isLoading) return <p>'Loading...'</p>

  if (error) return <pre>An error has occurred: {JSON.stringify(error, null, 2)}</pre>

  return <pre>
    {JSON.stringify(data, null, 2)}
  </pre>

  return (
          <div className="App">
            <header className="App-header">
              <img src={logo} className="App-logo" alt="logo"/>
              <p>
                Edit <code>src/App.tsx</code> and save to reload.
              </p>
              <a
                      className="App-link"
                      href="https://reactjs.org"
                      target="_blank"
                      rel="noopener noreferrer"
              >
                Learn React
              </a>
            </header>
          </div>
  );
}

export default App;
