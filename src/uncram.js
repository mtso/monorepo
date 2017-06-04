// recursively search through repo directories
// ignore the directories in .gitignore (like node_modules?)
// uncram files with .cram in them

'use strict';

var fs = require('fs');
var addrs = fs.readdirSync('./');
var ext = /\.cram$/;

var ignored = fs
  .readFileSync('.gitignore')
  .toString()
  .split('\n')
  .filter(function(i) { return i !== ''; });

addrs.forEach(function(addr) {
  if (!ext.test(addr)) {
    return;
  }
  // ignore if in gitignore
  if (ignored.some(function(i) {
    var regex = new RegExp(i);
    return regex.test(addr)
  })) {
    return;
  }
  var data = fs.readFileSync(addr).toString();
  fs.writeFileSync(addr.replace(ext, ''), uncram(data));
});

function uncram(data) {
  return data.replace(/\t/g, '\n');
}
