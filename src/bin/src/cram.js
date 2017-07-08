// #!/usr/bin/env node

var fs = require('fs')
var args = process.argv
var name = args[args.length - 1]
var file = fs.readFileSync(name).toString()
file = file.replace(/\n/g, '\t')
fs.writeFileSync(name + '.cram', file)
