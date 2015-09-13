'use strict';

var webpack = require('webpack');
var path = require('path');

function absDir(relative) {
  return path.join(__dirname, 'public_gen', relative);
}

var webpackConfig = {
  context: __dirname + '/public_gen/app/',
  entry: {
    'app'        : './app.js',
  },
  output: {
    path: __dirname + '/public_gen/app/',
    filename: 'bundle.js',
  },
  plugins: [
    // This replaces shim stuff in RequireJS.
    new webpack.ProvidePlugin({
      $: "jquery",
      jQuery: "jquery",
      "windows.jQuery": "jquery",
    })
  ],
  module: {
    unknownContextRegExp: /$^/,
    noParse: [
      /\.html/,
      /\.ts/,
      /\.map/,
      /cloudwatch/,
      /app\/panels/,
      /\.json/,
      /moment\.js/
    ],
    // so that JSX can be used.
    loaders: [
    ]
  },
  resolve: {
    alias: {
      "kbn":  absDir('/app/components/kbn.js'),
      "extend-jquery":  absDir('/app/components/extend-jquery.js'),
      "lodash":  absDir('/app/components/lodash.extended.js'),
      "settings":  absDir('/app/components/settings.js'),
      "config": absDir('/app/components/config.js'),
      "store": absDir('/app/components/store.js'),
      "app": absDir('/app/app.js'),
      "aws-sdk":  absDir('/vendor/aws-sdk/dist/aws-sdk.js'),
      "angular": absDir('/vendor/angular/angular.js'),
      "angular-ui": absDir('/vendor/angular-ui/angular-bootstrap.js'),
      "angular-dragdrop": absDir('/vendor/angular-native-dragdrop/draganddrop.js'),
      "angular-sanitize": absDir('/vendor/angular-sanitize/angular-sanitize.js'),
      "angular-route": absDir('/vendor/angular-route/angular-route.js'),
      "angular-strap": absDir('/vendor/angular-other/angular-strap.js'),
      "moment": absDir('/vendor/moment.js'),
      "bootstrap": absDir('/vendor/bootstrap/bootstrap.js'),
      "jquery": absDir('/vendor/jquery/dist/jquery.js'),
      "filesaver": absDir('/vendor/filesaver.js'),
      "ZeroClipboard": absDir('/vendor/ZeroClipboard.js'),
      "bootstrap-tagsinput": absDir('/vendor/tagsinput/bootstrap-tagsinput.js'),
      "spectrum": absDir('/vendor/spectrum.js'),
      "lodash-src": absDir('/vendor/lodash.js'),
      "bindonce": absDir('/vendor/angular-bindonce/bindonce.js'),
    },
    root: [
      __dirname + '/public_gen',
    ],
  }
};

module.exports = webpackConfig;
