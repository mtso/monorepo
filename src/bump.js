// #!/usr/bin/env node

const packageFile = './package.json' // path relative to command invocation
const partTable = {
  major: 0,
  minor: 1,
  patch: 2,
}

const fs = require('fs');

fs.readFile(packageFile, parseInfo(function(info) {
  const previous = info.version;
  const version = info.version.split('.');

  const args = process.argv;
  const partArg = args[args.length - 1];
  var part = partTable[partArg.toLowerCase()];
  
  if (part === undefined) {
    part = 2;
  }

  // Bump.
  version[part] = +version[part] + 1;
  info.version = version.join('.');

  fs.writeFile(packageFile, JSON.stringify(info, null, 2), {flag: 'w'},
    function(err) {
      if (err) {
        throw err;
      }
      console.log('Bumped version', previous + '->' + info.version);
    }
  );
}));

// Helper funcs:

function parseInfo(callback) {
  return function(err, data) {
    if (err) {
      throw err;
    }
    const info = JSON.parse(data.toString());
    callback(info);
  }
}
