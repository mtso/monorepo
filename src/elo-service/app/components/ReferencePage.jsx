import React from 'react'
import docs from '../../API.md'
import marked from 'marked'

const markup = {
  __html: marked(docs),
}

const ReferencePage = () => (
  <div dangerouslySetInnerHTML={markup} />
)

export default ReferencePage
