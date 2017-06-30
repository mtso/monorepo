import path from 'path'
import Express from 'express'
import React from 'react'
import { renderToString } from 'react-dom/server'
import { StaticRouter } from 'react-router'
import App from './App'

const app = new Express()
app.set('view engine', 'ejs')
app.set('views', __dirname)
app.use(Express.static(path.resolve(__dirname, '../dist')))

app.get('*', (req, res) => {
  const context = {}

  const markup = renderToString(
    <StaticRouter
      location={req.url}
      context={context}
    >
      <App/>
    </StaticRouter>
  )

  if (context.url) {
    res.redirect(301, context.url)
  } else {
    res.render('index', { markup })
  }
})

const port = process.env.PORT || 3750
const env = process.env.NODE_ENV || 'production'
app.listen(port, err => {
  if (err) {
    return console.error(err)
  }
  console.info(`[${env}] listening on ${port}`)
})