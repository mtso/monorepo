const http = require('http');

const log = (handler) => {
  return (req, res) => {
    const start = new Date;

    handler(req, res);

    const diff = Date.now() - start;
    console.log(diff + 'ms', req.url);
  }
}

const handler = (req, res) => {
  res.writeHead(200)
  res.end('hello~')
}

const server = http.createServer(log(handler))
server.listen(3750, () => console.log('listening on 3750'))
