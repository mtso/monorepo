const cluster = require('cluster');
const http = require('http');
const cpus = require('os').cpus();

if (cluster.isMaster) {
  console.log(`Master ${process.pid} is running`);

  // Fork to create cluster.Worker instances.
  cpus.forEach(cluster.fork);

  cluster.on('exit', (worker, _c, _s) => {
    console.log(`worker ${worker.process.pid} died`);
  });
} else {
  // Workers can share any TCP connection.
  http
    .createServer((req, res) => {
      res.writeHead(200);
      res.end('hello~\n');
      console.log(`Worker ${process.pid} ${req.url}`);
    })
    .listen(3750);

  console.log(`Worker ${process.pid} started`);
}
