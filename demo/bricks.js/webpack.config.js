const path = require('path')
const HtmlWebpackPlugin = require('html-webpack-plugin')

const injectIndex = new HtmlWebpackPlugin({
  title: 'picdemo',
  inject: 'body',
})

module.exports = {
  entry: {
    bundle: path.resolve(__dirname, 'index'),
  },
  output: {
    path: path.resolve(__dirname, 'dist'),
    filename: '[name].js',
  },
  resolve: {
    extensions: [ '.js', '.jsx', '.json' ],
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
      {
        test: /\.s[ac]ss$/,
        exclude: /node_modules/,
        loader: 'style-loader!css-loader!sass-loader',
      },
    ],
  },
  plugins: [
    injectIndex,
  ],
}
