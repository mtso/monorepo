const express = require('express')
const webpack = require('webpack')
let renderMarkup // = require('./dist/render')

const webpackConfig = require('./webpack.config')
const app = module.exports = express()

app.use(express.static('dist'))

app.get('/*', (req, res) => {
  // res.send('hi~')
  // const markup = renderMarkup.default({})
  // res.send(JSON.stringify(renderMarkup))
  res.status(200).send(renderIndex(renderMarkup(req.path)))
})

function renderIndex(markup) {
  return `
  <!doctype html>
  <html>
    <body>
      <div id="app">${markup}</div>
    
    <script src='/bundle.js'></script>
    </body>
  </html>
  `
}

webpack(webpackConfig)
  .run((err) => {
    if (err) { throw err }

    renderMarkup = require('./dist/render.js')

    const listener = app.listen(process.env.PORT || 3750, () => {
      console.log('Listening on', listener.address().port)
    })
  })
