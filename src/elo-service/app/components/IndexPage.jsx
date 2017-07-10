import React, { Component } from 'react'
import { Link, Redirect, withRouter } from 'react-router-dom'
import request from 'superagent'

const IndexPage = ({ league, onCreate, ...props }) => {
  if (!!league) {
    return (
      <Redirect
        push
        to={{
          pathname: '/' + league.id,
          state: { league },
        }}
      />
    )
  }

  return (
    <div>
      <h1>ELO</h1>
      <form onSubmit={onCreate}>
        <input
          type='text'
          name='title'
          placeholder='League Title'
          className='textfield'
        />
        <input
          type='submit'
          value='Create League'
          className='action-button'
        />
      </form>
      <span>© 2017 <a href='https://github.com/mtso'>
        mtso
      </a> | <Link to='/api'>
        API Reference
      </Link> | <Link to='/about'>
        About
      </Link>
      </span>
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
