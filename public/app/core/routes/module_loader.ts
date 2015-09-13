
declare var require : any;

export class ModuleLoader {
  lazy: any;

  constructor(moduleName) {

    this.lazy = ["$q", "$route", "$rootScope", function($q, $route, $rootScope) {
      var defered = $q.defer();

      require.ensure([], function () {
        require('../../features/org/all');
        defered.resolve();
      });

      return defered.promise;
    }];

  }
}
