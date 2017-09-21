const path = require('path')

module.exports = {
  entry: {
    bundle: './client.js',
  },
  output: {
    path: path.resolve(__dirname, 'dist'),
    filename: '[name].js',
  },
  resolve: {
    extensions: ['.js', '.marko'],
  },
  module: {
    rules: [
      {
        test: /\.marko/,
        loader: 'marko-loader',
      }
    ],
  },
}
