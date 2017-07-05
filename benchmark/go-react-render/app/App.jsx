import React from 'react'
import { Route } from 'react-router'

const Greeting = ({ name }) => (
  <h1>Hello~ {name}!</h1>
)

const App = ({ state }) => (
  <div>
    <Greeting {...state} />
    <Route
      path='/'
      render={({ location }) => (
        <span>Location: {location.pathname}</span>
      )}
    />
  </div>
)

export default App
