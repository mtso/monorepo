const fs = require('fs');
const path = require('path');
const webpack = require('webpack');

const loadModules = fs
  .readdirSync('node_modules')
  .reduce(function(acc, mod) {
    if (mod === '.bin') {
      return acc;
    }

    acc[mod] = 'commonjs ' + mod;
    return acc;
  });

module.exports = [
  {
    target: 'node',
    entry: { noderender: path.resolve(__dirname, 'renderer/node-render') },
    output: {
      path: path.resolve(__dirname, 'build'),
      filename: '[name].js',
    },
    resolve: { extensions: [ '.js', '.jsx' ] },
    externals: [ loadModules ],
    node: {
      console: false,
      global: false,
      process: false,
      __filename: false,
      __dirname: false,
      Buffer: false,
    },
    module: {
      rules: [
        {
          test: /\.jsx?$/,
          include: [
            path.resolve(__dirname, 'app'),
            path.resolve(__dirname, 'renderer'),
          ],
          loader: 'babel-loader',
          query: { presets: [ 'es2015', 'react' ] },
        },
      ],
    },
    plugins: [
      new webpack.DefinePlugin({ "global.GENTLY": false }),
    ],
  },
  {
    target: 'node',
    entry: { renderMarkup: path.resolve(__dirname, 'renderer/index') },
    output: {
      path: path.resolve(__dirname, 'build'),
      filename: '[name].js',
      libraryTarget: 'commonjs',
      // library: 'renderMarkup',
    },
    resolve: { extensions: [ '.js', '.jsx' ] },
    module: {
      rules: [
        {
          test: /\.jsx?$/,
          include: [
            path.resolve(__dirname, 'app'),
            path.resolve(__dirname, 'renderer'),
          ],
          loader: 'babel-loader',
          query: { presets: [ 'es2015', 'react' ] },
        },
      ],
    },
  },
];
