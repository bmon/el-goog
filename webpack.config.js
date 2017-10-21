var path = require('path');
var webpack = require('webpack');
var CopyWebpackPlugin = require('copy-webpack-plugin');

var config = {
  context: __dirname + '/assets/src', // `__dirname` is root of project and `src` is source
  entry: {
    app: './index.jsx',
  },
  output: {
    path: __dirname + '/assets/dist', // `dist` is the destination
    publicPath: "/assets/",
    filename: 'bundle.js',
  },

  module: {
    loaders: [
      {
        test: /\.jsx?$/, exclude: /node_modules/, loader: 'babel-loader',
        query: {
          presets:['env', 'react']
        }
      }, // to transform JSX into JS
      {
        test: /\.css$/,
        loader: 'style-loader!css-loader'
      },
    ],
  },

  plugins: [
    new CopyWebpackPlugin([
      { from: '../static', to: '.' }
    ])
  ],
  resolve: {
    extensions: ['.js', '.jsx']
  }
};

module.exports = config;
