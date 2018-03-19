var debug = process.env.NODE_ENV !== "production";
var webpack = require('webpack');

module.exports = {
  context: __dirname,
  devtool: debug ? "inline-sourcemap" : false, // true o false value required
  entry: "./index.js",
  output: {
    path: __dirname + "/public/",
    filename: "bundle.js"
  },
  plugins: debug ? [] : [
    new webpack.optimize.OccurrenceOrderPlugin(), // name was changed from OccurenceOrderPlugin to OccurrenceOrderPlugin
    new webpack.optimize.UglifyJsPlugin({ mangle: false, sourcemap: false }),
  ],
  module: {
    rules: [
        {
          test: /\.jsx?$/,
          exclude: /node_modules/,
          use: {
            loader: "babel-loader",
            options: {
                "presets": ["es2015", "react"]
            }
          }
        }
    ]
  },
  performance: { hints: false }
};