import React, { Component } from 'react'
import { Route, Switch, withRouter } from 'react-router-dom'
import IndexPage from './components/IndexPage'
import AboutPage from './components/AboutPage'
import LeaguePage from './components/LeaguePage'
import ReferencePage from './components/ReferencePage'

const App = ({ setLeague }) => (
  <Switch>
    <Route exact path='/' component={() => (
      <IndexPage
        setLeague={setLeague}
      />
    )} />
    <Route path='/api' component={ReferencePage} />
    <Route path='/about' component={AboutPage} />
    <Route path='/:id' component={() => (
      <LeaguePage />
    )} />
  </Switch>
)

class AppContainer extends Component {
  constructor(props) {
    super(props)
  }
  render() {
    return (
      <App
        {...this.state}
        setLeague={this.setLeague}
      />
    )
  }
}

export default withRouter(AppContainer)
