var jwt = require('jsonwebtoken')

var token = jwt.sign({foo: 'bar', message: 'hello'}, process.env.SECRET)
console.log(token)
