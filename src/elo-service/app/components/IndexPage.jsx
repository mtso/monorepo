import React, { Component } from 'react'
import { Redirect, withRouter } from 'react-router-dom'
import request from 'superagent'

const IndexPage = ({ league, onCreate, ...props }) => {
  if (!!league) {
    return (
      <Redirect to={{
        pathname: '/' + league.id,
        state: { league },
      }} />
    )
  }

  return (
    <div>
      <form onSubmit={onCreate}>
        <input
          type='text'
          name='title'
          placeholder='League Title'
        />
        <input
          type='submit'
          value='Create League'
        />
      </form>
    </div>
  )
}

class IndexPageContainer extends Component {
  constructor(props) {
    super(props)
    this.state = {
      league: null,
    }
    this.onCreate = this.onCreate.bind(this)
  }

  onCreate(e) {
    e.preventDefault()
    const inputs = e.target.elements
    const title = inputs["title"].value

    request
      .post('/api/new')
      .send({ title })
      .then(({ body }) => body)
      .then(({ ok, league, message }) => {
        if (!ok) {
          throw new Error(message)
        }
        this.setState({
          league,
        })
      })
      .catch(console.error)
  }

  render() {
    return (
      <IndexPage
        {...this.state}
        {...this.props}
        onCreate={this.onCreate}
      />
    )
  }
}

export default withRouter(IndexPageContainer)
