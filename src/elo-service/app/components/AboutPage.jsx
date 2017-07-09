import React from 'react'
import readme from '../../README.md'
import marked from 'marked'
import { Link } from 'react-router-dom'

const markup = {
  __html: marked(readme),
}

const AboutPage = () => (
  <div>
    <Link to='/'>ELO</Link>
    <div dangerouslySetInnerHTML={markup} />
  </div>
)

export default AboutPage
