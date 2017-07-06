const async = require('async');
const request = require('request');

const req = () => request(
  'http://localhost:3750',
  (_e, _r, body) => process.stdout.write(body)
);

const reqset = Array(20).fill(req);

async.parallel(
  reqset,
  console.error
);
