import React from 'react'
import { render } from 'react-dom'
import { BrowserRouter } from 'react-router-dom'
import App from './App'

const mount = document.createElement('div')
document.body.appendChild(mount)

render(
  <BrowserRouter>
    <App />
  </BrowserRouter>,
  mount
)
