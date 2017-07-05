import React from 'react'
import { StaticRouter } from 'react-router'
import { renderToString } from 'react-dom/server'
import App from '../app'

const renderMarkup = ({ url, state }) => {
  const ctx = {}
  const output = {}
  
  const markup = renderToString(
    <StaticRouter
      context={ctx}
      location={url}
    >
      <App state={state} />
    </StaticRouter>
  )

  const redirectUrl = (!!ctx.url) ? ctx.url : null

  return {
    markup,
    redirectUrl,
  }
}

const test = renderMarkup({
  state: { name: 'world' },
  url: '/',
})

console.log(test)
