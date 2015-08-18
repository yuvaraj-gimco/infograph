define([
  'angular',
],
function (angular) {
  'use strict';

  var module = angular.module('grafana.directives');

  module.directive('metricQueryEditorMeta', function() {
    return {templateUrl: 'app/plugins/datasource/meta/partials/query.editor.html'};
  });

});
