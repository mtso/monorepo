module.exports = {
  target: 'node',
  entry: __dirname + '/src/uncram.js',
  output: {
    path: __dirname + '/bin',
    filename: 'uncram.min.js',
  }
}
