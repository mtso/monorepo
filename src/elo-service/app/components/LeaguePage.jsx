import React, { Component } from 'react'
import { withRouter } from 'react-router-dom'

const LeaguePage = ({ league }) => (
  <div>
    <div>{league && league.title}</div>
  </div>
)

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
      />
    )
  }
}

export default withRouter(LeaguePageContainer)
