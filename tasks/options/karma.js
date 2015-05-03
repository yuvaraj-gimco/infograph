module.exports = function(config) {
  'use strict';

  return {
    dev: {
      configFile: '<%= srcDir %>/test/karma.conf.js',
      singleRun: false,
    },
    debug: {
      configFile: '<%= srcDir %>/test/karma.conf.js',
      singleRun: false,
      browsers: ['Chrome']
    },
    test: {
      configFile: '<%= srcDir %>/test/karma.conf.js',
    },
    coveralls: {
      configFile: '<%= srcDir %>/test/karma.conf.js',
      singleRun: true,

      plugins: ['karma-*', '../../../tasks/karma/coveralls'],

      reporters: ['dots','coverage','gf-coveralls'],
      preprocessors: {
        'public/app/**/*.js': ['coverage']
      },
      coverageReporter: {
        type: 'lcov',
        dir: 'coverage/'
      }
    }
  };
};
