import React, { Component } from 'react'

class Link extends Component {
  render() {
    const { description, url } = this.props.link
    return (
      <div>
        <div>{description} ({url})</div>
      </div>
    )
  }

  _voteForLink = async () => {

  }
}

export default Link
