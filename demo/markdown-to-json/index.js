const fs = require('fs')
const marked = require('marked')
const parseString = require('xml2js').parseString

const markdown = fs.readFileSync('./test.md')
const markup = marked(markdown.toString())
console.log(markup)
parseString(markup, function(err, result) {
  console.log(JSON.stringify(result, null, 2))
})
