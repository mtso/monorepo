import React from 'react'
import { Route, withRouter } from 'react-router-dom'

const App = ({ match, location }) => {
  const { url } = match
  return (
    <div>
      <p>hello~</p>
      <pre>url = {url}</pre>
      <pre>match = {JSON.stringify(match, null, 2)}</pre>
      <pre>location = {JSON.stringify(location, null, 2)}</pre>
    </div>
  )
}

const RoutedApp = (props) => (
  <Route
    path='/:param'
    component={() => (
      <App
        {...props}
      />
    )}
  />
)

export default withRouter(RoutedApp)
