import React from 'react'
import { render } from 'react-dom'
import { BrowserRouter } from 'react-router-dom'
import App from './App'

const BasicExample = () => (
  <BrowserRouter>
    <App/>
  </BrowserRouter>
)

render(
  <BasicExample/>,
  document.getElementById('app')
)
