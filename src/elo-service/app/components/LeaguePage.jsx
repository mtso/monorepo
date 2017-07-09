import React, { Component } from 'react'
import { Link, Route, withRouter } from 'react-router-dom'
import request from 'superagent'
import path from 'path'

const LeaguePage = ({ league, match, history, ...props }) => (
  <div>
    <div className='titlebar'>
      <div>{league && league.title}</div>
      <div>{league && league.id}</div>
    </div>
    <div>
      <Link to={{
        pathname: path.join(match.url, 'players'),
        state: { league },
      }}>Leaderboard</Link>
      <Link to={{
        pathname: path.join(match.url, 'games'),
        state: { league },
      }}>Game History</Link>

      <Route
        path={path.join(match.url, 'players')}
        render={() => (<div>player1</div>)}
      />
      <Route
        path={path.join(match.url, 'games')}
        render={() => (<div>game history</div>)}
      />
    </div>
  </div>
)
      // <button onClick={() => history.push(path.join(match.url, 'players'))}>Leaderboard</button>
      // <button onClick={() => history.push(path.join(match.url, 'games'))}>Game History</button>
      // <Link to={path.join(match.url, 'players')}>Leaderboard</Link>
      // <Link to={path.join(match.url, 'games')}>Game History</Link>

class LeaguePageContainer extends Component {
  constructor(props) {
    super(props)
    this.state = {
      league: null,
    }
  }
  componentDidMount() {
    const { location, match } = this.props
    const { state } = location
    if (!!state) {
      return this.setState({
        ...state,
      })
    }
    // Check if component re-mounts on URL path change.
    // console.log('mounting...')

    const { params } = match
    const { id } = params
    request
      .get('/api/'+id)
      .then(({ body }) => body)
      .then(({ ok, league, message }) => {
        if (!ok) {
          throw new Error(message)
        }
        this.setState({
          league,
        })
      })
  }
  render() {
    return (
      <LeaguePage
        {...this.state}
        {...this.props}
      />
    )
  }
}

export default withRouter(LeaguePageContainer)
