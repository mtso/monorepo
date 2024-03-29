import React from 'react'
import docs from '../../API.md'
import marked from 'marked'
import { Link } from 'react-router-dom'

const markup = {
  __html: marked(docs),
}

const ReferencePage = () => (
  <div>
    <Link to='/'>ELO</Link>
    <div dangerouslySetInnerHTML={markup} />
  </div>
)

export default ReferencePage
