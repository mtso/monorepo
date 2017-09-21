import { h, render } from 'preact'
import { Router } from 'preact-router'
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

render(
  <Router>
    <App path='/' />
    <Hello path='/:name' />
  </Router>,
  document.querySelector('#app'),
  document.querySelector('#app')
)
