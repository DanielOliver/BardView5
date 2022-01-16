import React from 'react'
import ReactDOM from 'react-dom'
import App from './App'
import reportWebVitals from './reportWebVitals'
import { QueryClient, QueryClientProvider } from 'react-query'
import { BrowserRouter, Route, Routes } from 'react-router-dom'
import RegisterRoute from './routes/register.route'

// import 'semantic-ui-css/semantic.min.css'
import './App.scss'
import LoginRoute from './routes/login.route'
import HomeRoute from './routes/home.route'

const queryClient = new QueryClient()

ReactDOM.render(
        <React.StrictMode>
          <QueryClientProvider client={queryClient}>
            <BrowserRouter>
              <Routes>
                <Route path="/" element={<App/>}>
                  <Route path="/" element={<HomeRoute/>}/>
                  <Route path="/register" element={<RegisterRoute/>}/>
                  <Route path="/login" element={<LoginRoute/>}/>
                </Route>
              </Routes>
            </BrowserRouter>
          </QueryClientProvider>
        </React.StrictMode>,
        document.getElementById('root')
)

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals()
