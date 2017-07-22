const http = require('http')

const app = module.exports = http.createServer((req, res) => {
  const headerkeys = Object.keys(req.headers)
  headerkeys.forEach((key) => console.log(key, req.headers[key]))
  res.write('hello~')
  res.end()
})

app.listen(3750, console.log.bind(null, 'listening:3750'))
