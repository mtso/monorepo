var summarize = require('summarize')
var superagent = require('superagent')

superagent
  .get('http://blog.felipe.rs/2017/07/07/where-do-type-systems-come-from/')
  // .get('https://news.ycombinator.com/item?id=14726130')
  .then(function(res) {
    console.log(summarize(res.text))
  })
  .catch(console.warn)
