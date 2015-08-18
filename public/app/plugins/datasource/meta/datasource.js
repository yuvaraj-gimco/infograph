define([
  'angular',
  'lodash',
  'kbn',
  './directives',
],
function (angular) {
  'use strict';

  var module = angular.module('grafana.services');

  module.factory('MetaDatasource', function() {
    return function MetaDatasource(){};
  });

});
