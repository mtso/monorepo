const path = require('path')

// const MATCH_ALL_NON_RELATIVE_IMPORTS = /^\w.*$/i;

module.exports = [
  {
    target: 'web',
    entry: {
      bundle: path.resolve(__dirname, 'render'),
    },
    output: {
      path: path.resolve(__dirname, 'dist'),
      filename: '[name].js',
    },
    resolve: {
      extensions: ['.js', '.jsx'],
    },
    module: {
      rules: [
        {
          test: /\.jsx?$/,
          exclude: /node_modules/,
          loader: 'babel-loader',
          query: {
            presets: ['es2015', 'stage-0'],
            plugins: [
              // 'external-helpers',
              // 'transform-class-properties',
              ['transform-react-jsx', { pragma: 'h' }],
            ],
          },
        },
      ],
    },
  },
  {
    target: 'node',
    entry: {
      render: path.resolve(__dirname, 'render'),
    },
    // externals: [MATCH_ALL_NON_RELATIVE_IMPORTS, {
    //   './render.js': 'commonjs ./render.js',
    // }],
    output: {
      path: path.resolve(__dirname, 'dist'),
      filename: '[name].js',
      libraryTarget: 'commonjs2',
    },
    resolve: {
      extensions: ['.js', '.jsx'],
    },
    module: {
      rules: [
        {
          test: /\.jsx?$/,
          exclude: /node_modules/,
          loader: 'babel-loader',
          query: {
            presets: ['es2015', 'stage-0'],
            plugins: [
              // 'external-helpers',
              // 'transform-class-properties',
              ['transform-react-jsx', { pragma: 'h' }],
            ],
          },
        },
      ],
    },
    // node: {
    //   // global: true,
    //   // process: false,
    //   // Buffer: false,
    //   __filename: false,
    //   __dirname: false,
    //   // setImmediate: false,
    // },
  }
]
