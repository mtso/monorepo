import { h, Component } from 'preact'
import { Router, Link, route } from 'preact-router'

export const Hello = ({name}) => {
  console.log(name)
  return (
  <div>
    Hello, {name}
  </div>
)
}

class RedirectingButton extends Component {
  constructor() {
    super()
    this.go = this.go.bind(this)
  }

  go(e) {
    e.preventDefault()
    console.log('GO!!')
  }

  render() {
    return (<button onClick={this.go}>what</button>)
  }
}

const Main = () => {
  return (
  <div>
    hi~
    <RedirectingButton />
  </div>
)
}

const App = () => (
  <Main />
)

export default App

