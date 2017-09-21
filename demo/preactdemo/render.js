import { Router } from 'preact-router'
import { h } from 'preact'
import render from 'preact-render-to-string'
// import App, { Hello } from './app'
import App from './app'
/** @jsx h */

// module.exports = function(state) {
//   return render(
//     <App />
//   )
// }
const Hello = ({name}) => {
  console.log(name)
  return (
  <div>
    Hello, {name}
  </div>
)
}

const renderMarkup = function(url, state) {
  const serverHistory = {
    getCurrentLocation() {
      return url
    }
  }

  return render(
    <Router url={url} history={serverHistory}>
      <App path='/' />
      <Hello path='/:name' />
    </Router>
  )
}

module.exports = renderMarkup
