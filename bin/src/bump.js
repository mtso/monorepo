// #!/usr/bin/env node
// bump is a script to increment the version string in a package.json file.
// Assumes the semver format: [major].[minor].[patch]
//
// Usage:
// $ bump [major|minor|patch]

'use strict';

const fs = require('fs');

// Assume we are running in the root directory of the package
// (relative path to package.json from command invocation).
const packageFile = './package.json';

// Maps parts to the index of an array containing the strings split by '.'
const partTable = {
  major: 0,
  minor: 1,
  patch: 2,
};

fs.readFile(packageFile, parseInfo(function(info) {
  // Save string for rendering result.
  const previous = info.version;
  const version = info.version.split('.');

  const args = process.argv;
  
  // Take the last argument in the command to determine part.
  const partArg = args[args.length - 1];
  var part = partTable[partArg.toLowerCase()];
  
  // Default to the patch part of the version if none is given.
  if (part === undefined) {
    part = 2;
  }

  // Increment part.
  version[part] = +version[part] + 1;
  info.version = version.join('.');

  // Overwrite package.json with the new version.
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

// parseInfo parses the read data into JSON.
function parseInfo(callback) {
  return function(err, data) {
    if (err) {
      throw err;
    }
    // TODO: wrap in a try-catch block.
    const info = JSON.parse(data.toString());
    callback(info);
  }
}
