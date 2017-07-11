const path = require('path')
const HtmlWebpackPlugin = require('html-webpack-plugin')

const injectHtml = new HtmlWebpackPlugin({
  filename: '404.html',
  inject: 'body',
  template: path.resolve(__dirname, '404.template.html'),
})

module.exports = {
  entry: {
    bundle: path.resolve(__dirname, 'app/index'),
  },
  resolve: {
    extensions: [ '.js', '.jsx' ],
  },
  output: {
    path: path.resolve(__dirname, 'dist'),
    filename: '[name].js',
  },
  module: {
    rules: [
      {
        test: /\.jsx?$/,
        exclude: /node_modules/,
        loader: 'babel-loader',
        query: {
          presets: [ 'es2015', 'react', 'stage-0' ],
        },
      },
    ],
  },
  plugins: [
    injectHtml,
  ],
}
