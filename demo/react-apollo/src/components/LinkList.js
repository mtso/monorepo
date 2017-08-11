import React, { Component } from 'react'
import Link from './Link'
import { graphql, gql } from 'react-apollo'

class LinkList extends Component {
  render() {
    const { query } = this.props

    if (query && query.loading) {
      return (<div>Loading</div>)
    }

    if (query && query.error) {
      console.warn(query.error)
      return (<div>Error</div>)
    }

    const { allLinks } = query

    return (
      <div>
        {allLinks.map((link) => (
          <Link key={link.id} link={link} />
        ))}
      </div>
    )
  }
}

const ALL_LINKS_QUERY = gql`
  query AllLinksQuery {
    allLinks {
      id
      createdAt
      url
      description
    }
  }
`

export default graphql(
  ALL_LINKS_QUERY,
  {name: 'query'}
)(LinkList)
