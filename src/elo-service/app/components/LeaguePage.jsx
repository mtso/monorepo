import React, { Component } from 'react'
import { Link, Route, withRouter } from 'react-router-dom'
import request from 'superagent'
import path from 'path'

const Players = ({ players }) => (
  <table>
    <thead>
      <tr>
        <th>Name</th>
        <th>ELO</th>
      </tr>
    </thead>
    <tbody>
    {players && players.map(({ username, elo }) => (
      <tr key={username}>
        <td>{username}</td>
        <td>{elo}</td>
      </tr>
    ))}
    </tbody>
  </table>
)

const Games = ({ games }) => (
  <table>
    <thead>
      <tr>
        <th>Winner</th>
        <th>Loser</th>
      </tr>
    </thead>
    <tbody>
    {games && games.map(({ winner, loser }, i) => (
      <tr key={i}>
        <td>{winner.username}</td>
        <td>{loser.username}</td>
      </tr>
    ))}
    </tbody>
  </table>
)

class GamesContainer extends Component {
  constructor(props) {
    super(props)
    const { games } = this.props
    this.state = { games }
  }
  componentDidMount() {
    const { league } = this.props
    if (!league) { return }
    request
      .get(`/api/${league.id}/games`)
      .then(({ body }) => body)
      .then(({ ok, games, message }) => {
        if (!ok) {
          throw new Error(message)
        }
        this.setState({
          games,
        })
      })
  }
  render() {
    return (
      <Games
        {...this.state} 
        {...this.props}
      />
    )
  }
}

class PlayersContainer extends Component {
  constructor(props) {
    super(props)
    const { players } = this.props
    this.state = {
      players,
    }
  }
  componentDidMount() {
    const { league } = this.props
    if (!league) {
      return
    }
    request
      .get(`/api/${league.id}/players`)
      .then(({ body }) => body)
      .then(({ ok, players, message }) => {
        if (!ok) {
          throw new Error(message)
        }
        this.setState({
          players,
        })
      })
  }
  render() {
    return (
      <Players
        {...this.state}
        {...this.props}
      />
    )
  }
}

const RoutedPlayers = withRouter(PlayersContainer)

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
        component={() => (<PlayersContainer league={league} />)}
      />
      <Route
        path={path.join(match.url, 'games')}
        component={() => (<GamesContainer league={league} />)}
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
