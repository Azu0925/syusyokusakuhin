import React from 'react'
import './App.css'
import Router from './routes/Router'
import { RecoilRoot } from 'recoil'

function App() {
  return (
    <div className="App">
      <RecoilRoot>
        <Router />
      </RecoilRoot>
    </div>
  )
}

export default App
