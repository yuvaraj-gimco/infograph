///<reference path="../../../headers/common.d.ts" />

import _ from 'lodash';
import {coreModule, appEvents} from 'app/core/core';

export class DashMenuCtrl {

  /** @ngInject */
  constructor($scope, $rootScope) {
  }
}

export function dashMenuDirective() {
  return {
    restrict: 'E',
    templateUrl: 'public/app/features/dashboard/dashnav/menu.html',
    controller: DashMenuCtrl,
    bindToController: true,
    controllerAs: 'ctrl',
    scope: {
      dashboard: "="
    },
    link: function(scope, elem) {
      elem.click(function() {

        elem.toggleClass('navbar-dash-btn--open');

      });
    }
  };
}

coreModule.directive('dashnavMenu', dashMenuDirective);
