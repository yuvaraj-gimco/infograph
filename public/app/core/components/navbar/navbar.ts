///<reference path="../../../headers/common.d.ts" />

import config from 'app/core/config';
import _ from 'lodash';
import $ from 'jquery';
import coreModule from '../../core_module';

import './menu';

export class NavbarCtrl {
  model: any;
  section: any;
  isSearching: boolean;

  /** @ngInject */
  constructor(private $rootScope, private contextSrv) {
    this.section = this.model.section;
  }

  showSearch() {
    console.log('show search');
    this.$rootScope.appEvent('show-dash-search');
    this.isSearching = !this.isSearching;
  }
}

export function navbarDirective() {
  return {
    restrict: 'E',
    templateUrl: 'public/app/core/components/navbar/navbar.html',
    controller: NavbarCtrl,
    bindToController: true,
    transclude: true,
    controllerAs: 'ctrl',
    scope: {
      model: "=",
    },
    link: function(scope, elem) {
      elem.addClass('navbar');
    }
  };
}

coreModule.directive('navbar', navbarDirective);
