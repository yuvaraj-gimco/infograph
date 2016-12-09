

export class NavModel {

  getAlertingNav(subPage) {
    return {
      section: {
        title: 'Alerting',
        url: '/plugins',
        icon: 'icon-gf icon-gf-alert'
      },
      navItems: [
        {title: 'Alert List', active: subPage === 0, url: 'alerting/list', icon: 'fa fa-list-ul'},
        {title: 'Notifications', active: subPage === 1, url: 'alerting/notifications', icon: 'fa fa-bell-o'},
      ]
    };
  }
}


var navModel = new NavModel();
export {navModel};
