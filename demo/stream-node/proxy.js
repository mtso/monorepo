const http = require('http')
const request = require('request')

const app = module.exports = http.createServer((req, res) => {
  req.pipe(request.get({
      url: 'http://localhost:3750',
      headers: {test: 'foo-bar'}, //Object.assign({}, req.headers, {test: 'foo-bar'}),
  })).pipe(res)
})

app.listen(3751, console.log.bind(null, 'listening:3751'))
